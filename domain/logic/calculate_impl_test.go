package logic

import (
	"github.com/golang/mock/gomock"
	"go-tdd-lab/container"
	"go-tdd-lab/mocks"
	"os"
	"testing"
)

var MockContainer *container.Dependencies

var sut LunchTimeCalculator

func TestMain(m *testing.M) {
	MockContainer = &container.Dependencies{}
	sut = NewLunchTimeCalculatorImpl(MockContainer)
	// Esto va por nomenclatura del TestMain
	exitVal := m.Run()
	os.Exit(exitVal)
}

func Test_ShouldCallRepositoryWhenCalculatingTimeUntilLunch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepository := mocks.NewMockLunchTimeRepo(ctrl)

	mockRepository.EXPECT().GetLunchTime()

	sut.ObtainMinutesUntilLunchTime()
}
