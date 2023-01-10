package album

import (
	"time"

	"github.com/google/uuid"
)

// Album Model that represents the Album
type Album struct {
	ID          uuid.UUID
	Title       string
	Artist      string
	Price       float64
	Description string
	CreatedAT   time.Time
}
