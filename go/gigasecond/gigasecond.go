package gigasecond

import "time"

//import "fmt"

const gigasecond = time.Duration(1e9 * time.Second)

func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
