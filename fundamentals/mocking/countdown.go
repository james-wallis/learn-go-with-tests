package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type (
	// Sleeper : interface that defines our dependency to Sleep
	Sleeper interface {
		Sleep()
	}

	// ConfigurableSleeper : takes a duration of how long to sleep for and a function to call when Sleep() is called
	ConfigurableSleeper struct {
		duration time.Duration
		// sleep is a function, this allows us to pass in a Spy or time.Sleep
		sleep func(time.Duration)
	}
)

const finalWord = "Go!"
const countdownStart = 3

// Sleep : Calls the function that is stored as c.sleep with the value for c.duration
// When we run main this is time.Sleep()
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Countdown : Counts down from the value declared in countdownStart to 0 with a pause in between outputs
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{duration: 1 * time.Second, sleep: time.Sleep}
	Countdown(os.Stdout, sleeper)
}
