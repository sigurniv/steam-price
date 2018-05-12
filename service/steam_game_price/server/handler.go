package server

import (
	"net/http"
	"github.com/spf13/viper"
	"errors"
	"github.com/gorilla/mux"
	"github.com/sigurniv/steam-price/service"
	"github.com/sigurniv/steam-price/service/steam_game_price/game"
	"github.com/sigurniv/steam-price/service/steam_game_price/currency"
	"github.com/sigurniv/workgroup"
)

type Handler struct {
	config          *viper.Viper
	currencies      map[string]struct{}
	gameService     *game.Service
	currencyService *currency.Service
}

func NewHandler(config *viper.Viper, gmeService *game.Service, currencyService *currency.Service) *Handler {
	return &Handler{
		config:          config,
		gameService:     gmeService,
		currencyService: currencyService,
	}
}

var errBadRequest = errors.New("Bad Request")

func (handler *Handler) info(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	data := map[string]interface{}{
		"name": handler.config.GetString("service.name"),
	}
	writer.Write(service.Response(data))
}

func (handler *Handler) game(writer http.ResponseWriter, request *http.Request) {
	var response GetGameResponse
	vars := mux.Vars(request)

	paramAppId, ok := vars["appId"]
	if !ok || paramAppId == "" {
		writer.Write(service.ErrResponse(response, errBadRequest))
		return
	}

	paramCur, ok := vars["currency"]
	if !ok || paramCur == "" {
		writer.Write(service.ErrResponse(response, errBadRequest))
		return
	}

	group := workgroup.New()
	group.Go("game", func() (interface{}, error) {
		return handler.gameService.Game(paramAppId)
	})
	group.Go("currency", func() (interface{}, error) {
		return handler.currencyService.Rate("RUB", paramCur)
	})

	results, errs := group.Wait()
	for _, errRaw := range errs {
		err := errRaw.(error)
		writer.Write(service.ErrResponse(response, err))
		return
	}

	gameResponse := results["game"].(game.GameResponse)
	currencyResponse := results["currency"].(currency.RateResponse)

	response.Name = gameResponse.Data.Name
	response.AppId = gameResponse.Data.AppId
	response.Currency = paramCur
	response.Price = currencyResponse.Data.Rate * gameResponse.Data.Price

	writer.Write(service.Response(response))
}
