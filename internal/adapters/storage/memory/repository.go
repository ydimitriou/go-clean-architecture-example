package memory

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/ydimitriou/go-clean-architecture-example/internal/domain/album"
)

// Repository is the in-memory repository implementation
type Repository struct {
	albums map[string]album.Album
}

// NewRepository constructor
func NewRepository() Repository {
	return Repository{
		albums: make(map[string]album.Album),
	}
}

// Create creates the given album
func (mr Repository) Create(album album.Album) error {
	mr.albums[album.ID.String()] = album

	return nil
}

// Update updates the given album
func (mr Repository) Update(album album.Album) error {
	_, exists := mr.albums[album.ID.String()]
	if !exists {
		return fmt.Errorf("update failed, id %v does not exist", album.ID.String())
	}
	mr.albums[album.ID.String()] = album

	return nil
}

// Delete deletes the album with the given id
func (mr Repository) Delete(id uuid.UUID) error {
	_, exists := mr.albums[id.String()]
	if !exists {
		return fmt.Errorf("album with id %v not found", id.String())
	}
	delete(mr.albums, id.String())

	return nil
}

// GetByID retrieves the album with the given id
func (mr Repository) GetByID(id uuid.UUID) (*album.Album, error) {
	val, exists := mr.albums[id.String()]
	if !exists {
		err := fmt.Errorf("album with id %v not found", id.String())
		return nil, err
	}

	return &val, nil
}

// GetAll retrieves all stored albums
func (mr Repository) GetAll() ([]album.Album, error) {
	var albums []album.Album
	for _, v := range mr.albums {
		albums = append(albums, v)
	}

	return albums, nil
}
