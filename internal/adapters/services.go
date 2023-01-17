package adapters

import (
	"github.com/ydimitriou/go-clean-architecture-example/internal/adapters/storage/memory"
	"github.com/ydimitriou/go-clean-architecture-example/internal/domain/album"
)

type Services struct {
	AlbumRepository album.Repository
}

func NewServices() Services {
	return Services{
		AlbumRepository: memory.NewRepository(),
	}
}
