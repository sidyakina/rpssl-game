package domain

import (
	"fmt"

	"github.com/pkg/errors"
)

const (
	Won  = "won"
	Lost = "lost"
	Tie  = "tie"
)

func getRules() map[string]string {
	return map[string]string{
		Scissors + Paper:  "cuts",
		Paper + Rock:      "covers",
		Rock + Lizard:     "crushes",
		Lizard + Spock:    "poisons",
		Spock + Scissors:  "smashes",
		Scissors + Lizard: "decapitates",
		Lizard + Paper:    "eat",
		Paper + Spock:     "disproves",
		Spock + Rock:      "vaporizes",
		Rock + Scissors:   "crushes",
	}
}

func GetResult(player1ChoiceID, player2ChoiceID int32) (result, message string, err error) {
	player1Choice, err := GetChoiceByID(player1ChoiceID)
	if err != nil {
		return "", "", errors.Wrap(err, "wrong player1 choice id")
	}

	player2Choice, err := GetChoiceByID(player2ChoiceID)
	if err != nil {
		return "", "", errors.Wrap(err, "wrong player2 choice id")
	}

	rules := getRules()

	action, ok := rules[player1Choice+player2Choice]
	if ok {
		return Won, formatMessage(player1Choice, player2Choice, action), nil
	}

	action, ok = rules[player2Choice+player1Choice]
	if ok {
		return Lost, formatMessage(player2Choice, player1Choice, action), nil
	}

	return Tie, "", nil
}

func formatMessage(wonPlayerChoice, lostPlayerChoice, action string) string {
	return fmt.Sprintf("%v %v %v", wonPlayerChoice, action, lostPlayerChoice)
}
