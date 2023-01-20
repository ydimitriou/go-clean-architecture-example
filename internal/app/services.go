package app

import (
	"github.com/ydimitriou/go-clean-architecture-example/internal/app/album/commands"
	"github.com/ydimitriou/go-clean-architecture-example/internal/app/album/queries"
	"github.com/ydimitriou/go-clean-architecture-example/internal/domain/album"
	"github.com/ydimitriou/go-clean-architecture-example/internal/pkg/time"
	"github.com/ydimitriou/go-clean-architecture-example/internal/pkg/uuid"
)

// Commands contains all command handlers
type Commands struct {
	CreateAlbumHandler commands.CreateAlbumHandler
}

// Queries contains all query handlers
type Queries struct {
	GetAllAlbumsHAndler queries.GetAllAlbumsHandler
	GetAlbumHandler     queries.GetAlbumHandler
}

// AlbumServices contains all the album queries and commands of app layer
type AlbumServices struct {
	Commands Commands
	Queries  Queries
}

// Services contains all exposed services of the application layer
type Services struct {
	AlbumServices AlbumServices
}

// NewServices Bootstraps Application layer dependencies
func NewServices(up uuid.Provider, tp time.Provider, repo album.Repository) Services {
	return Services{
		AlbumServices: AlbumServices{
			Commands: Commands{
				CreateAlbumHandler: commands.NewCreateAlbumHandler(up, tp, repo),
			},
			Queries: Queries{
				GetAllAlbumsHAndler: queries.NewGetAllAlbumsHandler(repo),
				GetAlbumHandler:     queries.NewGetAlbumHandler(repo),
			},
		},
	}
}
