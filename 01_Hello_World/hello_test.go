package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	// définition d'une méthode helper pour le test
	checkCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("should say hello to visitor", func(t *testing.T) {
		got := Hello("Jean")
		want := "Hello, Jean"
		checkCorrectMessage(t, got, want) // utilisation du helper
	})

	t.Run("should say hello word if no name send", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		checkCorrectMessage(t, got, want)
	})
}
