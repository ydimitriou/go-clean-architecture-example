package queries

import (
	"github.com/google/uuid"
	"github.com/ydimitriou/go-clean-architecture-example/internal/domain/album"
	"time"
)

// GetAlbumRequest get album request model
type GetAlbumRequest struct {
	ID uuid.UUID
}

// GetAlbumResult represents the result of GetAlbum request
type GetAlbumResult struct {
	ID          uuid.UUID
	Title       string
	Artist      string
	Price       float64
	Description string
	CreatedAt   time.Time
}

// GetAlbumHandler interface for getting an album from memory
type GetAlbumHandler interface {
	Handle(albumRequest GetAlbumRequest) (*GetAlbumResult, error)
}

type getAlbumHandler struct {
	repo album.Repository
}

// NewGetAlbumHandler constructor
func NewGetAlbumHandler(repo album.Repository) GetAlbumHandler {
	return getAlbumHandler{repo: repo}
}

// Handle handles get album by id requests
func (h getAlbumHandler) Handle(albumRequest GetAlbumRequest) (*GetAlbumResult, error) {
	var albumResult *GetAlbumResult
	res, err := h.repo.GetByID(albumRequest.ID)
	if err != nil || res == nil {
		return albumResult, err
	}
	albumResult = &GetAlbumResult{
		ID:          res.ID,
		Title:       res.Title,
		Artist:      res.Artist,
		Price:       res.Price,
		Description: res.Description,
		CreatedAt:   res.CreatedAt,
	}

	return albumResult, nil
}
