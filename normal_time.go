package time

import (
	"time"
)

//UseNormalTime switches the package to use the normal
//time package and start reporting the actual system
//time. Any go processes that are sleeping might be
//awakened if the time they expected to wake up is
//less than the new reference time set.
func UseNormalTime() {
	n := normalClock{}
	switchClock(n)
}

type normalClock struct{}

func (n normalClock) after(d Duration) <-chan Time {
	c := make(chan Time)
	o := time.After(d.d)
	go func() {
		c <- Time{<-o}
		close(c)
	}()

	return c
}

func (n normalClock) sleep(d Duration) {
	time.Sleep(d.d)
}

func (n normalClock) tick(d Duration) <-chan Time {
	c := make(chan Time)
	o := time.Tick(d.d)
	go func() {
		c <- Time{<-o}
		close(c)
	}()

	return c
}

func (n normalClock) now() Time {
	return Time{time.Now()}
}

func (n normalClock) init() {
	if fans != nil {
		//cleanup any remaining threads
	}
}

func (n normalClock) cleanup() {}

func (n normalClock) timeSwitch() {}
