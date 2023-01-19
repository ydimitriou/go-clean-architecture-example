package time

import "time"

// Provider interface for time abstraction
type Provider interface {
	Now() time.Time
}

type timeProvider struct {
}

// NewTimeProvider constructor to return the default time provider
func NewTimeProvider() Provider {
	return timeProvider{}
}

func (t timeProvider) Now() time.Time {
	return time.Now()
}
