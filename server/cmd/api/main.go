package main

import (
	"app/internal/handler"
	"app/internal/middleware"
	"app/internal/service"
	"context"
	"fmt"
	"net/http"

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

	// Register global middlewares
	for _, m := range params.Middlewares {
		r.Use(m.Handle())
	}

	// Register all routes automatically
	for _, h := range params.Handlers {
		h.RegisterRoutes(r)
	}

	// Define what happens on Start and Stop
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Server starting on :8080...")
			// Run in a goroutine so it doesn't block Fx
			go func() {
				if err := r.Run(":8080"); err != nil && err != http.ErrServerClosed {
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
