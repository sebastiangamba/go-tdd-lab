package main

import (
	"fmt"
	"go-tdd-lab/container"
	"go-tdd-lab/domain/logic"
	"go-tdd-lab/infraestructure/backendImpl"
	"go-tdd-lab/infraestructure/http/transport"
	"log"
	"net/http"
)

func main() {
	var (
		httpAddr = fmt.Sprintf(":%s", "8080")
	)

	//Dependencies
	dependenciesContainer := &container.Dependencies{}
	// Acá se agregan las implementaciones de las dependencias!
	dependenciesContainer.DomainLogic = logic.NewLunchTimeCalculatorImpl(dependenciesContainer)
	dependenciesContainer.LunchTimeRepository = backendImpl.NewLunchTimeRepositoryImpl()

	serveMux := http.NewServeMux()
	lunchTimeTransport := transport.NewTransport(dependenciesContainer, "/lunchtime", serveMux)

	// Endpoint handling
	http.Handle("/", accessControl(serveMux))
	_ = lunchTimeTransport.HandleEndpoint()

	errs := make(chan error)
	go func() {
		server := &http.Server{
			Addr:    httpAddr,
			Handler: nil,
		}
		log.Printf("Server started")
		errs <- server.ListenAndServe()
	}()
	log.Printf("Server started %v\n", <-errs)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
