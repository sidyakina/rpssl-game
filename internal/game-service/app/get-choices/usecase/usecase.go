package usecase

import (
	"log"

	"github.com/sidyakina/rpssl-game/internal/game-service/app/domain"
)

type Usecase struct {
}

func New() *Usecase {
	return &Usecase{}
}

func (u *Usecase) GetChoices() ([]string, error) {
	log.Println("getting choices")

	choices := domain.GetChoices()

	log.Printf("choices: %v", choices)

	return choices, nil
}
