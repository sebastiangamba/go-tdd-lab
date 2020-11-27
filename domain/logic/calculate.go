package logic

type LunchTimeCalculator interface {
	ObtainMinutesUntilLunchTime() int
}