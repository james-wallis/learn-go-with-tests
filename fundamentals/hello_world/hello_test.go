package main

import "testing"

// Writing a test is just like writing a function, with a few rules
// * It needs to be in a file with a name like xxx_test.go
// * The test function must start with the word Test
// * The test function takes one argument only t *testing.T
func TestHello(t *testing.T) {
	got := Hello("James")
	want := "Hello, James"

	if got != want {
		// %q is useful as it wraps the value in double quotes
		t.Errorf("got %q want %q", got, want)
	}
}
