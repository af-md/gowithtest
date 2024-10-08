package main

import (
	"fmt"
	"io"
	"time"
)

// func main() {
// 	sleeper := &DefaultSleeper{}
// 	Countdown(os.Stdout, sleeper)
// }

type SpyCountdownOperations struct {
	Calls     []string
	SleepTime int
}

func (s *SpyCountdownOperations) Sleep() {
	for i := 0; i < s.SleepTime; i++ {
		s.Calls = append(s.Calls, sleep)
	}
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls     int
	SleepTime int
}

type DefaultSleeper struct {
	SleepTime int
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(time.Duration(d.SleepTime) * time.Second)
}

const finalWord = "Go!"
const countdownStart = 3

func (s *SpySleeper) Sleep() {
	for i := 0; i <= s.SleepTime; i++ {
		s.Calls++
	}
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}
