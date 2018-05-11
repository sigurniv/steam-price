package rates

import (
	"github.com/spf13/viper"
	"net/http"
	"time"
	"fmt"
	"io/ioutil"
	"crypto/tls"
	"encoding/json"
)

type Service struct {
	apiUrl     string
	currencies map[string]struct{}
}

func New(config *viper.Viper) *Service {
	return &Service{
		apiUrl: config.GetString("rates.url"),
	}
}

func (srv *Service) CurrencyList() ([]string, error) {
	response, err := srv.request("currencies")
	if err != nil {
		return []string{}, err
	}

	var curListResponse CurrencyListResponse
	err = json.Unmarshal(response, &curListResponse)

	var currencies []string
	for currency, _ := range curListResponse.Results {
		currencies = append(currencies, currency)
	}

	return currencies, err
}

func (srv *Service) Rate(pairs string) (map[string]float64, error) {
	var ratesResponse RatesResponse
	rates := make(map[string]float64)

	pairsStr := fmt.Sprintf("convert?q=%s", pairs)
	response, err := srv.request(pairsStr)
	if err != nil {
		return rates, err
	}

	err = json.Unmarshal(response, &ratesResponse)
	for cur, item := range ratesResponse.Results {
		rates[cur] = item.Val
	}

	return rates, err
}

func (srv *Service) request(method string) ([]byte, error) {
	var response []byte
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	client.Timeout = 5 * time.Second
	respRaw, err := client.Get(fmt.Sprintf("%s%s", srv.apiUrl, method))
	if err != nil {
		return response, err
	}

	defer respRaw.Body.Close()

	return ioutil.ReadAll(respRaw.Body)
}
