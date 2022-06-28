package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerCHealthCheckAdapter struct{}

func NewControllerHealthCheckAdapter() *ControllerCHealthCheckAdapter {
	return &ControllerCHealthCheckAdapter{}
}

// HealthCheckHandler godoc
// @summary      Health Check
// @description  Health checking for the service
// @id           HealthCheckHandler
// @produce      plain
// @response     200  {string}  string  "OK"
// @router       /healthcheck [get]
func (*ControllerCHealthCheckAdapter) HealthCheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
