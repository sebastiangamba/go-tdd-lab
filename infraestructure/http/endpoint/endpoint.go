package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go-tdd-lab/container"
)

func makePostDocumentDataEndpoint(dependencies *container.Dependencies) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return nil, nil
	}
}
