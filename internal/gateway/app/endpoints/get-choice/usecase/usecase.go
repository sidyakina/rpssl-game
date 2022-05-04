package usecase

import (
	"log"

	"github.com/pkg/errors"

	"github.com/sidyakina/rpssl-game/internal/gateway/app/domain"
)

type Usecase struct {
	repo ChoicesRepo
}

type ChoicesRepo interface {
	GetRandomChoice() (*domain.Choice, error)
}

func New(repo ChoicesRepo) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) GetChoice() (*domain.Choice, error) {
	log.Println("getting random choice")

	choice, err := u.repo.GetRandomChoice()
	if err != nil {
		return nil, errors.Wrap(err, "repo.GetRandomChoice")
	}

	log.Printf("random choice: %v", choice)

	return choice, nil
}
