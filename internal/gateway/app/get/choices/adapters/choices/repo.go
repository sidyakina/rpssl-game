package choices

import "github.com/sidyakin/rpssl-game/internal/gateway/app/get/choices/usecase"

type Repo struct {
}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) GetChoices() ([]usecase.Choice, error) {
	// todo use grpc method from game-service
	return []usecase.Choice{{ID: 1, Name: "test"}}, nil
}
