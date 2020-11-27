package backendImpl

import (
	"testing"
	"time"
)

var sut LunchTimeRepositoryImpl

func Test_ShouldReturnLunchTimeAfterNow(t *testing.T) {

	sut := NewLunchTimeRepositoryImpl()

	lunchTime := sut.GetLunchTime()
	if time.Now().After(lunchTime) {
		t.Error("Lunch time should be after current time")
	}
}
