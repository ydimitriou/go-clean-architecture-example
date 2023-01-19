package queries

import (
	"github.com/google/uuid"
	"github.com/ydimitriou/go-clean-architecture-example/internal/domain/album"
	"time"
)

// GetAllAlbumsResult represents the result of GetAllAlbums handler
type GetAllAlbumsResult struct {
	ID          uuid.UUID
	Title       string
	Artist      string
	Price       float64
	Description string
	CreatedAt   time.Time
}

// GetAllAlbumsHandler interface for getting all albums form memory
type GetAllAlbumsHandler interface {
	Handle() ([]GetAllAlbumsResult, error)
}

type getAllAlbumsHandler struct {
	repo album.Repository
}

// NewGetAllAlbumsHandler constructor
func NewGetAllAlbumsHandler(repo album.Repository) GetAllAlbumsHandler {
	return getAllAlbumsHandler{repo: repo}
}

// Handle handles get all albums requests
func (h getAllAlbumsHandler) Handle() ([]GetAllAlbumsResult, error) {
	var albums []GetAllAlbumsResult
	res, err := h.repo.GetAll()
	if err != nil {
		return albums, err
	}

	for _, v := range res {
		a := GetAllAlbumsResult{
			ID:          v.ID,
			Title:       v.Title,
			Artist:      v.Artist,
			Price:       v.Price,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
		}
		albums = append(albums, a)
	}

	return albums, nil
}
