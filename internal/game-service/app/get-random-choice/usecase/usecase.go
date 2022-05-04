package usecase

import (
	"log"

	"github.com/pkg/errors"

	"github.com/sidyakina/rpssl-game/internal/game-service/app/domain"
)

type Usecase struct {
	generator RandomNumberGenerator
}

type RandomNumberGenerator interface {
	GetRandomNumber() (int32, error)
}

func New(generator RandomNumberGenerator) *Usecase {
	return &Usecase{generator: generator}
}

func (u *Usecase) GetRandomChoice() (choice string, choiceID int32, err error) {
	log.Println("getting random choice")

	randomNumber, err := u.generator.GetRandomNumber()
	if err != nil {
		return "", 0, errors.Wrap(err, "get random number")
	}

	log.Printf("received random number: %v", randomNumber)

	choice, id, err := domain.GetChoiceByRandomNumber(randomNumber)
	if err != nil {
		return "", 0, errors.Wrap(err, "get choice by random number")
	}

	log.Printf("random choice: %v with id %v", choice, id)

	return choice, id, nil
}
