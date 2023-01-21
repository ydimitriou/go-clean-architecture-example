package commands

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/ydimitriou/go-clean-architecture-example/internal/domain/album"
)

// DeleteAlbumRequest request model for album deletion
type DeleteAlbumRequest struct {
	ID uuid.UUID
}

// DeleteAlbumHandler interface for album deletion
type DeleteAlbumHandler interface {
	Handle(req DeleteAlbumRequest) error
}

type deleteAlbumHandler struct {
	repo album.Repository
}

// NewDeleteAlbumHandler constructor
func NewDeleteAlbumHandler(repo album.Repository) DeleteAlbumHandler {
	return deleteAlbumHandler{repo: repo}
}

// Handle handles delete album requests
func (h deleteAlbumHandler) Handle(req DeleteAlbumRequest) error {
	a, err := h.repo.GetByID(req.ID)
	if err != nil {
		return err
	}
	if a == nil {
		return fmt.Errorf("album with id %v does not exist", req.ID.String())
	}
	return h.repo.Delete(req.ID)
}
