package request

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sidyakina/rpssl-game/internal/gateway/app/domain"
	apigateway "github.com/sidyakina/rpssl-game/pkg/api/gateway"
)

type Handler struct {
	usecase Usecase
}

type Usecase interface {
	GetChoice() (*domain.Choice, error)
}

func New(usecase Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) Handle() (response []byte, code int) {
	rawChoice, err := h.usecase.GetChoice()
	if err != nil {
		log.Printf("failed to handle get choice request: %v", err)

		return nil, http.StatusInternalServerError
	}

	choice := apigateway.Choice{ID: rawChoice.ID, Name: rawChoice.Name}

	data, err := json.Marshal(&choice)
	if err != nil {
		log.Printf("failed to marshal get choice response: %v", err)

		return nil, http.StatusInternalServerError
	}

	return data, http.StatusOK
}
