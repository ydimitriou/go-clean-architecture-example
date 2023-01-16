package commands

import (
	"github.com/ydimitriou/go-clean-architecture-example/internal/domain/album"
	"github.com/ydimitriou/go-clean-architecture-example/internal/pkg/time"
	"github.com/ydimitriou/go-clean-architecture-example/internal/pkg/uuid"
)

// AddAlbumRequest Model of CreateAlbumRequestHandler
type AddAlbumRequest struct {
	Title       string
	Artist      string
	Price       float64
	Description string
}

type addAlbumHandler struct {
	uuidProvider uuid.Provider
	timeProvider time.Provider
	repo         album.Repository
}

func (h addAlbumHandler) Handle(req AddAlbumRequest) error {
	a := album.Album{
		ID:          h.uuidProvider.NewUUID(),
		Title:       req.Title,
		Artist:      req.Artist,
		Price:       req.Price,
		Description: req.Description,
		CreatedAT:   h.timeProvider.Now(),
	}

	err := h.repo.Add(a)
	if err != nil {
		return err
	}

	return nil
}
