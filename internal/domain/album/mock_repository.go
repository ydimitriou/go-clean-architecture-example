package album

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

//MockRepository mocks the album repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetByID(id uuid.UUID) (*Album, error) {
	args := m.Called(id)

	return args.Get(0).(*Album), args.Error(1)
}
