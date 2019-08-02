package clock

import "fmt"

type Clock struct {
	minute int
	hour   int
}

func timeConvert(h int, m int) (hour int, minute int) {

	if m < 0 {
		h = h - (1 + -m/60)
		m = 60 - -m%60
	}
	if h < 0 {
		h = (24 * (1 + -h/24)) + h
	}
	minute = m
	//	fmt.Printf("------entre [%d:%d]\n", h, m)
	for minute > 59 {
		hr := minute / 60
		h += hr
		minute -= 60 * hr
		//		fmt.Printf("------[%d:%d]\n", h, minute)
	}
	hour = h % 24
	return
}

func New(h int, m int) Clock {

	h, m = timeConvert(h, m)

	cl := Clock{minute: m, hour: h}
	return cl
}

func (clock Clock) String() string {

	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

func (clock Clock) Add(minutes int) Clock {
	return New(clock.hour, clock.minute+minutes)
}

func (clock Clock) Subtract(minutes int) Clock {
	return New(clock.hour, clock.minute-minutes)
}
