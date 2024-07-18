package main

import (
	"testing"
)

// concatstring funcion
func TestConcatString(t *testing.T) {
	//test1
	result := ConcatString("Ishu, ", "Padsala")
	expected := "Ishu, Padsala"
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}

	//test2
	result = ConcatString("", "")
	expected = ""
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}
