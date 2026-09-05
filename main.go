package main

//kvgo -> key value go project
import (
	"fmt"
	"sort"
)

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Get(key string) (string, bool) {
	value, ok := s.data[key]
	return value, ok
}

func (s *Store) Set(key, value string) {
	s.data[key] = value
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
	// TODO: move the value at oldKey over to newKey, then delete() oldKey
	if val, ok := s.Get(oldKey); ok {
		s.Set(newKey, val)
		s.Delete(oldKey)
	}

}

// Pop hands back the value at key and removes it in one go.
func (s *Store) Pop(key string) (string, bool) {
	val, ok := s.Get(key)
	s.Delete(key)
	// TODO: read s.data[key] with the two-value form, delete() it, return both
	return val, ok
}

func main() {
	fmt.Println("Gokv project")

	store := NewStore()
	fmt.Println(store)  // pointer to a store -> address
	fmt.Println(*store) // dereference -> value
	fmt.Println(&store) // address itself

	store.Set("name", "Ricardo")
	fmt.Println(*store) // dereference -> value

	store.Set("lastName", "Albe")
	fmt.Println(*store) // dereference -> value

	value, returned := store.Get("nam")
	fmt.Println(value, returned)

	fmt.Println(store.data)
	store.Delete("name")
	fmt.Println(store.data)
	store.Delete("nam")
	fmt.Println(store.data)

	fmt.Println(0x62)
}
