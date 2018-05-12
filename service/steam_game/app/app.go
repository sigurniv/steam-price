package app

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"fmt"
	"context"
	"github.com/sigurniv/steam-price/service/steam_game/server"
)

type Application struct {
	Server *server.Server
	Config *viper.Viper
	Logger *zap.SugaredLogger
}

func New(config *viper.Viper) (app *Application, err error) {
	zapLogger, _ := zap.NewProduction()
	logger := zapLogger.Sugar()

	srv, err := server.New(config, logger)
	if err != nil {
		return nil, err
	}

	return &Application{
		Server: srv,
		Logger: logger,
		Config: config,
	}, err
}

func (app *Application) Run() {
	mode := "production"
	if app.Config.GetBool("app.debug") {
		mode = "debug"
	}

	app.Logger.Infow(fmt.Sprintf("App running in %s mode", mode))

	go app.Server.Run()
}

func (app *Application) Shutdown(ctx context.Context) error {
	app.Logger.Infow("Shutting down app")

	var err error
	err = app.Server.Shutdown(ctx)

	return err
}
