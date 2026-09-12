package store

import (
	"errors"
)

// ErrEmptyKey sentinel error -> global error variable exported at the package level that represents a specific error condition
var (
	ErrEmptyKey  = errors.New("key is mandatory")
	ErrStoreFull = errors.New("store is full")
)
