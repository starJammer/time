package time

import (
	"errors"
	"sync"
)

var (
	lock *sync.Mutex
	//threads represents the goroutines that have called
	//the package and are waiting on some time event
	fans []chan Time

	//a channel used to fan out time change events
	//to the threads
	fanner chan Time
)

//UseWarpClock switches the package to use a warped clock.
//The warp parameter specifies how many seconds
//the clock will advance, or regress, for every second
//of real time that transpires.
//warp is the duration that will transpire for each
//reference duration.
func UseWarpClock(warp Duration, reference Duration) error {

	if int64(reference.d) < 0 {
		return errors.New("Reference time must be positive")
	}

	if fans == nil {
		lock = new(sync.Mutex)
		fans = make([]chan Time, 0)
		fanner = make(chan Time, 0)
	}

	switch v := int64(warp.d); {

	case v == 0:
		switchClock(frozenClock{})

	case v > 0:
		switchClock(forwardClock{warp, reference})

	case v < 0:
		switchClock(backwardClock{warp, reference})
	}

	return nil
}

//distortedClock represents a clock that runs
//either faster than normal or slower than normal.
//It can also go back in time or be frozen in time.
type forwardClock struct {
	warp Duration
	ref  Duration
}

func (f forwardClock) init() {
}

func (f forwardClock) timeSwitch() {

}

func (f forwardClock) after(d Duration) <-chan Time {
	return nil
}

func (f forwardClock) sleep(d Duration) {
}

func (f forwardClock) tick(d Duration) <-chan Time {
	return nil
}

func (f forwardClock) now() Time {
	return Time{}
}

func (f forwardClock) cleanup() {
}

type backwardClock struct {
	warp Duration
	ref  Duration
}

func (f backwardClock) init() {

}

func (f backwardClock) timeSwitch() {

}

func (f backwardClock) after(d Duration) <-chan Time {
	return nil
}

func (f backwardClock) sleep(d Duration) {
}

func (f backwardClock) tick(d Duration) <-chan Time {
	return nil
}

func (f backwardClock) now() Time {
	return Time{}
}

func (f backwardClock) cleanup() {
}

type frozenClock struct {
}

func (f frozenClock) init() {
}

func (f frozenClock) timeSwitch() {
}

func (f frozenClock) after(d Duration) <-chan Time {
	c := make(chan Time)
	lock.Lock()
	me := make(chan Time, 0)
	ref := referenceTime
	fans = append(fans, me)
	lock.Unlock()

	go func() {
		select
		t := <-me

	}()
	return c
}

func (f frozenClock) sleep(d Duration) {
}

func (f frozenClock) tick(d Duration) <-chan Time {
	return nil
}

func (f frozenClock) now() Time {
	return referenceTime
}

func (f frozenClock) cleanup() {
}
