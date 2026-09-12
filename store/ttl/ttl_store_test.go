package ttl

import (
	"testing"
	"time"
)

func TestGet_CheckExpiredTime(t *testing.T) {
	ttl := NewTtlStore(50 * time.Millisecond)
	ttl2 := NewTtlStore(120 * time.Millisecond)

	ttl.Set("a", "1")
	ttl2.Set("b", "2")

	time.Sleep(100 * time.Millisecond)

	if _, err := ttl.Get("a"); err == nil {
		t.Error("expected error from Get(), returned nil")
	}

	if _, err := ttl2.Get("b"); err != nil {
		t.Error("expected nil from Get(), returned err")
	}
}
