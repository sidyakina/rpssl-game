package getchoices

import (
	"github.com/sidyakina/rpssl-game/internal/game-service/app/get-choices/request"
	"github.com/sidyakina/rpssl-game/internal/game-service/app/get-choices/usecase"
)

func Setup() *request.Handler {
	uc := usecase.New()

	return request.New(uc)
}
