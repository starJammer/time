package time

//UseWarpClock switches the package to use a warped clock.
//The warp parameter specifies how many seconds
//the clock will advance, or regress, for every second
//of real time that transpires.
//warp is the duration that will transpire for each
//reference duration.
func UseWarpClock(warp Duration, reference Duration) {
	//w := warpClock{warp, reference}
	//switchClock(w)
}

//distortedClock represents a clock that runs
//either faster than normal or slower than normal.
//It can also go back in time or be frozen in time.
type warpClock struct {
	warp Duration
	ref  Duration
}
