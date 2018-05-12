package server

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"github.com/gorilla/mux"
	"net/http"
	"context"
	"errors"
	"github.com/sigurniv/steam-price/service/steam_game_price/game"
	"github.com/sigurniv/steam-price/service/steam_game_price/currency"
)

type Server struct {
	srv     *http.Server
	logger  *zap.SugaredLogger
	config  *viper.Viper
	Handler *Handler
}

func New(config *viper.Viper, logger *zap.SugaredLogger) (*Server, error) {
	var err error

	port := config.GetString("server.port")
	if port == "" {
		return nil, errors.New("server.port is not specified")
	}

	gameService := game.New(config)
	currencyService := currency.New(config)
	handler := NewHandler(config, gameService, currencyService)

	srv := &Server{
		srv:     &http.Server{Addr: ":" + port},
		logger:  logger,
		config:  config,
		Handler: handler,
	}

	router := mux.NewRouter()
	router.HandleFunc("/info/", handler.info)
	router.HandleFunc("/game/{appId}/{currency}", handler.game)
	http.Handle("/", router)

	return srv, err
}

func (s *Server) Run() {
	if err := s.srv.ListenAndServe(); err != nil {
		s.logger.Errorw("Error starting webserver", "err", err.Error())
		return
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
