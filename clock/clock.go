package clock

import "fmt"

const testVersion = 4

type Clock int

func New(hour, minute int) Clock {
	time := (hour*60 + minute) % (60 * 24)
	if time < 0 {
		time += 60 * 24
	}
	return Clock(time)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}