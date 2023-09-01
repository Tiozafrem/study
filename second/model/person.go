package model

import "time"

type Person struct {
	Id        int
	Surname   string
	Age       int
	TimeStart time.Time
	TimeEnd   time.Time
	Interval  time.Duration
}
