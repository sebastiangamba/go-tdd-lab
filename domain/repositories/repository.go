package repositories

import "time"

type LunchTimeRepo interface {
	GetLunchTime() time.Time
}
