package clock

import "time"

type systemClock struct{}

func NewSystemClock() *systemClock {
	return &systemClock{}
}

func (c *systemClock) Now() time.Time {
	return time.Now().UTC().Add(3 * time.Hour)
}
