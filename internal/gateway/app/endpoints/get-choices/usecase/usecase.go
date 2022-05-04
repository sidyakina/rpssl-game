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
	GetChoices() ([]domain.Choice, error)
}

func New(repo ChoicesRepo) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) GetChoices() ([]domain.Choice, error) {
	log.Println("getting choices")

	choices, err := u.repo.GetChoices()
	if err != nil {
		return nil, errors.Wrap(err, "repo.GetChoices")
	}

	log.Printf("choices: %v", choices)

	return choices, nil
}
