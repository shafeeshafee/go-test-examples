package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Shafee")
	want := "Hello, Shafee!"

	if got != want {
		t.Errorf("Got %q. Wanted %q.", got, want)
	}
}
