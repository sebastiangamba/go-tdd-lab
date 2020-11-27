package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go-tdd-lab/container"
	"go-tdd-lab/domain/logic"
)

func makePostDocumentDataEndpoint(dependencies *container.Dependencies) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		domainLogic := dependencies.DomainLogic.(logic.LunchTimeCalculator)
		return domainLogic.ObtainMinutesUntilLunchTime() // <- esto retorna el resultado y el err, por lo tanto los retornos son compatibles
	}
}
