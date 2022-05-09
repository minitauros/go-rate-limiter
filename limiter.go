package rate

import (
	"math"
	"time"
)

// Limiter is a rate limiter.
type Limiter struct {
	tokenCh chan struct{}
	stopCh  chan struct{}
}

// NewLimiter returns a new rate limiter.
func NewLimiter(allowedCalls int64, per time.Duration) *Limiter {
	l := &Limiter{
		tokenCh: make(chan struct{}, 1),
		stopCh:  make(chan struct{}),
	}

	go func() {
		replenishInterval := time.Duration(math.Round(float64(per) / float64(allowedCalls)))
		ticker := time.NewTicker(replenishInterval)
		for {
			select {
			case <-l.stopCh:
				return
			case <-ticker.C:
				if len(l.tokenCh) == 1 {
					break
				}
				l.tokenCh <- struct{}{}
			}
		}
	}()

	return l
}

// Take takes one token from the rate limiter. It blocks until it can.
func (l *Limiter) Take() {
	<-l.tokenCh
}

// Stop stops the rate limiter.
func (l *Limiter) Stop() {
	close(l.stopCh)
}
