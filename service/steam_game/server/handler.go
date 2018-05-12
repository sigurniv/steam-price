package server

import (
	"net/http"
	"github.com/spf13/viper"
	"github.com/sigurniv/steam-price/service"
	"github.com/sigurniv/steam-price/service/steam_game/steam"
	"github.com/gorilla/mux"
	"errors"
)

type Handler struct {
	config       *viper.Viper
	currencies   map[string]struct{}
	steamService *steam.Service
}

func NewHandler(config *viper.Viper, steamService *steam.Service) *Handler {
	return &Handler{
		config:       config,
		steamService: steamService,
	}
}

var (
	errBadRequest  = errors.New("Bad Request")
	errUnknownGame = errors.New("unknown game")
)

func (handler *Handler) info(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	data := map[string]interface{}{
		"name": handler.config.GetString("service.name"),
	}
	writer.Write(service.Response(data))
}

func (handler *Handler) gameSearch(writer http.ResponseWriter, request *http.Request) {
	var response GameSearchResponse
	vars := mux.Vars(request)
	paramGame, ok := vars["game"]

	if !ok || paramGame == "" {
		writer.Write(service.ErrResponse(response, errBadRequest))
		return
	}

	games, err := handler.steamService.GameSearch(paramGame)
	if err != nil {
		writer.Write(service.ErrResponse(response, err))
		return
	}

	response.Games = games
	writer.Write(service.Response(response))
}

func (handler *Handler) game(writer http.ResponseWriter, request *http.Request) {
	var response steam.GameDetails
	vars := mux.Vars(request)
	appId, ok := vars["appId"]
	if !ok || appId == "" {
		writer.Write(service.ErrResponse(response, errBadRequest))
		return
	}

	gameDetails, err := handler.steamService.AppDetails(appId)
	if err != nil {
		writer.Write(service.ErrResponse(response, errBadRequest))
		return
	}

	if gameDetails.Name == "" {
		writer.Write(service.ErrResponse(response, errUnknownGame))
		return
	}

	writer.Write(service.Response(gameDetails))
}
