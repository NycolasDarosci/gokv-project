package store

import (
	"encoding/base64"
	"errors"
)

// ErrEmptyKey sentinel error -> global error variable exported at the package level that represents a specific error condition
var (
	ErrEmptyKey  = errors.New("key is mandatory")
	ErrStoreFull = errors.New("store is full")
)

type Storer interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string)
	Keys() []string
	Rename(oldKey, newKey string)
}

func SetWithEncryption(s Storer, key, value string) (string, error) {
	encrypted := base64.StdEncoding.EncodeToString([]byte(value))
	if err := s.Set(key, encrypted); err != nil {
		return "", err
	}
	return s.Get(key)
}
