package store

import (
	"errors"
	"pmg/pass"
)

// Store is a store for all token
type Store interface {
	// Set save token by key, error when failed
	Set(key, passwd string, values ...any) error

	// Get get token by key
	Get(key string) (*pass.Token, error)

	// Persistence save all data somewhere permanently
	Persistence() error
}

var (
	ErrKeyNotFound = errors.New("")

	ErrPersistenceFailed = errors.New("")
)
