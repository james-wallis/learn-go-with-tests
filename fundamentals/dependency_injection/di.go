package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet : Writes a string to a given io.Writer (os.Stdout, a buffer)
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler : handles requests for a HTTP server
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
