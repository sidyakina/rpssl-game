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
	GetRandomChoice() (choice string, choiceID int32, err error)
}

func New(usecase Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) GetRandomChoice(
	_ context.Context, _ *apigameservice.GetRandomChoiceRequest,
) (
	*apigameservice.GetRandomChoiceResponse, error,
) {
	log.Printf("new get random choice request")

	choice, choiceID, err := h.usecase.GetRandomChoice()
	if err != nil {
		log.Printf("failed to handle get random choice request: %v", err)

		return nil, apigameservice.ErrInternal
	}

	response := &apigameservice.GetRandomChoiceResponse{
		Choice: &apigameservice.Choice{
			ID:   choiceID,
			Name: choice,
		},
	}

	log.Printf("get random choice response: %+v", response)

	return response, nil
}
