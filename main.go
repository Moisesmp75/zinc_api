package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.AllowContentType("application/json", "text/xml"))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	}))

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	apiv1 := chi.NewRouter()

	// controllers.AddControllers(apiv1)
	r.Mount("/api/v1", apiv1)

	log.Println("Running on port: 3000")
	http.ListenAndServe(fmt.Sprintf(":%v", 3000), r)
}
