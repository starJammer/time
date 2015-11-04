package time

import (
	ti "time"
)

func init() {
	myClock = normalClock{}
	referenceTime = Time{ti.Now()}
}

/******************\
	Completed
\******************/

// A Month specifies a month of the year (January = 1, ...).
type Month int

// String returns the English name of the month ("January", "February", ...).
func (m Month) String() string { return months[m-1] }

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

var months = [...]string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time {
	return Time{
		ti.Date(
			year,
			ti.Month(month),
			day,
			hour,
			min,
			sec,
			nsec,
			loc.l,
		),
	}
}

type Location struct {
	l *ti.Location
}

func FixedZone(name string, offset int) *Location {
	return &Location{ti.FixedZone(name, offset)}
}

func LoadLocation(name string) (*Location, error) {
	l, err := ti.LoadLocation(name)
	return &Location{l}, err
}

/******************\
	Todo
\******************/

var (
	myClock       clock
	referenceTime Time
)

type clock interface {
	after(d Duration) <-chan Time
	sleep(d Duration)
	tick(d Duration) <-chan Time
	now() Time
	//init is called right after the clock is switched
	init()
	//cleanup is called before the clock is switched
	//out for another one
	cleanup()
	//timeSwitch notifies the clock of the reference time changing
	timeSwitch()
}

func switchClock(c clock) {

	if myClock != nil {
		myClock.cleanup()
	}

	myClock = c
	myClock.init()
}

//setReferenceTime sets the reference time to be
//used by the distorted clocks. If you use the normal
//clock this won't have any effect
func SetReferenceTime(t Time) {
	referenceTime = t
	myClock.timeSwitch()
}

func ReferenceTime() Time {
	return referenceTime
}

func After(d Duration) <-chan Time {
	return myClock.after(d)
}

func Sleep(d Duration) {
	myClock.sleep(d)
}

func Now() Time {
	return myClock.now()
}

func Tick(d Duration) <-chan Time {
	return myClock.tick(d)
}

func ParseDuration(s string) (Duration, error) {
	d, err := ti.ParseDuration(s)
	return Duration{d}, err
}

func Since(t Time) Duration {
	return Now().Sub(t)
}

func Parse(layout, value string) (Time, error) {
	t, err := ti.Parse(layout, value)
	return Time{t}, err
}

func ParseInLocation(layout, value string, loc *Location) (Time, error) {
	t, err := ti.ParseInLocation(layout, value, loc.l)
	return Time{t}, err
}

func Unix(sec int64, nsec int64) Time {
	return Time{ti.Unix(sec, nsec)}
}

type Duration struct {
	d ti.Duration
}

type Ticker struct {
	C <-chan Time
}

func NewTicker(d Duration) *Ticker {

	return nil
}

func (*Ticker) Stop() {

}

type ParseError struct {
	*ti.ParseError
}

type Time struct {
	t ti.Time
}

func (t Time) Add(d Duration) Time {
	return Time{t.t.Add(d.d)}
}

func (t Time) AddDate(years int, months int, days int) Time {
	return Time{t.t.AddDate(years, months, days)}
}

func (t Time) After(u Time) bool {
	return t.t.After(u.t)
}

func (t Time) Before(u Time) bool {
	return t.t.Before(u.t)
}

func (t Time) Equal(u Time) bool {
	return t.t.Equal(u.t)
}

func (t Time) In(loc *Location) Time {
	return Time{t.t.In(loc.l)}
}

func (t Time) Local() Time {
	return Time{t.t.Local()}
}

func (t Time) Location() *Location {
	return &Location{t.t.Location()}
}

func (t Time) Month() Month {
	return Month(t.t.Month())
}

func (t Time) Round(d Duration) Time {
	return Time{t.t.Round(d.d)}
}

func (t Time) Sub(u Time) Duration {
	return Duration{t.t.Sub(u.t)}
}

func (t Time) Truncate(d Duration) Time {
	return Time{t.t.Truncate(d.d)}
}

func (t Time) UTC() Time {
	return Time{t.t.UTC()}
}

func (t Time) Weekday() Weekday {
	return Weekday(t.t.Weekday())
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

var days = [...]string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

// String returns the English name of the day ("Sunday", "Monday", ...).
func (d Weekday) String() string { return days[d] }
