package album

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// MockRepository mocks the album repository
type MockRepository struct {
	mock.Mock
}

// GetByID mock
func (m *MockRepository) GetByID(id uuid.UUID) (*Album, error) {
	args := m.Called(id)

	return args.Get(0).(*Album), args.Error(1)
}

// GetAll mock
func (m *MockRepository) GetAll() ([]Album, error) {
	args := m.Called()

	return args.Get(0).([]Album), args.Error(1)
}

// Create mock
func (m *MockRepository) Create(album Album) error {
	args := m.Called(album)

	return args.Error(0)
}

// Update mock
func (m *MockRepository) Update(album Album) error {
	args := m.Called(album)

	return args.Error(0)
}

// Delete mock
func (m *MockRepository) Delete(id uuid.UUID) error {
	args := m.Called(id)

	return args.Error(0)
}
