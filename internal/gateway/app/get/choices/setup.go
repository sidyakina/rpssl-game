package getchoices

import (
	"github.com/sidyakin/rpssl-game/internal/gateway/app/get/choices/adapters/choices"
	"github.com/sidyakin/rpssl-game/internal/gateway/app/get/choices/request"
	"github.com/sidyakin/rpssl-game/internal/gateway/app/get/choices/usecase"
)

func Setup() *request.Handler {
	repo := choices.New()
	uc := usecase.New(repo)

	return request.New(uc)
}
