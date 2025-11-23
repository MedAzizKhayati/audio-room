package handler

import "go.uber.org/fx"

// Module exports all handler dependencies.
// It automatically annotates them as RouteRegistrars.
var Module = fx.Module("handler",
	fx.Provide(
		asRouteRegistrar(NewFibonacciHandler),
		asRouteRegistrar(NewHealthHandler),
		asRouteRegistrar(NewAudioHandler),
	),
)

// asRouteRegistrar is a helper to annotate handlers so they get picked up
// by the main application loop automatically.
func asRouteRegistrar(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(RouteRegistrar)),
		fx.ResultTags(`group:"routes"`),
	)
}
