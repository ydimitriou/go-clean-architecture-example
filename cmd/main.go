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
	adapters := adapters.NewServices()
	appServices := app.NewServices(up, tp, adapters.AlbumRepository)
	ports := ports.NewServices(appServices)
	ports.HTTPServer.ListenAndServe(":8080")
}
