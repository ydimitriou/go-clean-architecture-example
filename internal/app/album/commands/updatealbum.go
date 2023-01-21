package commands

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/ydimitriou/go-clean-architecture-example/internal/domain/album"
)

// UpdateAlbumRequest request model for update an album
type UpdateAlbumRequest struct {
	ID          uuid.UUID
	Title       string
	Artist      string
	Price       float64
	Description string
}

// UpdateAlbumHandler interface for updating album details
type UpdateAlbumHandler interface {
	Handle(req UpdateAlbumRequest) error
}

type updateAlbumHandler struct {
	repo album.Repository
}

// NewUpdateAlbumHandler constructor
func NewUpdateAlbumHandler(repo album.Repository) UpdateAlbumHandler {
	return updateAlbumHandler{repo: repo}
}

// Handle handles update album requests
func (h updateAlbumHandler) Handle(req UpdateAlbumRequest) error {
	a, err := h.repo.GetByID(req.ID)
	if err != nil {
		return err
	}
	if a == nil {
		return fmt.Errorf("update failed, album with id %v does not exist", req.ID.String())
	}
	a.Title = req.Title
	a.Artist = req.Artist
	a.Price = req.Price
	a.Description = req.Description

	return h.repo.Update(*a)
}
