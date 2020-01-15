package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rl404/go-malscraper/internal/controller"
	"github.com/rs/cors"
)

// startHTTP is a function to register routes and start the HTTP serve.
func startHTTP(port string) error {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Default().Handler)

	controller.RegisterBaseRoutes(r)

	r.Mount("/v1", controller.RegisterRoutesV1())

	fmt.Printf("Server listen at :%v\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

func main() {
	port := "8005"

	err := startHTTP(port)
	if err != nil {
		panic("fail")
	}
}
