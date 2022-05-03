package gateway

type Choice struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type PlayRequest struct {
	PlayerChoice int32 `json:"player"`
}

type PlayResponse struct {
	Result         string `json:"result"`
	PlayerChoice   int32  `json:"player"`
	ComputerChoice int32  `json:"computer"`
}
