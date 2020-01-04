package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("wingsico")
	want := "Hello, wingsico"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
