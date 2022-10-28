package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Saying hello to user", func(t *testing.T) {
		got := Hello("Shafee")
		want := "Hello, Shafee"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Say 'Hello, World' when an argument is empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// tell the test suite that this method is a helper
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
