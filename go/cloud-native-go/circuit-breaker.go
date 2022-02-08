package main

import (
	"context"
	"errors"
	"sync"
	"time"
)

type Ciruit func(context.Context) (string, error)

// Uses uint to force a positive number? I can see potential errors with that...
func Breaker(circuit Ciruit, failureThreshold uint) Ciruit {
	var consecutiveFailures int = 0
	var lastAttempt = time.Now()
	var m sync.RWMutex

	// Closure
	return func(ctx context.Context) (string, error) {
		m.RLock() // Read lock

		d := consecutiveFailures - int(failureThreshold)

		// There are too many consecutive failures, so we should use our backoff logic
		if d >= 0 {
			shouldRetryAt := lastAttempt.Add(time.Second * 2 << d)
			// Block any requests that are to fast
			if !time.Now().After(shouldRetryAt) {
				m.RUnlock()
				return "", errors.New("service unreachable")
			}
		}

		m.RUnlock() // Release read lock

		response, err := circuit(ctx)

		m.Lock()
		defer m.Unlock()

		lastAttempt = time.Now()

		if err != nil {
			consecutiveFailures++
			return response, err
		}

		// Successful request! So reset failures ticker
		consecutiveFailures = 0
		return response, err
	}
}
