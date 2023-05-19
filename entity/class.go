package entity

import "time"

type Class struct {
	Id      int
	Name    string
	StartAt time.Time
	EndAt   time.Time
}
