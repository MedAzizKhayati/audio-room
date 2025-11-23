package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) RegisterRoutes(router *gin.Engine) {
	createPath := GetControllerPathCreator("/health")
	router.GET(createPath("/ready"), WrapHandler(h.handleReady))
	router.GET(createPath("/live"), WrapHandler(h.handleLive))
}

func (h *HealthHandler) handleReady(c *gin.Context) (any, int, error) {
	return gin.H{
		"ok": true,
	}, http.StatusOK, nil
}

func (h *HealthHandler) handleLive(c *gin.Context) (any, int, error) {
	return gin.H{
		"ok": true,
	}, http.StatusOK, nil
}
