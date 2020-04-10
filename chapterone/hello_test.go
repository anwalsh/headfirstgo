package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := Hello()
	want := "Hello, world!"

	if got != want {
		t.Errorf("Received %q want %q", got, want)
	}
}
