package rates

import (
	"github.com/spf13/viper"
	"github.com/sigurniv/steam-price/service/currency_rate/storage"
	"time"
	"bytes"
	"go.uber.org/zap"
	"strings"
)

type Updater struct {
	storageService    *storage.Service
	ratesService      *Service
	updateIntervalMin int
	logger            *zap.SugaredLogger
}

func NewUpdater(config *viper.Viper, logger *zap.SugaredLogger, storageService *storage.Service, ratesService *Service) *Updater {
	return &Updater{
		storageService:    storageService,
		ratesService:      ratesService,
		updateIntervalMin: config.GetInt("rates.updateIntervalMin"),
		logger:            logger,
	}
}

func (updater *Updater) Run() {
	ticker := time.NewTicker(time.Duration(updater.updateIntervalMin) * time.Second) //todo
	isUpdating := false

	for {
		select {
		case <-ticker.C:
			if !isUpdating {
				isUpdating = true
				updater.updateRates()
				isUpdating = false
			}
		}
	}
}

func (updater *Updater) updateRates() {
	var pairs []string
	for pair, _ := range updater.storageService.GetPairs() {
		pairs = append(pairs, pair)
	}

	buf := bytes.Buffer{}
	for _, pair := range pairs {
		buf.WriteString(pair)
		buf.WriteString(",")
	}

	pairStr := strings.TrimRight(buf.String(), ",")
	if pairStr == "" {
		return
	}

	rates, err := updater.ratesService.Rate(pairStr)
	if err != nil {
		updater.logger.Errorw("Error updating rates", "error", err.Error())
	}

	for pair, _ := range updater.storageService.GetPairs() {
		if rate, ok := rates[pair]; ok {
			updater.storageService.SetRate(pair, rate)
		}
	}
}
