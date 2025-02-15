package main

import "testing"

func TestAdd(t *testing.T) {
	result := Sub(1, 2)

	expected := -1

	if result != expected {
		t.Errorf("Nahi chala %d", result)
	}
}
