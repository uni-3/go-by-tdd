package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

type Sleeper interface {
	Sleep()
}

/*
type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
*/

// mock
/*
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
*/

type CountdownOperationSpy struct {
	Calls []string
}

const (
	write = "write"
	sleep = "sleep"
)

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (int, error) {
	s.Calls = append(s.Calls, write)
	return 0, nil
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// SleepのDI
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)

}
