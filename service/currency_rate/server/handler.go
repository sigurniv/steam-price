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
	"go.uber.org/zap"
)

type Handler struct {
	config         *viper.Viper
	ratesService   *rates.Service
	StorageService *storage.Service
	logger         *zap.SugaredLogger
}

func NewHandler(config *viper.Viper, ratesService *rates.Service, storageService *storage.Service, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		config:         config,
		ratesService:   ratesService,
		StorageService: storageService,
		logger:         logger,
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

	response.Success = handler.StorageService.AddPair(addPairRequest.Pair)

	//update rate on the fly for testing purposes
	if response.Success {
		ratesData, err := handler.ratesService.Rate(addPairRequest.Pair)
		if err != nil {
			handler.logger.Errorw("Error updating rates", "error", err.Error())
		}

		for pair, _ := range handler.StorageService.GetPairs() {
			if rate, ok := ratesData[pair]; ok {
				handler.StorageService.SetRate(pair, rate)
			}
		}
	}

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
	if response.Rate == 0 {
		//update rate on the fly for testing purposes
		ratesData, err := handler.ratesService.Rate(pair)
		if err != nil {
			handler.logger.Errorw("Error updating rates", "error", err.Error())
		}

		if rate, ok := ratesData[pair]; ok {
			handler.StorageService.SetRate(pair, rate)
			response.Rate = rate
		}
	}

	response.Pair = pair
	writer.Write(service.Response(response))
}
