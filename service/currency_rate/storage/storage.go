package storage

import (
	"github.com/spf13/viper"
)

type Service struct {
	storage map[string]float64
}

func New(config *viper.Viper) *Service {
	return &Service{
		storage: make(map[string]float64),
	}
}

func (svc *Service) AddPair(pair string) bool {
	svc.storage[pair] = 0
	return true
}

func (svc *Service) GetPair(pair string) float64 {
	rate, _ := svc.storage[pair]
	return rate
}

func (svc *Service) GetPairs() map[string]float64 {
	return svc.storage
}

func (svc *Service) SetRate(pair string, rate float64) {
	svc.storage[pair] = rate
}
