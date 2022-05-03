package getchoice

import (
	"github.com/sidyakin/rpssl-game/internal/gateway/app/endpoints/get-choice/adapters/choices"
	"github.com/sidyakin/rpssl-game/internal/gateway/app/endpoints/get-choice/request"
	"github.com/sidyakin/rpssl-game/internal/gateway/app/endpoints/get-choice/usecase"
)

func Setup() *request.Handler {
	repo := choices.New()
	uc := usecase.New(repo)

	return request.New(uc)
}
