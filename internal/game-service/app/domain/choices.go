package domain

import "github.com/pkg/errors"

const (
	NumberChoices = 5

	Rock     = "rock"
	Paper    = "paper"
	Scissors = "scissors"
	Spock    = "spock"
	Lizard   = "lizard"
)

func GetChoices() []string {
	return []string{
		Rock,
		Paper,
		Scissors,
		Spock,
		Lizard,
	}
}

func GetChoiceByID(id int32) (string, error) {
	choices := GetChoices()

	if id < 1 || int(id) > len(choices) {
		return "", errors.Errorf("wrong id %v", id)
	}

	return choices[int(id)-1], nil
}

func GetChoiceByRandomNumber(randomNumber int32) (choice string, id int32, err error) {
	id = randomNumber%NumberChoices + 1 // id [1, NumberChoices]

	choice, err = GetChoiceByID(id)

	return choice, id, err
}
