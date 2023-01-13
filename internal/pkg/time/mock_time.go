package time

import (
	"github.com/stretchr/testify/mock"
	"time"
)

// MockProvider mocks time provider
type MockProvider struct {
	mock.Mock
}

// Now mock
func (m *MockProvider) Now() time.Time {
	args := m.Called()

	return args.Get(0).(time.Time)
}
