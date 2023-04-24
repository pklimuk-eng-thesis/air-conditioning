package http

import "github.com/gin-gonic/gin"

var enabledEndpoint = "/enabled"
var getInfoEndpoint = "/info"
var updateEndpoint = "/update"

func SetupRouter(r *gin.Engine, h *ACHandler) {
	r.GET(getInfoEndpoint, h.GetInfo)
	r.PATCH(enabledEndpoint, h.ToggleEnabled)
	r.PATCH(updateEndpoint, h.UpdateACSettings)
}
