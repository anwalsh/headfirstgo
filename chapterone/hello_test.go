package main

import "testing"

func TestHelloWorld(t *testing.T) {
	input := Hello()
	expected := "Hello, world!"

	if input != expected {
		t.Errorf("Received %q expected %q", input, expected)
	}
}
