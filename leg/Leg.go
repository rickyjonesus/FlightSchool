package leg

import "time"

type Leg struct {
	Id          int
	Origin      string
	Destination string
	Departs     time.Time
}