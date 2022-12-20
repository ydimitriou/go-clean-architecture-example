package album

import (
	"github.com/google/uuid"
)

//Repository Interface for album
type Repository interface {
	GetByID(id uuid.UUID) (*Album, error)
	GetAll() ([]Album, error)
	Add(album Album) error
	Update(album Album) error
	Delete(id uuid.UUID) error
}
