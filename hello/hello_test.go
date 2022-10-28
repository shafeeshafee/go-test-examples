package hello

import "testing"

func TestHello(t *testing.T) {

	t.Run("Saying hello to user", func(t *testing.T) {
		got := Hello("Shafee", "")
		want := "Hello, Shafee"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when an argument is empty string", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	// in spanish
	t.Run("Say hello to user in Spanish", func(t *testing.T) {
		got := Hello("Carla", "Spanish")
		want := "Hola, Carla"
		assertCorrectMessage(t, got, want)
	})

	// in french
	t.Run("Saying hello to user in French", func(t *testing.T) {
		got := Hello("Polnaref", "French")
		want := "Bonjour, Polnaref"
		assertCorrectMessage(t, got, want)
	})

	// in japanese
	t.Run("Saying hello to user in Japanese", func(t *testing.T) {
		got := Hello("Sasago", "Japanese")
		want := "こんにちは, Sasago"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// tell the test suite that this method is a helper
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
