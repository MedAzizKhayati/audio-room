package handler

import "github.com/gin-gonic/gin"

// RouteRegistrar defines the interface that all handlers must implement
// to automatically register their routes with the Gin engine.
type RouteRegistrar interface {
	RegisterRoutes(r *gin.Engine)
}
