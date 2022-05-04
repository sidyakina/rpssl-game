package play

import (
	"github.com/sidyakina/rpssl-game/internal/gateway/app/endpoints/play/request"
	"github.com/sidyakina/rpssl-game/internal/gateway/app/endpoints/play/usecase"
)

func Setup(service usecase.Service) *request.Handler {
	uc := usecase.New(service)

	return request.New(uc)
}
