package choices

import (
	"github.com/sidyakin/rpssl-game/internal/gateway/app/domain"
)

type Repo struct {
}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) GetRandomChoice() (*domain.Choice, error) {
	// todo use grpc method from game-service
	return &domain.Choice{ID: 1, Name: "test"}, nil
}

func (r *Repo) Play(playerChoice, computerChoice int32) (result string, err error) {
	// todo use grpc method from game-service
	return "win", nil
}
