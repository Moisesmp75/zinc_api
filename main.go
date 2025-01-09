package main

import (
	"fmt"
	"log"
	repositories "mamuro_api/pkg/Email/Infraestructure/Persistence/Zincsearch/Repositories"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

var (
	username = "admin"
	password = "Complexpass#123"
	base_url = "http://localhost:4080"
)

func validar() {
	repository := repositories.NewMessageRepository(base_url, username, password)

	data, err := repository.Search("Hola", 10, 0, "date")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(len(data))

	fmt.Println(("Terminado"))
}

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

	validar()

	log.Println("Running on port: 3000")
	http.ListenAndServe(fmt.Sprintf(":%v", 3000), r)
}
