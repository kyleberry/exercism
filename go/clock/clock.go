// Package clock does things
package clock

import (
	"fmt"
)

// Clock creates a struct for our clock
type Clock struct {
	hour   int
	minute int
}

const (
	hoursInADay     = 24
	minutesInAnHour = 60
)

// New creates a new clock for us, and handles negative hours/minutes to return
// a real time.
func New(hour, minute int) Clock {
	newMinute := minute % minutesInAnHour
	// If we're passed a negative value for `minute` we add the negative number
	// to 60 and decrement `hour` once.
	if newMinute < 0 {
		newMinute = minutesInAnHour + newMinute
		hour--
	}
	newHour := (hour + (minute / minutesInAnHour)) % hoursInADay
	// We do the same here if we have a negative `hour`, adding the negative
	// number to 24
	if newHour < 0 {
		newHour = hoursInADay + newHour
	}
	return Clock{newHour, newMinute}
}

// String returns a fancy string of our clock
func (c *Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds minutes to our clock and creates a new one calling New()
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

// Subtract subtracts minutes to our clock and creates a new one calling New()
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute-minutes)
}
