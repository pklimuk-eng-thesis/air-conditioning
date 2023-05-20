package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pklimuk-eng-thesis/air-conditioning/pkg/domain"
	"github.com/pklimuk-eng-thesis/air-conditioning/pkg/service"
	"github.com/stretchr/testify/assert"
)

func TestACHandler_GetInfo(t *testing.T) {
	acService := new(service.MockACService)
	expectedAC := &domain.AC{Enabled: true, Temperature: 20.0, Humidity: 40.0}
	acService.EXPECT().GetInfo().Return(expectedAC, nil)

	handler := NewACHandler(acService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.GetInfo(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"enabled": true, "temperature": 20.0, "humidity": 40.0}`, w.Body.String())
}

func TestACHandler_GetInfo_Error(t *testing.T) {
	acService := new(service.MockACService)
	acService.EXPECT().GetInfo().Return(nil, errors.New("error"))

	handler := NewACHandler(acService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.GetInfo(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "error", w.Body.String())
}

func TestACHandler_ToggleEnabled(t *testing.T) {
	acService := new(service.MockACService)
	expectedAC := &domain.AC{Enabled: true, Temperature: 20.0, Humidity: 40.0}
	acService.EXPECT().ToggleEnabled().Return(expectedAC, nil)

	handler := NewACHandler(acService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.ToggleEnabled(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"enabled": true, "temperature": 20.0, "humidity": 40.0}`, w.Body.String())
}

func TestACHandler_ToggleEnabled_Error(t *testing.T) {
	acService := new(service.MockACService)
	acService.EXPECT().ToggleEnabled().Return(nil, errors.New("error"))

	handler := NewACHandler(acService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.ToggleEnabled(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "error", w.Body.String())
}

func TestACHandler_UpdateACSettings(t *testing.T) {
	acService := new(service.MockACService)
	expectedAC := &domain.AC{Enabled: true, Temperature: 20.0, Humidity: 40.0}
	acService.EXPECT().UpdateACSettings(float32(20.0), float32(40.0)).Return(expectedAC, nil)

	handler := NewACHandler(acService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader(`{"enabled": true, "temperature": 20, "humidity": 40}`))
	c.Request.Header.Set("Content-Type", "application/json")
	handler.UpdateACSettings(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"enabled": true, "temperature": 20, "humidity": 40}`, w.Body.String())
}

func TestACHandler_UpdateACSettings_BadRequest_NotEnabled(t *testing.T) {
	acService := new(service.MockACService)

	handler := NewACHandler(acService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader(`{"enabled": false, "temperature": 20, "humidity": 40}`))
	c.Request.Header.Set("Content-Type", "application/json")
	handler.UpdateACSettings(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "AC must be enabled", w.Body.String())
}

func TestACHandler_UpdateACSettings_BadRequest_WrongBody(t *testing.T) {
	acService := new(service.MockACService)

	handler := NewACHandler(acService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader(`{}`))
	c.Request.Header.Set("Content-Type", "application/json")
	handler.UpdateACSettings(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestACHandler_UpdateACSettings_Error(t *testing.T) {
	acService := new(service.MockACService)
	acService.EXPECT().UpdateACSettings(float32(20.0), float32(40.0)).Return(nil, errors.New("error"))

	handler := NewACHandler(acService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader(`{"enabled": true, "temperature": 20, "humidity": 40}`))
	c.Request.Header.Set("Content-Type", "application/json")
	handler.UpdateACSettings(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "error", w.Body.String())
}
