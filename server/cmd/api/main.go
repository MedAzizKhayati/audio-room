package main

import (
	"app/internal/handler"
	"app/internal/middleware"
	"app/internal/service"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		// Load high-level modules
		service.Module,
		handler.Module,
		middleware.Module,

		// Invoke starts the application
		fx.Invoke(registerHooks),
	).Run()
}

// RouteParams groups all handlers together
type RouteParams struct {
	fx.In
	Handlers    []handler.RouteRegistrar `group:"routes"`
	Middlewares []middleware.Middleware  `group:"middlewares"`
}

// registerHooks is the entry point that Fx calls.
func registerHooks(lc fx.Lifecycle, params RouteParams) {
	r := gin.Default()

	// CORS middleware for Railway deployment
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Register global middlewares
	for _, m := range params.Middlewares {
		r.Use(m.Handle())
	}

	// Register all routes automatically
	for _, h := range params.Handlers {
		h.RegisterRoutes(r)
	}

	// Get port from environment variable (Railway sets this)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Define what happens on Start and Stop
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Printf("Server starting on :%s...\n", port)
			// Run in a goroutine so it doesn't block Fx
			go func() {
				if err := r.Run(":" + port); err != nil && err != http.ErrServerClosed {
					fmt.Printf("Error starting server: %v\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Server stopping...")
			return nil
		},
	})
}
