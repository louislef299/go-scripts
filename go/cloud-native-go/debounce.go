package main

import (
	"context"
	"sync"
	"time"
)

// Create a circuit type that will be used for our running function
type Circuit func(context.Context) (string, error)

// Debounce
func DebounceFirst(circuit Circuit, d time.Duration) Circuit {
	var threshold time.Time
	var result string
	var err error
	var m sync.Mutex

	return func(ctx context.Context) (string, error) {
		m.Lock()

		defer func() {
			// This means that the function was called with a result
			// returned, so I guess we extend the threshold time?
			threshold = time.Now().Add(d)
			m.Unlock()
		}()

		// Time before threshold, so return cached result
		if time.Now().Before(threshold) {
			return result, err
		}

		// Time after threshold, so return new result
		result, err = circuit(ctx)
		return result, err
	}
}
