package main

import (
	"github.com/ydimitriou/go-clean-architecture-example/internal/adapters"
	"github.com/ydimitriou/go-clean-architecture-example/internal/app"
	"github.com/ydimitriou/go-clean-architecture-example/internal/pkg/time"
	"github.com/ydimitriou/go-clean-architecture-example/internal/pkg/uuid"
	"github.com/ydimitriou/go-clean-architecture-example/internal/ports"
)

func main() {
	tp := time.NewTimeProvider()
	up := uuid.NewUUIDProvider()
	outputAdapters := adapters.NewServices()
	appServices := app.NewServices(up, tp, outputAdapters.AlbumRepository)
	inputPorts := ports.NewServices(appServices)
	inputPorts.HTTPServer.ListenAndServe(":8080")
}
