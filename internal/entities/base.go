package entities

import "time"

type BaseEntity struct {
	ID      int
	Created *time.Time
	Updated *time.Time
}
