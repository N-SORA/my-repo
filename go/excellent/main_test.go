package main

import "testing"

// TestEvenOrOdd tests the EvenOrOdd function.
func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "even" {
		t.Errorf("expected: even, actual: %s", result)
	}
}