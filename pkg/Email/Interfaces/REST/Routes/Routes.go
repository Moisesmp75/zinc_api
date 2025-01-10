package routes

import (
	handlers "mamuro_api/pkg/Email/Interfaces/REST/Handlers"

	"github.com/go-chi/chi/v5"
)

func EmailSetRoutes(r *chi.Mux, messageController *handlers.MessageController) {
	r.Route("/email", func(r chi.Router) {
		r.Get("/search", messageController.SearchMessage)
	})
}
