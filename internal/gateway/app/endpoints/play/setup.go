package play

import (
	"github.com/sidyakin/rpssl-game/internal/gateway/app/endpoints/play/adapters/choices"
	"github.com/sidyakin/rpssl-game/internal/gateway/app/endpoints/play/request"
	"github.com/sidyakin/rpssl-game/internal/gateway/app/endpoints/play/usecase"
)

func Setup() *request.Handler {
	repo := choices.New()
	uc := usecase.New(repo)

	return request.New(uc)
}
