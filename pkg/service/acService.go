package service

import (
	"errors"
	"log"
	"time"

	"github.com/pklimuk-eng-thesis/air-conditioning/pkg/domain"
)

var ErrACDisabled = errors.New("AC is disabled")

//go:generate --name ACService --output mock_acService.go
type ACService interface {
	GetInfo() (*domain.AC, error)
	ToggleEnabled() (*domain.AC, error)
	UpdateACSettings(desiredTemp float32, desiredHumidity float32) (*domain.AC, error)
}

type acService struct {
	ac *domain.AC
}

func NewACService(ac *domain.AC) ACService {
	return &acService{
		ac: ac,
	}
}

func (s *acService) GetInfo() (*domain.AC, error) {
	if !s.ac.Enabled {
		return &domain.AC{Enabled: false, Temperature: 0.0, Humidity: 0.0}, nil
	}
	return s.ac, nil
}

func (s *acService) ToggleEnabled() (*domain.AC, error) {
	s.ac.Enabled = !s.ac.Enabled
	if !s.ac.Enabled {
		return &domain.AC{Enabled: false, Temperature: 0.0, Humidity: 0.0}, nil
	}
	return s.ac, nil
}

func (s *acService) UpdateACSettings(desiredTemp float32, desiredHumidity float32) (*domain.AC, error) {
	if !s.ac.Enabled {
		return &domain.AC{Enabled: false, Temperature: 0.0, Humidity: 0.0}, ErrACDisabled
	}
	if desiredTemp < 15 || desiredTemp > 30 {
		return s.ac, errors.New("Desired temperature is out of range(15 - 30)")
	}
	if desiredHumidity < 30 || desiredHumidity > 60 {
		return s.ac, errors.New("Desired humidity is out of range (30 - 60)")
	}
	changeSteps := 10
	sleepTime := 5 * time.Second
	tempDelta := (s.ac.Temperature - desiredTemp) / float32(changeSteps)
	humidityDelta := (s.ac.Humidity - desiredHumidity) / float32(changeSteps)
	go func() {
		log.Printf("CHANGING SETTINGS TO --> Temperature: %.2f, Humidity: %.2f", desiredTemp, desiredHumidity)
		for i := 0; i < changeSteps; i++ {
			s.ac.Temperature -= tempDelta
			s.ac.Humidity -= humidityDelta
			time.Sleep(sleepTime)
			log.Printf("CHANGING SETTINGS. Temperature: %.2f, Humidity: %.2f", s.ac.Temperature, s.ac.Humidity)
		}
		log.Printf("CHANGING SETTINGS DONE. Temperature: %.2f, Humidity: %.2f", s.ac.Temperature, s.ac.Humidity)
	}()
	return s.ac, nil
}
