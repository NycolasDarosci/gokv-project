package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestKeys_ReturnAllKeysSorted(t *testing.T) {
	store := NewStore(0)
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
	s := NewStore(0)
	s.Set("key1", "value1")
	s.Rename("key1", "keyOne")

	val, ok := s.Get("keyOne")

	if val != "value1" || ok != nil {
		t.Errorf("The value is incorrect.: %s. Should be value1", val)
	}

	val, ok = s.Get("key1")
	if ok == nil || val == "value1" {
		t.Error("key1 should not exists anymore")
	}

}

func TestPop_ReturnsAndRemoves(t *testing.T) {
	store := NewStore(0)
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
	s := NewStore(0)
	s.Set("key1", "a")
	s.Rename("key2", "keyTwo")

	if _, ok := s.Get("keyTwo"); ok == nil {
		t.Error("Get() should not find the keyTwo key")
	}

	if keys := s.Keys(); len(keys) != 1 {
		t.Error("Store should hold just one key")
	}
}

func TestSetGet_EmptyKey(t *testing.T) {
	s := NewStore(0)

	if _, ok := s.Get("key1"); ok == nil {
		t.Error("Get() should not find a key which does not exists")
	}

	if err := s.Set("", "value1"); !errors.Is(err, ErrEmptyKey) {
		t.Error("Set() should not set a key-value pair which key is empty")
	}
}

func TestSet_CheckMaxSize(t *testing.T) {
	s := NewStore(2)
	_ = s.Set("a", "1")
	_ = s.Set("b", "2")

	if err := s.Set("c", "3"); err == nil {
		t.Errorf("Set() should have thrown an error for setting a new key. Error: %s", err.Error())
	}

	if len(s.data) != 2 {
		t.Error("Set() should have had just 2 keys")
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
			store := NewStore(0)
			for k, v := range tc.setup {
				store.Set(k, v)
			}

			store.Delete(tc.deleteKey)

			if got := len(store.Keys()); got != tc.wantLen {
				t.Errorf("after Delete(%q): Len() = %d, want %d", tc.deleteKey, got, tc.wantLen)
			}

			if _, ok := store.Get(tc.deleteKey); ok == nil {
				t.Errorf("after Delete(%q): key still present", tc.deleteKey)
			}
		})
	}
}
