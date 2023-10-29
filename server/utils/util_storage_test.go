package utils

import (
	"strconv"
	"testing"
)

func TestStorage(t *testing.T) {
	storage := NewStorage()

	// Test case 1: Add an item
	storage.Add("item1")
	if !storage.Check("item1") {
		t.Error("Check failed for an existing item")
	}
	if storage.Check("item2") {
		t.Error("Check passed for a non-existing item")
	}

	// Test case 2: Remove an item
	storage.Remove("item1")
	if storage.Check("item1") {
		t.Error("Check passed for a removed item")
	}

	// Test case 3: Remove a non-existing item
	storage.Remove("item2")
	if storage.Check("item2") {
		t.Error("Check passed for a non-existing item")
	}
}

func TestConcurrentStorage(t *testing.T) {
	// Test case: Add and remove items concurrently
	storage := NewStorage()

	const numGoroutines = 100
	const itemsPerGoroutine = 100
	done := make(chan struct{})

	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			for j := 0; j < itemsPerGoroutine; j++ {
				key := "item" + strconv.Itoa(i*itemsPerGoroutine+j)
				storage.Add(key)
				// Make sure the item exists
				if !storage.Check(key) {
					t.Error("Check failed for an existing item in a concurrent test")
				}
				// Remove the item
				storage.Remove(key)
				// Make sure the item doesn't exist after removal
				if storage.Check(key) {
					t.Error("Check passed for a removed item in a concurrent test")
				}
			}
			done <- struct{}{}
		}(i)
	}

	// Wait for all goroutines to finish
	for i := 0; i < numGoroutines; i++ {
		<-done
	}

	// Ensure that the storage is empty after all goroutines finish
	if len(storage.storage) != 0 {
		t.Errorf("Storage is not empty after concurrent operations")
	}
}
