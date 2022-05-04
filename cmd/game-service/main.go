package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	getchoices "github.com/sidyakina/rpssl-game/internal/game-service/app/get-choices"
	getrandomchoice "github.com/sidyakina/rpssl-game/internal/game-service/app/get-random-choice"
	"github.com/sidyakina/rpssl-game/internal/game-service/app/play"
	apigameservice "github.com/sidyakina/rpssl-game/pkg/api/game-service"
)

func main() {
	config, err := parseConfig()
	if err != nil {
		log.Panicf("failed to parse config: %v", err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", config.GRPCPort))
	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	service := NewService(
		getrandomchoice.Setup(config.RandomNumberGeneratorURI),
		getchoices.Setup(),
		play.Setup(),
	)

	apigameservice.RegisterGameServiceServer(grpcServer, service)

	go func() {
		err := grpcServer.Serve(listener)
		if err != nil {
			log.Panicf("failed to serve grpc server: %v", err)
		}
	}()

	defer grpcServer.Stop()

	log.Printf("server started")

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGKILL)

	<-ch

	log.Println("got signal, server stopped")
}
