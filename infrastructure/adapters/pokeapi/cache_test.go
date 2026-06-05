package pokeapi

import (
	"testing"
	"time"
)

func TestCacheAddAndGet(t *testing.T) {
	cache := NewCache(5 * time.Minute)
	key := "test-key"
	val := []byte("test-value")

	cache.Add(key, val)

	got, ok := cache.Get(key)

	if !ok {
		t.Errorf("expected value to be in cache")
	}
	if string(got) != "test-value" {
		t.Errorf("expected value to be 'test-value', got %s", got)
	}
}
