package getchoices

import (
	"github.com/sidyakin/rpssl-game/internal/gateway/app/endpoints/get-choices/adapters/choices"
	"github.com/sidyakin/rpssl-game/internal/gateway/app/endpoints/get-choices/request"
	"github.com/sidyakin/rpssl-game/internal/gateway/app/endpoints/get-choices/usecase"
)

func Setup() *request.Handler {
	repo := choices.New()
	uc := usecase.New(repo)

	return request.New(uc)
}
