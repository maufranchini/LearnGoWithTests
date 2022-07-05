package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Mauro","English")
		want := "Hello, Mauro"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("","")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T){
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T){
		got := Hello("Tom", "French")
		want := "Bonjour, Tom"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T){
		got := Hello("Max", "German")
		want := "Hallo, Max"
		assertCorrectMessage(t, got, want)
	})
}
