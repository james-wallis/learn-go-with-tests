package main

import (
	"os"
	"time"

	"github.com/james-wallis/learn-go-with-tests/fundamentals/maths/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
