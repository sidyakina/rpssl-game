package usecase

import (
	"log"

	"github.com/pkg/errors"

	"github.com/sidyakin/rpssl-game/internal/gateway/app/domain"
)

type Usecase struct {
	repo ChoicesRepo
}

type ChoicesRepo interface {
	GetRandomChoice() (*domain.Choice, error)
	Play(playerChoice, computerChoice int32) (result string, err error)
}

func New(repo ChoicesRepo) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) Play(playerChoice int32) (*domain.PlayRoundInfo, error) {
	log.Printf("playing game round with player choice id: %v", playerChoice)

	computerChoice, err := u.repo.GetRandomChoice()
	if err != nil {
		return nil, errors.Wrap(err, "repo.GetRandomChoice")
	}

	log.Printf("computer choice: %v", computerChoice)

	result, err := u.repo.Play(playerChoice, computerChoice.ID)
	if err != nil {
		return nil, errors.Wrap(err, "repo.Play")
	}

	log.Printf("result is: %v", result)

	info := &domain.PlayRoundInfo{
		Result:         result,
		PlayerChoice:   playerChoice,
		ComputerChoice: computerChoice.ID,
	}

	return info, nil
}
