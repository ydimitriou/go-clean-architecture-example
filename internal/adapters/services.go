package adapters

import (
	"github.com/ydimitriou/go-clean-architecture-example/internal/adapters/storage/memory"
	"github.com/ydimitriou/go-clean-architecture-example/internal/domain/album"
)

// Services contains the available adapters services
type Services struct {
	AlbumRepository album.Repository
}

// NewServices instantiates the adapter services
func NewServices() Services {
	return Services{
		AlbumRepository: memory.NewRepository(),
	}
}
