package HelloWorld

import (
	"fmt"
	"io"
	"time"
)

const (
	write = "write"
	sleep = "sleep"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct{}

type CountdownOperationsSpy struct {
	Calls []string
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Spies are a kind of mock which can record how a dependency is used.
// They can record the arguments sent in, how many times it has been called, etc.
// In our case, we're keeping track of how many times Sleep() is called so we can check it in our test.
func (s *SpySleeper) Sleep() {
	s.Calls++
}

const (
	countdownStart = 3
	finalWord      = "Go!"
)

func CountDown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	//for i := countdownStart; i > 0; i-- {
	//	fmt.Fprintln(out, i)
	//}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
