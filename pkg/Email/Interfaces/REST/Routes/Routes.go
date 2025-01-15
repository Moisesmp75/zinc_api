package routes

import (
	email_handlers "mamuro_api/pkg/Email/Interfaces/REST/Controllers/Email"

	"github.com/go-chi/chi/v5"
)

func EmailSetRoutes(r *chi.Mux, messageController *email_handlers.MessageController) {
	r.Route("/email", func(r chi.Router) {
		r.Get("/search", messageController.SearchMessage)
	})
}
