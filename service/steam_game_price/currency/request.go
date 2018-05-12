package currency

type RateResponse struct {
	Data struct {
		Pair string  `json:"pair"`
		Rate float64 `json:"rate"`
	} `json:"data"`
	Error string `json:"error"`
}
