package logic

import "go-tdd-lab/container"

type LunchTimeCalculatorImpl struct {
	dependencies *container.Dependencies
}

func NewLunchTimeCalculatorImpl(dependencies *container.Dependencies) LunchTimeCalculator {
	return &LunchTimeCalculatorImpl{dependencies: dependencies}
}

func (l *LunchTimeCalculatorImpl) ObtainMinutesUntilLunchTime() (int, error) {
	return 0, nil
}




