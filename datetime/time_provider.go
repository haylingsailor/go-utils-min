package datetime

import (
	"sync"
	"time"
)

// TimeProvider is the interface that wraps methods which allow the setting and
// getting of 'now' which, for the purposes of testing, might have been set to
// an explicit value.
type TimeProvider interface {
	SetNow(t time.Time)

	Now() (t time.Time)

	Until(t time.Time) time.Duration

	Since(t time.Time) time.Duration
}

// NowTimeProvider implements interface TimeProvider to provide the real value
// of 'now' (which is always changing), or an artificially-set value (which
// does not change over time automatically).
type NowTimeProvider struct {
	sync.RWMutex
	// explicitNow is only set if the helper is to return a specific time rather
	// than the current time
	explicitNow *time.Time
}

// SetNow implements interface TimeProvider. If called, then all future calls to
// Now() will return the value passed.
func (p *NowTimeProvider) SetNow(t time.Time) {
	p.Lock()
	defer p.Unlock()

	p.explicitNow = &t
}

// Now implements interface TimeProvider. If SetNow() has not been called, then
// successive calls to Now() will return the actual real (changing) time. If
// SetNow() has been called, then the value returned will always remain the same
// - that passed to SetNow().
func (p *NowTimeProvider) Now() (t time.Time) {
	p.RLock()
	defer p.RUnlock()

	switch p.explicitNow {
	case nil:
		return time.Now()
	default:
		return *p.explicitNow
	}
}

// Since returns the duration that has elapsed since the given time to now. A
// negative duration is returned if the given time is in the future.
func (p *NowTimeProvider) Since(t time.Time) time.Duration {
	return p.Now().Sub(t)
}

// Until returns the duration remaining between now and the time specified. If
// the time given is already in the past, a negative duration is returned.
func (p *NowTimeProvider) Until(t time.Time) time.Duration {
	return -p.Since(t)
}

// NewNowTimeProvider constructs a new NowTimeProvider initialised to return the
// current (changing) time
func NewNowTimeProvider() *NowTimeProvider {
	return &NowTimeProvider{}
}
