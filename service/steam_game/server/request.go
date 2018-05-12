package server

import "github.com/sigurniv/steam-price/service/steam_game/steam"

type GameSearchResponse struct {
	Games [] steam.Game `json:"games"`
}

type GameResponse struct {
	Name     string  `json:"name"`
	AppId    int     `json:"appId"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}
