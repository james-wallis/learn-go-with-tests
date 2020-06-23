package main

import "fmt"

// Using constants should improve performance because
// it saves us creating the "Hello, " string instance everytime
// the Hello function is called
const helloPrefix = "Hello, "

// Hello : A simple function that returns a String
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}

func main() {
	// Separate Hello() and Println() to make Hello() easier to test
	fmt.Println(Hello("world"))
}
