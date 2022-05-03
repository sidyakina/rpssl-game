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
	Play(playerChoice int32) (*domain.PlayRoundInfo, error)
}

func New(usecase Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) Handle(rawRequest []byte) (response []byte, code int) {
	request := apigateway.PlayRequest{}
	err := json.Unmarshal(rawRequest, &request)
	if err != nil {
		log.Printf("failed to unmarshal play request from %s: %v", rawRequest, err)

		return nil, http.StatusBadRequest
	}

	info, err := h.usecase.Play(request.PlayerChoice)
	if err != nil {
		log.Printf("failed to handle play request: %v", err)

		if errors.Cause(err) == internalerrors.ErrWrongParameters {
			return nil, http.StatusBadRequest
		}

		return nil, http.StatusInternalServerError
	}

	rawResponse := apigateway.PlayResponse{
		Result:         info.Result,
		PlayerChoice:   info.PlayerChoice,
		ComputerChoice: info.ComputerChoice,
	}

	data, err := json.Marshal(&rawResponse)
	if err != nil {
		log.Printf("failed to marshal play response: %v", err)

		return nil, http.StatusInternalServerError
	}

	return data, http.StatusOK
}
