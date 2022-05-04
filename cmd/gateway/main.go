package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	gameservice "github.com/sidyakina/rpssl-game/internal/gateway/app/adapters/game-service"
	getchoice "github.com/sidyakina/rpssl-game/internal/gateway/app/endpoints/get-choice"
	getchoices "github.com/sidyakina/rpssl-game/internal/gateway/app/endpoints/get-choices"
	"github.com/sidyakina/rpssl-game/internal/gateway/app/endpoints/play"
	handlerwrapper "github.com/sidyakina/rpssl-game/internal/gateway/pkg/handler-wrapper"
	apigameservice "github.com/sidyakina/rpssl-game/pkg/api/game-service"
)

func main() {
	config, err := parseConfig()
	if err != nil {
		log.Panicf("failed to parse config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.GPPCConnectTimeout)
	defer cancel()

	gameServiceConn, err := grpc.DialContext(ctx, config.GameServiceURI)
	if err != nil {
		log.Panicf("failed to connect to game service")
	}

	defer func() {
		err := gameServiceConn.Close()
		if err != nil {
			log.Printf("failed to close connect to game service")
		}
	}()

	gameServiceClient := apigameservice.NewGameServiceClient(gameServiceConn)
	gameServiceAdapter := gameservice.New(gameServiceClient, config.GRPCRequestTimeout)

	getChoicesHandler := getchoices.Setup(gameServiceAdapter)
	getChoiceHandler := getchoice.Setup(gameServiceAdapter)
	playHandler := play.Setup(gameServiceAdapter)

	mux := http.NewServeMux()

	mux.HandleFunc("/choices", func(writer http.ResponseWriter, request *http.Request) {
		handlerwrapper.Get(writer, request, "choices", getChoicesHandler.Handle)
	})

	mux.HandleFunc("/choice", func(writer http.ResponseWriter, request *http.Request) {
		handlerwrapper.Get(writer, request, "choice", getChoiceHandler.Handle)
	})

	mux.HandleFunc("/play", func(writer http.ResponseWriter, request *http.Request) {
		handlerwrapper.Post(writer, request, "play", playHandler.Handle)
	})

	server := http.Server{
		Addr:         fmt.Sprintf(":%v", config.HTTPPort),
		Handler:      mux,
		ReadTimeout:  config.HTTPReadTimeout,
		WriteTimeout: config.HTTPWriteTimeout,
	}

	log.Printf("starting http server on %v", config.HTTPPort)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Panicf("failed to listen and serve: %v", err)
		}
	}()

	log.Printf("server started")

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGKILL)

	<-ch

	log.Println("got signal, server stopped")
}
