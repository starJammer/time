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

func (n normalClock) now() Time {
	return Time{time.Now()}
}

func (n normalClock) sleep(d Duration) {
	time.Sleep(d.Duration)
}

func (n normalClock) cleanup() {

}
