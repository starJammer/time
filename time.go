package time

import (
	t "time"
)

//NormalTime switches the package to use the normal
//time package and start reporting the actual system
//time. Any go processes that are sleeping might be
//awakened if the time they expected to wake up is
//less than the new reference time set.
func NormalTime() {

}

func FreezeClock() {

}

//Switches the package to use a distorted clock.
//The distortion parameter specifies how many seconds
//the clock will advance, or regress, for every second
//of real time that transpires.
func DistortClock(distortion interface{}) {

}

//SetReferenceTime sets the time of reference for the package.
//Use this when you want your entire system to think that it
//is a very specific time.
//Any go process that are sleeping might be awakened
//if the time they expected to wake up is less than the
//new reference time set.
func SetReferenceTime(t t.Time) {

}

//freezeClock represents a clock that is frozen
//in time
type freezeClock struct{}

//distortedClock represents a clock that runs
//either faster than normal or slower than normal.
//It can also go back in time
type distortedClock struct{}
