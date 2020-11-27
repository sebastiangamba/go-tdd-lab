package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-tdd-lab/container"
	"go-tdd-lab/infraestructure/http/endpoint"
	"go-tdd-lab/infraestructure/http/model"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
	kithttp "github.com/go-kit/kit/transport/http"
)

type Transport struct {
	dependencies *container.Dependencies
	Endpoint     string
	server       *http.ServeMux
}

func NewTransport(dependencies *container.Dependencies, endpoint string, server *http.ServeMux) *Transport {
	return &Transport{dependencies: dependencies, Endpoint: endpoint, server: server}
}

func (transport *Transport) HandleEndpoint() error {
	serverOptions := []kithttp.ServerOption{}

	postDocumentDataHandler := kithttp.NewServer(
		endpoint.MakeLunchTimeEndpoint(transport.dependencies),
		decodeRequest,
		encodeResponse,
		serverOptions...,
	)
	r := mux.NewRouter()
	r.Handle(transport.Endpoint, handlers.LoggingHandler(os.Stdout, postDocumentDataHandler))
	(transport.server).Handle(transport.Endpoint, r)
	return nil
}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func encodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	if response == nil {
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	responseModel := model.ResponseModel{
		RemainingMinutes: response.(int),
	}

	return json.NewEncoder(w).Encode(responseModel)

}