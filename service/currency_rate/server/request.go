package server

type AddPairRequest struct {
	Pair string `json:"pair"`
}

type AddPairResponse struct {
	Success bool `json:"success"`
}

type GetPairResponse struct {
	Pair string `json:"pair"`
	Rate float64 `json:"rate"`
}
