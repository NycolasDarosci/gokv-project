package ttl

import (
	"fmt"
	"redis-kvgo/store"
	"sort"
	"time"
)

// TTL -> time to live
// defined time how long an entry/value/data stored sonewhere can exist in fast accessed memory (cache) before it expires

type TtlStore struct {
	data map[string]ttlEntry
	ttl  time.Duration
}

type ttlEntry struct {
	value     string
	expiresAt time.Time
}

func NewTtlStore(ttl time.Duration) *TtlStore {
	return &TtlStore{
		data: make(map[string]ttlEntry),
		ttl:  ttl,
	}
}

func (s *TtlStore) Get(key string) (string, error) {
	if key == "" {
		return "", store.ErrEmptyKey
	}

	entry, exists := s.data[key]
	if !exists || time.Now().After(entry.expiresAt) {
		// lazy eviction -> no background job scanning for expired keys
		// deleted on demand
		s.Delete(key)
		return "", fmt.Errorf("Key %s does not exists!", key)
	}

	return entry.value, nil
}

func (t *TtlStore) Set(key, value string) error {
	if key == "" {
		return store.ErrEmptyKey
	}

	// entry, exists := s.data[key]
	// if !exists {
	// 	// if s.maxSize > 0 && len(s.data) >= s.maxSize && !exists {
	// 	// %w -> specific for Errorf to wrap an error
	// 	return fmt.Errorf("Set(%q): %w", key, ErrStoreFull)
	// }

	t.data[key] = ttlEntry{
		value:     value,
		expiresAt: time.Now().Add(t.ttl),
	}
	return nil
}

func (s *TtlStore) Delete(key string) { delete(s.data, key) }

// return all keys currently in the store, sorted alphabetically
func (s *TtlStore) Keys() []string {
	keys := make([]string, 0, len(s.data))

	for key := range s.data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

// Rename moves the value at oldKey over to newKey.
func (t *TtlStore) Rename(oldKey, newKey string) {
	if val, ok := t.Get(oldKey); ok == nil {
		t.Set(newKey, val)
		t.Delete(oldKey)
	} else {
		fmt.Printf("message: %s", ok.Error())
	}
}
