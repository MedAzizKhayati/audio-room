package service

import "go.uber.org/fx"

// Module exports all service dependencies.
// When you add a new service, just add it here.
var Module = fx.Module("service",
	fx.Provide(
		NewFibonacciService,
	),
)
