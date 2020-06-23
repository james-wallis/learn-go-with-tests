package main

import "testing"

// Writing a test is just like writing a function, with a few rules
// * It needs to be in a file with a name like xxx_test.go
// * The test function must start with the word Test
// * The test function takes one argument only t *testing.T
func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		// t.Helper() is needed to tell the test suite that this method is a helper.
		// By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
		t.Helper()
		if got != want {
			// %q is useful as it wraps the value in double quotes
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("James", "")
		want := "Hello, James"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("James", "Spanish")
		want := "Hola, James"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in French", func(t *testing.T) {
		got := Hello("James", "French")
		want := "Bonjour, James"
		assertCorrectMessage(t, got, want)
	})
}
