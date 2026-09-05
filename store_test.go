package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestKeys_ReturnAllKeysSorted(t *testing.T) {
	store := NewStore()
	store.Set("g", "1")
	store.Set("a", "2")
	store.Set("b", "3")
	got := store.Keys()
	expected := []string{"a", "b", "g"}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Keys() = %v, expected = %v", got, expected)
		return
	}
	fmt.Printf("Keys() = %v, expected = %v", got, expected)
}

func TestRename_MovesTheValue(t *testing.T) {
	// 1. NewStore(), then Set() one key.
	// 2. Rename() it to a new key.
	// 3. Assert Get() finds the value under the new key.
	// 4. Assert Get() no longer finds the old key.
	s := NewStore()
	s.Set("key1", "value1")
	s.Rename("key1", "keyOne")

	val, ok := s.Get("keyOne")

	if val != "value1" || !ok {
		t.Errorf("The value is incorrect.: %s. Should be value1", val)
	}

	val, ok = s.Get("key1")
	if ok || val == "value1" {
		t.Error("key1 should not exists anymore")
	}

}

func TestPop_ReturnsAndRemoves(t *testing.T) {
	store := NewStore()
	store.Set("keyOne", "a")
	val, ok := store.Pop("keyOne")

	t.Logf("%s, %v", val, ok)

	if val != "a" || !ok {
		t.Error("The popped result should be [a] and [true]")
	}

	val, ok = store.Pop("keyOne")
	if ok {
		t.Error("The popped result should be [empty string] and [false]")
	}
}

func TestRename_MissingKeyCreatesNothing(t *testing.T) {
	s := NewStore()
	s.Set("key1", "a")
	s.Rename("key2", "keyTwo")

	if _, ok := s.Get("keyTwo"); ok {
		t.Error("Get() should not find the keyTwo key")
	}

	if keys := s.Keys(); len(keys) != 1 {
		t.Error("Store should hold just one key")
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name      string
		setup     map[string]string
		deleteKey string
		wantLen   int
	}{
		{"deletes existing key", map[string]string{"a": "1", "b": "2"}, "a", 1},
		{"no-op on missing key", map[string]string{"a": "1"}, "x", 1},
		{"no-op on empty store", map[string]string{}, "x", 0},
		{"deletes last remaining key", map[string]string{"only": "value"}, "only", 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			store := NewStore()
			for k, v := range tc.setup {
				store.Set(k, v)
			}

			store.Delete(tc.deleteKey)

			if got := len(store.Keys()); got != tc.wantLen {
				t.Errorf("after Delete(%q): Len() = %d, want %d", tc.deleteKey, got, tc.wantLen)
			}

			if _, ok := store.Get(tc.deleteKey); ok {
				t.Errorf("after Delete(%q): key still present", tc.deleteKey)
			}
		})
	}
}
