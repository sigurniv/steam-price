package game

import (
	"github.com/spf13/viper"
	"net/http"
	"crypto/tls"
	"time"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type Service struct {
	config *viper.Viper
	apiUrl string
}

func New(config *viper.Viper) *Service {
	return &Service{
		config: config,
		apiUrl: config.GetString("backends.steam_game"),
	}
}

func (svc *Service) Game(appId string) (GameResponse, error) {
	var response GameResponse
	responseRaw, err := svc.request(fmt.Sprintf("%s/game/%s", svc.apiUrl, appId))
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(responseRaw, &response)
	if err != nil {
		return response, err
	}

	if response.Error != "" {
		return response, fmt.Errorf("backend error : %s", response.Error)
	}

	return response, err
}

func (svc *Service) request(url string) ([]byte, error) {
	var response []byte
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	client.Timeout = 5 * time.Second
	respRaw, err := client.Get(url)
	if err != nil {
		return response, err
	}

	defer respRaw.Body.Close()

	return ioutil.ReadAll(respRaw.Body)
}
