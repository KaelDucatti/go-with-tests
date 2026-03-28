package countdown

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++	
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (c *ConfigurableSleeper) Sleep() {
	time.Sleep(c.duration)
}

func NewConfigurableSleeper(duration time.Duration) *ConfigurableSleeper {
	return &ConfigurableSleeper{duration: duration}
}

const (
	Write = "write"
	Sleep = "sleep"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (sco *SpyCountdownOperations) Sleep() {
	sco.Calls = append(sco.Calls, Sleep)
}

func (sco *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	sco.Calls = append(sco.Calls, Write)
	return
}

func Countdown(w io.Writer, ss Sleeper) {
	for i := 3; i >= 1; i-- {
		fmt.Fprintln(w, i)
		ss.Sleep()
	}
	fmt.Fprintln(w, "Go")
}
