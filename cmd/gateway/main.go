package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	getchoice "github.com/sidyakin/rpssl-game/internal/gateway/app/endpoints/get-choice"
	getchoices "github.com/sidyakin/rpssl-game/internal/gateway/app/endpoints/get-choices"
	"github.com/sidyakin/rpssl-game/internal/gateway/app/endpoints/play"
	handlerwrapper "github.com/sidyakin/rpssl-game/internal/gateway/pkg/handler-wrapper"
)

func main() {
	config, err := parseConfig()
	if err != nil {
		log.Panicf("failed to parse config: %v", err)
	}

	getChoicesHandler := getchoices.Setup()
	getChoiceHandler := getchoice.Setup()
	playHandler := play.Setup()

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
		Addr:         fmt.Sprintf(":%v", config.Port),
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Printf("starting http server on %v", config.Port)

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
