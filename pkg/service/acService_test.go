package service

import (
	"testing"
	"time"

	"github.com/pklimuk-eng-thesis/air-conditioning/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func Test_ACService_NewACService(t *testing.T) {
	ac := &domain.AC{Enabled: true, Temperature: 20.0, Humidity: 40.0}
	acService := NewACService(ac)
	assert.NotNil(t, acService)
}

func Test_ACService_GetInfo(t *testing.T) {
	tests := []struct {
		name    string
		s       acService
		want    *domain.AC
		wantErr bool
	}{
		{name: "AC Enabled",
			s:       acService{&domain.AC{Enabled: true, Temperature: 20.0, Humidity: 40.0}},
			want:    &domain.AC{Enabled: true, Temperature: 20.0, Humidity: 40.0},
			wantErr: false},
		{name: "AC Disabled",
			s:       acService{&domain.AC{Enabled: false, Temperature: 20.0, Humidity: 40.0}},
			want:    &domain.AC{Enabled: false, Temperature: 0.0, Humidity: 0.0},
			wantErr: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.s.GetInfo()

			assert.Equal(t, test.wantErr, err != nil)
			assert.Equal(t, test.want.Enabled, got.Enabled)
			assert.Equal(t, test.want.Temperature, got.Temperature)
			assert.Equal(t, test.want.Humidity, got.Humidity)
		})
	}
}

func Test_ACService_ToggleEnabled(t *testing.T) {
	tests := []struct {
		name    string
		s       acService
		want    *domain.AC
		wantErr bool
	}{
		{name: "AC Enabled",
			s:       acService{&domain.AC{Enabled: true, Temperature: 20.0, Humidity: 40.0}},
			want:    &domain.AC{Enabled: false, Temperature: 0.0, Humidity: 0.0},
			wantErr: false},
		{name: "AC Disabled",
			s:       acService{&domain.AC{Enabled: false, Temperature: 20.0, Humidity: 40.0}},
			want:    &domain.AC{Enabled: true, Temperature: 20.0, Humidity: 40.0},
			wantErr: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.s.ToggleEnabled()

			assert.Equal(t, test.wantErr, err != nil)
			assert.Equal(t, test.want.Enabled, got.Enabled)
			assert.Equal(t, test.want.Temperature, got.Temperature)
			assert.Equal(t, test.want.Humidity, got.Humidity)
		})
	}
}

func Test_ACService_UpdateACSettings(t *testing.T) {
	tests := []struct {
		name        string
		s           acService
		desiredTemp float32
		desiredHum  float32
		want        *domain.AC
		wantErr     bool
	}{
		{name: "AC Enabled - Temperature out of range min",
			s:           acService{&domain.AC{Enabled: true, Temperature: 20.0, Humidity: 40.0}},
			desiredTemp: float32(10.0),
			desiredHum:  float32(50.0),
			want:        &domain.AC{Enabled: true, Temperature: 20.0, Humidity: 40.0},
			wantErr:     true},
		{name: "AC Enabled - Temperature out of range max",
			s:           acService{&domain.AC{Enabled: true, Temperature: 20.0, Humidity: 40.0}},
			desiredTemp: float32(35.0),
			desiredHum:  float32(50.0),
			want:        &domain.AC{Enabled: true, Temperature: 20.0, Humidity: 40.0},
			wantErr:     true},
		{name: "AC Enabled - Humidity out of range min",
			s:           acService{&domain.AC{Enabled: true, Temperature: 20.0, Humidity: 40.0}},
			desiredTemp: float32(25.0),
			desiredHum:  float32(10.0),
			want:        &domain.AC{Enabled: true, Temperature: 20.0, Humidity: 40.0},
			wantErr:     true},
		{name: "AC Enabled - Humidity out of range max",
			s:           acService{&domain.AC{Enabled: true, Temperature: 20.0, Humidity: 40.0}},
			desiredTemp: float32(25.0),
			desiredHum:  float32(70.0),
			want:        &domain.AC{Enabled: true, Temperature: 20.0, Humidity: 40.0},
			wantErr:     true},
		{name: "AC Disabled",
			s:           acService{&domain.AC{Enabled: false, Temperature: 20.0, Humidity: 40.0}},
			desiredTemp: float32(25.0),
			desiredHum:  float32(50.0),
			want:        &domain.AC{Enabled: false, Temperature: 0.0, Humidity: 0.0},
			wantErr:     true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.s.UpdateACSettings(test.desiredTemp, test.desiredHum)

			assert.Equal(t, test.wantErr, err != nil)
			assert.Equal(t, test.want.Enabled, got.Enabled)
			assert.Equal(t, test.want.Temperature, got.Temperature)
			assert.Equal(t, test.want.Humidity, got.Humidity)
		})
	}
}

func Test_UpdateACSettings_OK(t *testing.T) {
	acService := &acService{&domain.AC{Enabled: true, Temperature: 20, Humidity: 40}}
	desiredTemp := float32(30)
	desiredHumidity := float32(50)
	expectedAC := domain.AC{Enabled: true, Temperature: desiredTemp, Humidity: desiredHumidity}
	got, err := acService.UpdateACSettings(desiredTemp, desiredHumidity)

	// Wait for the goroutine to complete partially
	time.Sleep(5 * time.Second)

	assert.NoError(t, err)
	assert.Equal(t, expectedAC.Enabled, got.Enabled)
	assert.Greater(t, got.Temperature, float32(20))
	assert.Greater(t, got.Humidity, float32(40))
}
