package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResult(t *testing.T) {
	rockID := int32(1)
	paperID := int32(2)
	scissorsID := int32(3)
	spockID := int32(4)
	lizardID := int32(5)

	tests := []struct {
		name            string
		player1ChoiceID int32
		player2ChoiceID int32
		wantResult      string
		wantMessage     string
		wantErr         bool
	}{
		// wrong choices
		{
			name:            "wrong player1 choice",
			player1ChoiceID: 0,
			player2ChoiceID: paperID,
			wantResult:      "",
			wantMessage:     "",
			wantErr:         true,
		},
		{
			name:            "wrong player2 choice",
			player1ChoiceID: paperID,
			player2ChoiceID: 6,
			wantResult:      "",
			wantMessage:     "",
			wantErr:         true,
		},
		// tie
		{
			name:            "rock and rock",
			player1ChoiceID: rockID,
			player2ChoiceID: rockID,
			wantResult:      "tie",
			wantMessage:     "",
			wantErr:         false,
		},
		{
			name:            "paper and paper",
			player1ChoiceID: paperID,
			player2ChoiceID: paperID,
			wantResult:      "tie",
			wantMessage:     "",
			wantErr:         false,
		},
		{
			name:            "scissors and scissors",
			player1ChoiceID: scissorsID,
			player2ChoiceID: scissorsID,
			wantResult:      "tie",
			wantMessage:     "",
			wantErr:         false,
		},
		{
			name:            "spock and spock",
			player1ChoiceID: spockID,
			player2ChoiceID: spockID,
			wantResult:      "tie",
			wantMessage:     "",
			wantErr:         false,
		},
		{
			name:            "lizard and lizard",
			player1ChoiceID: lizardID,
			player2ChoiceID: lizardID,
			wantResult:      "tie",
			wantMessage:     "",
			wantErr:         false,
		},
		// won
		{
			name:            "won, scissors cuts paper",
			player1ChoiceID: scissorsID,
			player2ChoiceID: paperID,
			wantResult:      "won",
			wantMessage:     "scissors cuts paper",
			wantErr:         false,
		},
		{
			name:            "won, paper covers rock",
			player1ChoiceID: paperID,
			player2ChoiceID: rockID,
			wantResult:      "won",
			wantMessage:     "paper covers rock",
			wantErr:         false,
		},
		{
			name:            "won, rock crushes lizard",
			player1ChoiceID: rockID,
			player2ChoiceID: lizardID,
			wantResult:      "won",
			wantMessage:     "rock crushes lizard",
			wantErr:         false,
		},
		{
			name:            "won, lizard poisons spock",
			player1ChoiceID: lizardID,
			player2ChoiceID: spockID,
			wantResult:      "won",
			wantMessage:     "lizard poisons spock",
			wantErr:         false,
		},
		{
			name:            "won, spock smashes scissors",
			player1ChoiceID: spockID,
			player2ChoiceID: scissorsID,
			wantResult:      "won",
			wantMessage:     "spock smashes scissors",
			wantErr:         false,
		},
		{
			name:            "won, scissors decapitates lizard",
			player1ChoiceID: scissorsID,
			player2ChoiceID: lizardID,
			wantResult:      "won",
			wantMessage:     "scissors decapitates lizard",
			wantErr:         false,
		},
		{
			name:            "won, lizard eat paper",
			player1ChoiceID: lizardID,
			player2ChoiceID: paperID,
			wantResult:      "won",
			wantMessage:     "lizard eat paper",
			wantErr:         false,
		},
		{
			name:            "won, paper disproves spock",
			player1ChoiceID: paperID,
			player2ChoiceID: spockID,
			wantResult:      "won",
			wantMessage:     "paper disproves spock",
			wantErr:         false,
		},
		{
			name:            "won, spock vaporizes rock",
			player1ChoiceID: spockID,
			player2ChoiceID: rockID,
			wantResult:      "won",
			wantMessage:     "spock vaporizes rock",
			wantErr:         false,
		},
		{
			name:            "won, rock crushes scissors",
			player1ChoiceID: rockID,
			player2ChoiceID: scissorsID,
			wantResult:      "won",
			wantMessage:     "rock crushes scissors",
			wantErr:         false,
		},
		// lost
		{
			name:            "lost, scissors cuts paper",
			player1ChoiceID: paperID,
			player2ChoiceID: scissorsID,
			wantResult:      "lost",
			wantMessage:     "scissors cuts paper",
			wantErr:         false,
		},
		{
			name:            "lost, paper covers rock",
			player1ChoiceID: rockID,
			player2ChoiceID: paperID,
			wantResult:      "lost",
			wantMessage:     "paper covers rock",
			wantErr:         false,
		},
		{
			name:            "lost, rock crushes lizard",
			player1ChoiceID: lizardID,
			player2ChoiceID: rockID,
			wantResult:      "lost",
			wantMessage:     "rock crushes lizard",
			wantErr:         false,
		},
		{
			name:            "lost, lizard poisons spock",
			player1ChoiceID: spockID,
			player2ChoiceID: lizardID,
			wantResult:      "lost",
			wantMessage:     "lizard poisons spock",
			wantErr:         false,
		},
		{
			name:            "lost, spock smashes scissors",
			player1ChoiceID: scissorsID,
			player2ChoiceID: spockID,
			wantResult:      "lost",
			wantMessage:     "spock smashes scissors",
			wantErr:         false,
		},
		{
			name:            "lost, scissors decapitates lizard",
			player1ChoiceID: lizardID,
			player2ChoiceID: scissorsID,
			wantResult:      "lost",
			wantMessage:     "scissors decapitates lizard",
			wantErr:         false,
		},
		{
			name:            "lost, lizard eat paper",
			player1ChoiceID: paperID,
			player2ChoiceID: lizardID,
			wantResult:      "lost",
			wantMessage:     "lizard eat paper",
			wantErr:         false,
		},
		{
			name:            "lost, paper disproves spock",
			player1ChoiceID: spockID,
			player2ChoiceID: paperID,
			wantResult:      "lost",
			wantMessage:     "paper disproves spock",
			wantErr:         false,
		},
		{
			name:            "lost, spock vaporizes rock",
			player1ChoiceID: rockID,
			player2ChoiceID: spockID,
			wantResult:      "lost",
			wantMessage:     "spock vaporizes rock",
			wantErr:         false,
		},
		{
			name:            "lost, rock crushes scissors",
			player1ChoiceID: scissorsID,
			player2ChoiceID: rockID,
			wantResult:      "lost",
			wantMessage:     "rock crushes scissors",
			wantErr:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotMessage, err := GetResult(tt.player1ChoiceID, tt.player2ChoiceID)

			assert.Equalf(t, tt.wantErr, err != nil, "err: %v", err)
			assert.Equal(t, tt.wantResult, gotResult)
			assert.Equal(t, tt.wantMessage, gotMessage)
		})
	}
}
