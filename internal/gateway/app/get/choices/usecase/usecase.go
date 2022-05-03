package usecase

import (
	"log"

	"github.com/pkg/errors"
)

type Usecase struct {
	repo ChoicesRepo
}

type ChoicesRepo interface {
	GetChoices() ([]Choice, error)
}

func New(repo ChoicesRepo) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) GetChoices() ([]Choice, error) {
	log.Println("getting choices")

	choices, err := u.repo.GetChoices()
	if err != nil {
		return nil, errors.Wrap(err, "repo.GetChoices")
	}

	log.Printf("choices: %v", choices)

	return choices, nil
}
