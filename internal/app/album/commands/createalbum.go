package commands

import (
	"github.com/ydimitriou/go-clean-architecture-example/internal/domain/album"
	"github.com/ydimitriou/go-clean-architecture-example/internal/pkg/time"
	"github.com/ydimitriou/go-clean-architecture-example/internal/pkg/uuid"
)

// CreateAlbumRequest Model of CreateAlbumRequestHandler
type CreateAlbumRequest struct {
	Title       string
	Artist      string
	Price       float64
	Description string
}

// CreateAlbumHandler interface for creating albums
type CreateAlbumHandler interface {
	Handle(req CreateAlbumRequest) error
}

type createAlbumHandler struct {
	uuidProvider uuid.Provider
	timeProvider time.Provider
	repo         album.Repository
}

// NewCreateAlbumHandler constructor
func NewCreateAlbumHandler(up uuid.Provider, tp time.Provider, repo album.Repository) CreateAlbumHandler {
	return createAlbumHandler{
		uuidProvider: up,
		timeProvider: tp,
		repo:         repo,
	}
}

// Handle handles create album request
func (h createAlbumHandler) Handle(req CreateAlbumRequest) error {
	a := album.Album{
		ID:          h.uuidProvider.NewUUID(),
		Title:       req.Title,
		Artist:      req.Artist,
		Price:       req.Price,
		Description: req.Description,
		CreatedAt:   h.timeProvider.Now(),
	}

	err := h.repo.Create(a)
	if err != nil {
		return err
	}

	return nil
}
