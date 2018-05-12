package server

type GetGameResponse struct {
	AppId    int     `json:"appId"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}
