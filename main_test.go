package main

import "testing"

func TestAdd(t *testing.T) {
	t.Run("Add 1 and 2", func(t *testing.T) {
		result, err := add(1, 2)
		if err != nil {
			t.Fatal(err)
		}
		if result != 3 {
			t.Fatal("1 + 2 did not equal 3")
		}
	})
	t.Run("Add -1 and 2", func(t *testing.T) {
		result, err := add(-1, 2)
		if err == nil {
			t.Fatal("Expected an error")
		}
		if result != 0 {
			t.Fatal("Expected 0")
		}
	})
	t.Run("Add 1 and -2", func(t *testing.T) {
		result, err := add(1, -2)
		if err == nil {
			t.Fatal("Expected an error")
		}
		if result != 0 {
			t.Fatal("Expected 0")
		}
	})
	t.Run("Add -1 and -2", func(t *testing.T) {
		result, err := add(-1, -2)
		if err == nil {
			t.Fatal("Expected an error")
		}
		if result != 0 {
			t.Fatal("Expected 0")
		}
	})
}
