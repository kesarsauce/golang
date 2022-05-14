package main

import "testing"

func TestHello(t *testing.T) {
	assertGotWant := func(got, want string, t testing.TB) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to a name", func(t *testing.T){
		got := Hello("Sam", "")
		want := "Hello, Sam"

		assertGotWant(got, want, t)
	})
	t.Run("no name passed", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"

		assertGotWant(got, want, t)
	})
	t.Run("in spanish", func(t *testing.T){
		got := Hello("Sam", "spanish")
		want := "Hola, Sam"

		assertGotWant(got, want, t)
	})
	t.Run("in french", func(t *testing.T){
		got := Hello("Sam", "french")
		want := "Bonjour, Sam"

		assertGotWant(got, want, t)
	})
	t.Run("in hindi", func(t *testing.T){
		got := Hello("Sam", "hindi")
		want := "Namaskar, Sam"

		assertGotWant(got, want, t)
	})
	
}

