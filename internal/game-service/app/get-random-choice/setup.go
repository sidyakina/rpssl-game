package getrandomchoice

import (
	"github.com/sidyakina/rpssl-game/internal/game-service/app/get-random-choice/adapters/generator"
	"github.com/sidyakina/rpssl-game/internal/game-service/app/get-random-choice/request"
	"github.com/sidyakina/rpssl-game/internal/game-service/app/get-random-choice/usecase"
)

func Setup(url string) *request.Handler {
	generatorRandomNumbers := generator.New(url)
	uc := usecase.New(generatorRandomNumbers)

	return request.New(uc)
}
