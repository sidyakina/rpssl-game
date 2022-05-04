package gameservice

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/sidyakina/rpssl-game/internal/gateway/app/domain"
	internalerrors "github.com/sidyakina/rpssl-game/internal/gateway/pkg/internal-errors"
	apigameservice "github.com/sidyakina/rpssl-game/pkg/api/game-service"
)

type Service struct {
	client  Client
	timeout time.Duration
}

type Client interface {
	GetRandomChoice(
		ctx context.Context, request *apigameservice.GetRandomChoiceRequest, opts ...grpc.CallOption,
	) (
		*apigameservice.GetRandomChoiceResponse, error,
	)
	GetChoices(
		ctx context.Context, request *apigameservice.GetChoicesRequest, opts ...grpc.CallOption,
	) (
		*apigameservice.GetChoicesResponse, error,
	)
	Play(
		ctx context.Context, request *apigameservice.PlayRequest, opts ...grpc.CallOption,
	) (
		*apigameservice.PlayResponse, error,
	)
}

func New(client Client, timeout time.Duration) *Service {
	return &Service{
		client:  client,
		timeout: timeout,
	}
}

func (s *Service) GetRandomChoice() (*domain.Choice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	response, err := s.client.GetRandomChoice(ctx, &apigameservice.GetRandomChoiceRequest{})
	if err != nil {
		if status.Convert(err).Message() == apigameservice.ErrNotFound.Error() {
			return nil, internalerrors.ErrNotFound
		}

		return nil, err
	}

	if response.Choice == nil {
		return nil, errors.New("nil choice in response")
	}

	return &domain.Choice{ID: response.Choice.ID, Name: response.Choice.Name}, nil
}

func (s *Service) GetChoices() ([]domain.Choice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	response, err := s.client.GetChoices(ctx, &apigameservice.GetChoicesRequest{})
	if err != nil {
		if status.Convert(err).Message() == apigameservice.ErrNotFound.Error() {
			return nil, internalerrors.ErrNotFound
		}

		return nil, err
	}

	choices := make([]domain.Choice, 0, len(response.Choices))
	for i, rawChoice := range response.Choices {
		if rawChoice == nil {
			return nil, errors.Errorf("choices[%v] is nil", i)
		}

		choices = append(choices, domain.Choice{
			ID:   rawChoice.ID,
			Name: rawChoice.Name,
		})
	}

	return choices, nil
}

func (s *Service) Play(playerChoice, computerChoice int32) (result, message string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	request := &apigameservice.PlayRequest{
		Player1ChoiceID: playerChoice,
		Player2ChoiceID: computerChoice,
	}

	response, err := s.client.Play(ctx, request)
	if err != nil {
		if status.Convert(err).Message() == apigameservice.ErrWrongParameters.Error() {
			return "", "", internalerrors.ErrWrongParameters
		}

		return "", "", err
	}

	return response.Result, response.Message, nil
}
