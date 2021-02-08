package reservation

import "time"

type Reservation struct {
	Id int
	Starts time.Time
	Ends time.Time
	MemberId int
	AircraftId int
}