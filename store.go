package main

import (
	"errors"
	"fmt"
	"sort"
)

// sentinel error -> global error variable exported at the package level that represents a specific error condition
var ErrEmptyKey = errors.New("Key is mandatory")
var ErrStoreFull = errors.New("Store is full")

type Store struct {
	data    map[string]string
	maxSize int // 0 -> unlimited size
}

func NewStore(maxSize int) *Store {
	if maxSize == 0 {
		return &Store{
			data:    make(map[string]string),
			maxSize: maxSize,
		}
	}
	return &Store{
		data:    make(map[string]string, maxSize),
		maxSize: maxSize,
	}
}

func (s *Store) Get(key string) (string, error) {
	if key == "" {
		return "", ErrEmptyKey
	}

	value, ok := s.data[key]
	if !ok {
		return "", fmt.Errorf("Key %s does not exists!", key)
	}

	return value, nil
}

func (s *Store) Set(key, value string) error {
	if key == "" {
		return ErrEmptyKey
	}

	_, exists := s.data[key]
	if s.maxSize > 0 && len(s.data) >= s.maxSize && !exists {
		// %w -> specific for Errorf to wrap an error
		return fmt.Errorf("Set(%q): %w", key, ErrStoreFull)
	}

	s.data[key] = value
	return nil
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}

// return all keys currently in the store, sorted alphabetically
func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))

	for key := range s.data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

// Rename moves the value at oldKey over to newKey.
func (s *Store) Rename(oldKey, newKey string) {
	if val, ok := s.Get(oldKey); ok == nil {
		s.Set(newKey, val)
		s.Delete(oldKey)
	} else {
		fmt.Printf("message: %s", ok.Error())
	}
}

// Pop hands back the value at key and removes it in one go.
func (s *Store) Pop(key string) (string, bool) {
	val, ok := s.Get(key)
	if ok != nil {
		fmt.Printf("\nmessage: %s", ok.Error())
		return "", false
	}
	s.Delete(key)
	// TODO: read s.data[key] with the two-value form, delete() it, return both
	return val, true
}
