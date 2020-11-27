package logic

import (
	"go-tdd-lab/container"
	"go-tdd-lab/domain/repositories"
	"time"
)

type LunchTimeCalculatorImpl struct {
	dependencies *container.Dependencies
}

func NewLunchTimeCalculatorImpl(dependencies *container.Dependencies) LunchTimeCalculator {
	return &LunchTimeCalculatorImpl{dependencies: dependencies}
}

func (l *LunchTimeCalculatorImpl) ObtainMinutesUntilLunchTime() (int, error) {

	repo := l.dependencies.LunchTimeRepository.(repositories.LunchTimeRepo)
	lunchTime := repo.GetLunchTime()

	return int(lunchTime.Sub(time.Now()).Minutes()), nil
}




