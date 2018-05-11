package rates

type CurrencyListResponse struct {
	Results map[string]interface{} `json:"results"`
}

type RatesResponse struct {
	Results map[string]struct {
		Val float64 `json:"val"`
	} `json:"results"`
}

type RatesReponseItem struct{

}
