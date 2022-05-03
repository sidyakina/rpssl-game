package request

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/sidyakin/rpssl-game/internal/gateway/app/domain"
	internalerrors "github.com/sidyakin/rpssl-game/internal/gateway/pkg/internal-errors"

	apigateway "github.com/sidyakin/rpssl-game/pkg/api/gateway"
)

type Handler struct {
	usecase Usecase
}

type Usecase interface {
	GetChoices() ([]domain.Choice, error)
}

func New(usecase Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) Handle() (response []byte, code int) {
	rawChoices, err := h.usecase.GetChoices()
	if err != nil {
		log.Printf("failed to handle get choices request: %v", err)

		if errors.Cause(err) == internalerrors.ErrNotFound {
			return nil, http.StatusNotFound
		}

		return nil, http.StatusInternalServerError
	}

	choices := make([]apigateway.Choice, 0, len(rawChoices))
	for _, rawChoice := range rawChoices {
		choices = append(choices, apigateway.Choice{ID: rawChoice.ID, Name: rawChoice.Name})
	}

	data, err := json.Marshal(choices)
	if err != nil {
		log.Printf("failed to marshal get choices response: %v", err)

		return nil, http.StatusInternalServerError
	}

	return data, http.StatusOK
}
