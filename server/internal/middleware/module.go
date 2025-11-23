package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// Module exports all middleware dependencies.
var Module = fx.Module("middleware",
	fx.Provide(
		asMiddleware(NewTimingMiddleware),
	),
)

// asMiddleware annotates the middleware to be collected in the "middlewares" group
func asMiddleware(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Middleware)),
		fx.ResultTags(`group:"middlewares"`),
	)
}

type Middleware interface {
	Handle() gin.HandlerFunc
}
