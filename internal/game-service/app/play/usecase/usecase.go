package usecase

import (
	"log"

	"github.com/sidyakina/rpssl-game/internal/game-service/app/domain"
	apigameservice "github.com/sidyakina/rpssl-game/pkg/api/game-service"
)

type Usecase struct {
}

func New() *Usecase {
	return &Usecase{}
}

func (u *Usecase) Play(player1Choice, player2Choice int32) (result, message string, err error) {
	log.Printf(
		"playing game round with player1 choice id: %v, player2 choice id %v",
		player1Choice, player2Choice,
	)

	result, message, err = domain.GetResult(player1Choice, player2Choice)
	if err != nil {
		log.Printf("failed to get result: %v", err)

		return "", "", apigameservice.ErrWrongParameters
	}

	log.Printf("result: %v, message: %v", result, message)

	return result, message, nil
}
