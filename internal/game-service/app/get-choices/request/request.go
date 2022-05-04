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
	GetChoices() ([]string, error)
}

func New(usecase Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) GetChoices(
	_ context.Context, _ *apigameservice.GetChoicesRequest,
) (
	*apigameservice.GetChoicesResponse, error,
) {
	log.Printf("new get choices request")

	rawChoices, err := h.usecase.GetChoices()
	if err != nil {
		log.Printf("failed to handle get choices request: %v", err)

		return nil, apigameservice.ErrInternal
	}

	choices := make([]*apigameservice.Choice, 0, len(rawChoices))

	for i, choice := range rawChoices {
		choices = append(choices, &apigameservice.Choice{
			ID:   int32(i) + 1,
			Name: choice,
		})
	}

	response := &apigameservice.GetChoicesResponse{Choices: choices}

	log.Printf("get choices response: %+v", response)

	return response, nil
}
