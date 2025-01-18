package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"

	email_command_services "mamuro_api/pkg/Email/Application/CommandServices"
	email_query_services "mamuro_api/pkg/Email/Application/QueryServices"
	email_repositories "mamuro_api/pkg/Email/Infraestructure/Persistence/Zincsearch/Repositories"
	email_controllers "mamuro_api/pkg/Email/Interfaces/REST/Controllers/Email"
	email_routes "mamuro_api/pkg/Email/Interfaces/REST/Routes"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var (
	username = "admin"
	password = "Complexpass#123"
	base_url = "http://localhost:4080"
)

//go:embed web/dist
var embedFrontend embed.FS

func static(r *chi.Mux) {
	front, err := fs.Sub(embedFrontend, "web/dist")
	if err != nil {
		fmt.Println(err.Error())
	}

	staticServer := http.FileServer(http.FS(front))

	r.Get("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		staticServer.ServeHTTP(w, r)
	}))
}

func main() {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	}))

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!"))
	// })
	static(r)
	apiv1 := chi.NewRouter()

	apiv1.Use(middleware.AllowContentType("application/json", "text/xml"))
	apiv1.Use(middleware.Logger)

	messageRepository := email_repositories.NewMessageRepository(base_url, username, password)
	messageCommandService := email_command_services.NewMessageCommandService(messageRepository)
	messageQueryService := email_query_services.NewMessageQueryService(messageRepository)
	messageController := email_controllers.NewMessageController(messageCommandService, messageQueryService)

	email_routes.EmailSetRoutes(apiv1, messageController)
	r.Mount("/api/v1", apiv1)

	log.Println("Running on port: 3000")
	http.ListenAndServe(fmt.Sprintf(":%v", 3000), r)
}
