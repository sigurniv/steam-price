package steam

type GetAppListResponse struct {
	AppList struct {
		Apps struct {
			App []struct {
				AppId int    `json:"appId"`
				Name  string `json:"name"`
			} `json:"app"`
		} `json:"apps"`
	} `json:"applist"`
}

type AppDetailsResponse map[string]struct {
	Data struct {
		Name  string `json:"name"`
		AppId int    `json:"steam_appid"`
		PriceOverview struct {
			Currency string  `json:"currency"`
			Price    float64 `json:"final"`
		} `json:"price_overview"`
	} `json:"data"`
}
