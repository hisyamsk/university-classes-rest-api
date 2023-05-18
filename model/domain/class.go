package domain

import "time"

type Class struct {
	Id      int
	Name    string
	startAt time.Time
	endAt   time.Time
}
