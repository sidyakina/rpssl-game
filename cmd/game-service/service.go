package main

import (
	"context"

	apigameservice "github.com/sidyakina/rpssl-game/pkg/api/game-service"
)

type GetRandomChoiceHandler interface {
	GetRandomChoice(
		ctx context.Context, request *apigameservice.GetRandomChoiceRequest,
	) (
		*apigameservice.GetRandomChoiceResponse, error,
	)
}

type GetChoicesHandler interface {
	GetChoices(
		ctx context.Context, request *apigameservice.GetChoicesRequest,
	) (
		*apigameservice.GetChoicesResponse, error,
	)
}

type PlayHandler interface {
	Play(
		ctx context.Context, request *apigameservice.PlayRequest,
	) (
		*apigameservice.PlayResponse, error,
	)
}

type Service struct {
	GetRandomChoiceHandler
	GetChoicesHandler
	PlayHandler
}

func NewService(
	getRandomChoiceHandler GetRandomChoiceHandler,
	getChoicesHandler GetChoicesHandler,
	playHandler PlayHandler,
) *Service {
	return &Service{
		GetRandomChoiceHandler: getRandomChoiceHandler,
		GetChoicesHandler:      getChoicesHandler,
		PlayHandler:            playHandler,
	}
}
