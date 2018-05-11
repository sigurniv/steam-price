package server

import (
	"net/http"
	"github.com/spf13/viper"
	"github.com/sigurniv/steam-price/service/currency_rate/rates"
	"github.com/sigurniv/steam-price/service"
	"github.com/sigurniv/steam-price/service/currency_rate/storage"
	"io/ioutil"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"strings"
)

type Handler struct {
	config         *viper.Viper
	ratesService   *rates.Service
	StorageService *storage.Service
	currencies     map[string]struct{}
}

func NewHandler(config *viper.Viper, ratesService *rates.Service, storageService *storage.Service) *Handler {
	return &Handler{
		config:         config,
		ratesService:   ratesService,
		StorageService: storageService,
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

func (handler *Handler) list(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	curList, err := handler.ratesService.CurrencyList()
	writer.Write(service.ErrResponse(curList, err))
}

func (handler *Handler) addPair(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var response AddPairResponse

	requestRaw, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	if err != nil {
		writer.Write(service.ErrResponse(response, err))
		return
	}

	var addPairRequest AddPairRequest
	json.Unmarshal(requestRaw, &addPairRequest)

	if addPairRequest.Pair == "" {
		writer.Write(service.ErrResponse(response, errBadRequest))
		return
	}

	if !strings.Contains(addPairRequest.Pair, "_") {
		writer.Write(service.ErrResponse(response, errors.New("Bad Request format. Use CUR_CUR format")))
		return
	}

	//todo check if pair is valid
	response.Success = handler.StorageService.AddPair(addPairRequest.Pair)
	writer.Write(service.Response(response))
}

func (handler *Handler) getRate(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)
	pair, ok := vars["pair"]

	var response GetPairResponse

	if !ok || pair == "" {
		writer.Write(service.ErrResponse(response, errBadRequest))
		return
	}

	response.Rate = handler.StorageService.GetPair(pair)
	response.Pair = pair
	writer.Write(service.Response(response))
}
