package logic

import (
	"github.com/golang/mock/gomock"
	"go-tdd-lab/container"
	"go-tdd-lab/mocks"
	"os"
	"testing"
	"time"
)

var MockContainer *container.Dependencies

var sut LunchTimeCalculator

func TestMain(m *testing.M) {
	MockContainer = &container.Dependencies{}
	sut = NewLunchTimeCalculatorImpl(MockContainer)
	exitVal := m.Run()
	os.Exit(exitVal)
}

func Test_ShouldCallRepositoryWhenCalculatingTimeUntilLunch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepository := mocks.NewMockLunchTimeRepo(ctrl)

	// Se agrega el mock al contenedor de dependencias
	MockContainer.LunchTimeRepository = mockRepository

	mockRepository.EXPECT().GetLunchTime()

	sut.ObtainMinutesUntilLunchTime()
}

func Test_ShouldReturnSubtractionFromRepoResponseAndCurrentTime(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepository := mocks.NewMockLunchTimeRepo(ctrl)

	MockContainer.LunchTimeRepository = mockRepository

	//Setup
	location, _ := time.LoadLocation("America/Bogota")
	testLunchTime := time.Date(2020, 11, 27, 12, 30, 0, 0,location )
	expected := int(testLunchTime.Sub(time.Now()).Minutes())

	// se agrega el comportamiento
	mockRepository.EXPECT().GetLunchTime().Return(testLunchTime)

	got, _ := sut.ObtainMinutesUntilLunchTime()
	if got != expected {
		// Falla si el resultado no es el esperado
		t.Error("Result different from expected")
	}
}
