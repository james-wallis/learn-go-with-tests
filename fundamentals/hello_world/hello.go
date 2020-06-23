package main

import "fmt"

// Using constants should improve performance because
// it saves us creating the "Hello, " string instance everytime
// the Hello function is called
const spanish = "Spanish"
const french = "French"
const helloEnglishPrefix = "Hello, "
const helloSpanishPrefix = "Hola, "
const helloFrenchPrefix = "Bonjour, "

// Hello : A simple function that returns a String
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	// prefix is a named return value
	// * It is created with a zero value
	// * Is returned by just using "return" rather than "return prefix"
	switch language {
	case french:
		prefix = helloFrenchPrefix
	case spanish:
		prefix = helloSpanishPrefix
	default:
		prefix = helloEnglishPrefix
	}
	return
}

func main() {
	// Separate Hello() and Println() to make Hello() easier to test
	fmt.Println(Hello("world", ""))
}
