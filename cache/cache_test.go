package cache

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	cache := NewCache()
	if cache == nil {
		t.Fatal("cache can't be nil")
	}

	key := "apple"
	value := []byte("apple's value")
	cache.Add(key, value)

	val, ok := cache.Get(key)
	if !ok {
		t.Fatalf("did not get the value for the key %v", key)
	}

	if !reflect.DeepEqual(value, val) {
		t.Fatalf("expected: %v, got: %v", value, val)
	}
}

func TestExistingKeyValueUpdated(t *testing.T) {
	cache := NewCache()
	if cache == nil {
		t.Fatal("cache can't be nil")
	}
	key := "apple"
	existingValue := []byte("apple's value")
	cache.Add(key, existingValue)

	newValue := []byte("appls's new value")
	cache.Add(key, newValue)

	val, ok := cache.Get(key)
	if !ok {
		t.Fatalf("did not get the value for the key %v", key)
	}

	if !reflect.DeepEqual(newValue, val) {
		t.Fatalf("expected: %v, got: %v", newValue, val)
	}
}
