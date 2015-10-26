package time

import (
	ti "time"
)

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
			loc.Location,
		),
	}
}

type Location struct {
	*ti.Location
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

var myClock clock

type clock interface {
	after(d Duration) <-chan Time
	sleep(d Duration)
	tick(d Duration) <-chan Time
	now() Time
	cleanup()
}

func switchClock(c clock) {
	if myClock != nil {
		myClock.cleanup()
	}
}

func After(d Duration) <-chan Time {

	return nil
}

func Sleep(d Duration) {

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
	t, err := ti.ParseInLocation(layout, value, loc.Location)
	return Time{t}, err
}

func Unix(sec int64, nsec int64) Time {
	return Time{ti.Unix(sec, nsec)}
}

type Duration struct {
	ti.Duration
}

type Ticker struct {
	C <-chan Time
}

func NewTicker(d Duration) *Ticker {

	return nil
}

func (*Ticker) Stop() {

}

//SetReferenceTime sets the time of reference for the package.
//Use this when you want your entire system to think that it
//is a very specific time.
//Any go process that are sleeping might be awakened
//if the time they expected to wake up is less than the
//new reference time seti.
func SetReferenceTime(t ti.Time) {

}

func Tick(d Duration) <-chan Time {

	return nil
}

type ParseError struct {
	*ti.ParseError
}

type Time struct {
	ti.Time
}

func (t Time) Add(d Duration) Time {
	return Time{t.Time.Add(d.Duration)}
}

func (t Time) AddDate(years int, months int, days int) Time {
	return Time{t.Time.AddDate(years, months, days)}
}

func (t Time) After(u Time) bool {
	return t.Time.After(u.Time)
}

func (t Time) Before(u Time) bool {
	return t.Time.Before(u.Time)
}

func (t Time) Equal(u Time) bool {
	return t.Time.Equal(u.Time)
}

func (t Time) In(loc *Location) Time {
	return Time{t.Time.In(loc.Location)}
}

func (t Time) Local() Time {
	return Time{t.Time.Local()}
}

func (t Time) Location() *Location {
	return &Location{t.Time.Location()}
}

func (t Time) Month() Month {
	return Month(t.Time.Month())
}

func (t Time) Round(d Duration) Time {
	return Time{t.Time.Round(d.Duration)}
}

func (t Time) Sub(u Time) Duration {
	return Duration{t.Time.Sub(u.Time)}
}

func (t Time) Truncate(d Duration) Time {
	return Time{t.Time.Truncate(d.Duration)}
}

func (t Time) UTC() Time {
	return Time{t.Time.UTC()}
}

func (t Time) Weekday() Weekday {
	return Weekday(t.Time.Weekday())
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

func Now() Time {

	return Time{}
}

type Ticker interface {
	Stop()
}
