package getchoices

import (
	"github.com/sidyakina/rpssl-game/internal/gateway/app/endpoints/get-choices/request"
	"github.com/sidyakina/rpssl-game/internal/gateway/app/endpoints/get-choices/usecase"
)

func Setup(repo usecase.ChoicesRepo) *request.Handler {
	uc := usecase.New(repo)

	return request.New(uc)
}
