package usecase

import (
	"log"

	"github.com/pkg/errors"

	"github.com/sidyakina/rpssl-game/internal/gateway/app/domain"
)

type Usecase struct {
	service Service
}

type Service interface {
	GetRandomChoice() (*domain.Choice, error)
	Play(playerChoice, computerChoice int32) (result string, err error)
}

func New(service Service) *Usecase {
	return &Usecase{service: service}
}

func (u *Usecase) Play(playerChoice int32) (*domain.PlayRoundInfo, error) {
	log.Printf("playing game round with player choice id: %v", playerChoice)

	computerChoice, err := u.service.GetRandomChoice()
	if err != nil {
		return nil, errors.Wrap(err, "service.GetRandomChoice")
	}

	log.Printf("computer choice: %v", computerChoice)

	result, err := u.service.Play(playerChoice, computerChoice.ID)
	if err != nil {
		return nil, errors.Wrap(err, "service.Play")
	}

	log.Printf("result is: %v", result)

	info := &domain.PlayRoundInfo{
		Result:         result,
		PlayerChoice:   playerChoice,
		ComputerChoice: computerChoice.ID,
	}

	return info, nil
}
