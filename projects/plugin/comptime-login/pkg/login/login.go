package login

import (
	"context"
	"fmt"
	"maps"
	"slices"
	"sync"
)

var (
	driversMu sync.RWMutex
	drivers   = make(map[string]Login)
)

// Register makes a login driver available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, driver Login) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		panic("sql: Register driver is nil")
	}
	if _, dup := drivers[name]; dup {
		panic("sql: Register called twice for driver " + name)
	}
	drivers[name] = driver
}

// Drivers returns a sorted list of the names of the registered drivers.
func Drivers() []string {
	driversMu.RLock()
	defer driversMu.RUnlock()
	return slices.Sorted(maps.Keys(drivers))
}

type ConfigOptions any
type ConfigOptionsFunc func(*ConfigOptions) error
type Login interface {
	Login(ctx context.Context, opts ...ConfigOptionsFunc) error
}
type Credentials struct{}

func DLogin(driverName string) (*Credentials, error) {
	driversMu.RLock()
	driverLogin, ok := drivers[driverName]
	driversMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("sql: unknown driver %q (forgotten import?)", driverName)
	}

	err := driverLogin.Login(context.TODO())
	if err != nil {
		return nil, err
	}
	return &Credentials{}, nil
}
