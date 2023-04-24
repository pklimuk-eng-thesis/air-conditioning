package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pklimuk-eng-thesis/air-conditioning/pkg/domain"
	service "github.com/pklimuk-eng-thesis/air-conditioning/pkg/service"
)

type ACHandler struct {
	acService service.ACService
}

func NewACHandler(acService service.ACService) *ACHandler {
	return &ACHandler{
		acService: acService,
	}
}

func (h *ACHandler) GetInfo(c *gin.Context) {
	ac, err := h.acService.GetInfo()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, &ac)
}

func (h *ACHandler) ToggleEnabled(c *gin.Context) {
	ac, err := h.acService.ToggleEnabled()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, &ac)
}

func (h *ACHandler) UpdateACSettings(c *gin.Context) {
	var incomingAC domain.AC
	err := c.BindJSON(&incomingAC)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if !incomingAC.Enabled {
		c.String(http.StatusBadRequest, "AC must be enabled")
		return
	}
	ac, err := h.acService.UpdateACSettings(incomingAC.Temperature, incomingAC.Humidity)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, &ac)
}
