package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hour, minute int
}

func New(h, m int) Clock {
	return Clock{hour: 0, minute: 0}.Add(h*60 + m)
}

func (c Clock) Add(m int) Clock {
	minutes := c.hour*60 + c.minute + m
	if minutes >= 0 {
		c.minute = minutes % 60
		c.hour = (minutes / 60) % 24
	} else {
		c.minute = 60 + (minutes % 60)
		if c.minute == 60 {
			c.minute = 0
		}
		minutes -= c.minute
		c.hour = (minutes/60)%24 + 24
	}
	if c.hour == 24 {
		c.hour = 0
	}
	return Clock{c.hour, c.minute}
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02v:%02v", c.hour, c.minute)
}
