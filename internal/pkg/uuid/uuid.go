package uuid

import "github.com/google/uuid"

// Provider interface for uuid generator abstraction
type Provider interface {
	NewUUID() uuid.UUID
}

type uuidProvider struct {
}

// NewUUIDProvider constructor to return the default uuid provider
func NewUUIDProvider() Provider {
	return uuidProvider{}
}

func (u uuidProvider) NewUUID() uuid.UUID {
	return uuid.New()
}
