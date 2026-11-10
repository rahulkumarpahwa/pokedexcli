package cache

import (
	"reflect"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	key := "apple"
	cache := NewCache(time.Second * 5)
	if cache == nil {
		t.Fatal("cache can't be nil")
	}

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
	key := "apple"
	cache := NewCache(time.Second * 5)
	if cache == nil {
		t.Fatal("cache can't be nil")
	}
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

func TestReapLoopWithGettingAfterInterval(t *testing.T) {
	interval := 100 * time.Millisecond
	cache := NewCache(interval)
	key := "apple"
	value := []byte("apple's value")
	cache.Add(key, value)

	deadline := time.Now().Add(time.Second) // checking for the within a timeline of 1 second as it is not sure to be removed after the first interval and may get removed in second or next interval. so we will ping it to get after every 10 ms that it exist or not. even though we have given to get deleted after 100 ms but will check for 1 second atleast bcoz we aare not sure when it gets deleted in first try or second.

	for time.Now().Before(deadline) {
		_, ok := cache.Get(key)

		if !ok {
			return
		}

		time.Sleep(10 * time.Millisecond)
	}

	t.Fatalf("expected key %v to be removed after expiration", key)
}
