// Package gigasecond adds a gigasecond to the submitted time
package gigasecond

import "time"

// AddGigasecond adds a gigasecond a given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
