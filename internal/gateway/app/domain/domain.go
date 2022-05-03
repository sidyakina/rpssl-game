package domain

type Choice struct {
	ID   int32
	Name string
}

type PlayRoundInfo struct {
	Result         string
	PlayerChoice   int32
	ComputerChoice int32
}
