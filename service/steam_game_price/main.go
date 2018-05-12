package main

import (
	"log"
	"flag"
	"os"
	"os/signal"
	"time"
	"fmt"
	"path/filepath"
	"github.com/spf13/viper"
	"context"
	"github.com/sigurniv/steam-price/service/steam_game_price/app"
)

func main() {
	var configPath = flag.String("config", "", "absolute path to the config file directory")
	flag.Parse()
	config, err := initConfig(*configPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	application, err := app.New(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	application.Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	application.Shutdown(ctx)
}

func initConfig(path string) (*viper.Viper, error) {
	if path == "" {
		path = getBinaryDir(path)
	}

	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(path)
	err := v.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Fatal error config file: %s \n", err)
	}

	return v, err
}

func getBinaryDir(path string) string {
	if path == "" {
		path = "."
	}

	//current working app directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err == nil {
		path = dir
	}

	return path
}
