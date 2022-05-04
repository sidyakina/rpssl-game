package play

import (
	"github.com/sidyakina/rpssl-game/internal/game-service/app/play/request"
	"github.com/sidyakina/rpssl-game/internal/game-service/app/play/usecase"
)

func Setup() *request.Handler {
	uc := usecase.New()

	return request.New(uc)
}
