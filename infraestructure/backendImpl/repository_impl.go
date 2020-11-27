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
	panic("implement me")
}

