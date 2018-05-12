package game

type GameResponse struct {
	Data struct {
		Name     string  `json:"name"`
		AppId    int     `json:"appId"`
		Price    float64 `json:"price"`
		Currency string  `json:"currency"`
	} `json:"data"`
	Error string `json:"error"`
}
