package getchoice

import (
	"github.com/sidyakin/rpssl-game/internal/gateway/app/endpoints/get-choice/request"
	"github.com/sidyakin/rpssl-game/internal/gateway/app/endpoints/get-choice/usecase"
)

func Setup(repo usecase.ChoicesRepo) *request.Handler {
	uc := usecase.New(repo)

	return request.New(uc)
}
