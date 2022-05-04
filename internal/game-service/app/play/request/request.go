package request

import (
	"context"
	"log"

	apigameservice "github.com/sidyakina/rpssl-game/pkg/api/game-service"
)

type Handler struct {
	usecase Usecase
}

type Usecase interface {
	Play(player1Choice, player2Choice int32) (result string, err error)
}

func New(usecase Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) Play(
	_ context.Context, request *apigameservice.PlayRequest,
) (
	*apigameservice.PlayResponse, error,
) {
	log.Printf("new play request: %+v", request)

	result, err := h.usecase.Play(request.Player1ChoiceID, request.Player2ChoiceID)
	if err != nil {
		log.Printf("failed to handle play request: %v", err)

		return nil, apigameservice.ErrInternal
	}

	response := &apigameservice.PlayResponse{Result: result}

	log.Printf("play response: %+v", response)

	return response, nil
}
