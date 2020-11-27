package endpoint

import (
	"context"
	"github.com/golang/mock/gomock"
	"go-tdd-lab/container"
	"go-tdd-lab/mocks"
	"os"
	"testing"
)


var MockContainer *container.Dependencies

func TestMain(m *testing.M) {
	MockContainer = &container.Dependencies{}
	exitVal := m.Run()
	os.Exit(exitVal)
}

func Test_ShouldReturnDomainLogicResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDomainLogic := mocks.NewMockLunchTimeCalculator(ctrl)
	MockContainer.DomainLogic = mockDomainLogic

	var testContext context.Context
	var testRequest interface{}

	domainResponse := 10
		mockDomainLogic.EXPECT().ObtainMinutesUntilLunchTime().Return(domainResponse, nil)

	endpointFunc := makePostDocumentDataEndpoint(MockContainer)
	got, _ := endpointFunc(testContext, testRequest)

	if got != domainResponse {
		t.Error("Endpoint response Different from expected")
	}
}
