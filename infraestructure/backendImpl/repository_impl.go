package backendImpl

import (
	"time"
)

type LunchTimeRepositoryImpl struct {
}

func NewLunchTimeRepositoryImpl() *LunchTimeRepositoryImpl {
	return &LunchTimeRepositoryImpl{}
}

func (l *LunchTimeRepositoryImpl) GetLunchTime() time.Time {
	location, _ := time.LoadLocation("America/Bogota")
	return time.Date(2020, 11, 27, 12, 30, 0, 0,location )
}

