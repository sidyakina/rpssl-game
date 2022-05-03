package choices

import "github.com/sidyakin/rpssl-game/internal/gateway/app/domain"

type Repo struct {
}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) GetChoices() ([]domain.Choice, error) {
	// todo use grpc method from game-service
	return []domain.Choice{{ID: 1, Name: "test"}}, nil
}
