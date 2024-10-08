package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloWorld("Afzal", "")
		want := "Hello, Afzal"
		assertMessage(got, want, t)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := HelloWorld("", "")
		want := "Hello, world"
		assertMessage(got, want, t)
	})

	t.Run("say hello to people in spanish", func(t *testing.T) {
		got := HelloWorld("Afzal", "Spanish")
		want := "Hola, Afzal"
		assertMessage(got, want, t)
	})
}

func assertMessage(got, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("Test failed, got %q and want %q", got, want)
	}
}
