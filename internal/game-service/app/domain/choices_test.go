package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChoiceByRandomNumber(t *testing.T) {
	tests := []struct {
		name         string
		randomNumber int32
		wantID       int32
		wantChoice   string
		wantErr      bool
	}{
		{
			name:         "10%5 = 0 -> id = 1 -> rock",
			randomNumber: 10,
			wantID:       1,
			wantChoice:   Rock,
			wantErr:      false,
		},
		{
			name:         "11%5 = 1 -> id = 2 -> paper",
			randomNumber: 11,
			wantID:       2,
			wantChoice:   Paper,
			wantErr:      false,
		},
		{
			name:         "12%5 = 2 -> id = 3 -> scissors",
			randomNumber: 12,
			wantID:       3,
			wantChoice:   Scissors,
			wantErr:      false,
		},
		{
			name:         "13%5 = 3 -> id = 4 -> spock",
			randomNumber: 13,
			wantID:       4,
			wantChoice:   Spock,
			wantErr:      false,
		},
		{
			name:         "14%5 = 4 -> id = 5 -> lizard",
			randomNumber: 14,
			wantID:       5,
			wantChoice:   Lizard,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChoice, gotID, err := GetChoiceByRandomNumber(tt.randomNumber)

			assert.Equalf(t, tt.wantErr, err != nil, "err: %v", err)
			assert.Equal(t, tt.wantID, gotID)
			assert.Equal(t, tt.wantChoice, gotChoice)
		})
	}
}
