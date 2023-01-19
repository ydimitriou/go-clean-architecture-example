package ports

import (
	"github.com/ydimitriou/go-clean-architecture-example/internal/app"
	"github.com/ydimitriou/go-clean-architecture-example/internal/ports/http"
)

// Services contains ports services
type Services struct {
	HTTPServer http.Server
}

// NewServices instantiates ports services
func NewServices(as app.Services) Services {
	return Services{HTTPServer: http.NewServer(as)}
}
