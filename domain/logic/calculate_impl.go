package logic

import (
	"go-tdd-lab/container"
	"go-tdd-lab/domain/repositories"
)

type LunchTimeCalculatorImpl struct {
	dependencies *container.Dependencies
}

func NewLunchTimeCalculatorImpl(dependencies *container.Dependencies) LunchTimeCalculator {
	return &LunchTimeCalculatorImpl{dependencies: dependencies}
}

func (l *LunchTimeCalculatorImpl) ObtainMinutesUntilLunchTime() (int, error) {

	repo := l.dependencies.LunchTimeRepository.(repositories.LunchTimeRepo)
	repo.GetLunchTime()
	return 0, nil
}




