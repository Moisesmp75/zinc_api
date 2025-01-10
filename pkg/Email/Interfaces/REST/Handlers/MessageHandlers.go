package handlers

import (
	"fmt"
	"log"
	domain_services "mamuro_api/pkg/Email/Domain/Services"
	resources "mamuro_api/pkg/Email/Interfaces/REST/Resources"
	transform "mamuro_api/pkg/Email/Interfaces/REST/Transform"
	"net/http"
)

type MessageController struct {
	MessageCommandService domain_services.MessageCommandService
	MessageQueryService   domain_services.MessageQueryService
}

func NewMessageController(
	messageCommandService domain_services.MessageCommandService,
	messageQueryService domain_services.MessageQueryService) *MessageController {
	return &MessageController{
		MessageCommandService: messageCommandService,
		MessageQueryService:   messageQueryService,
	}
}

// bodyRequest := make([]byte, r.ContentLength)
// r.Body.Read(bodyRequest)

func (c *MessageController) SearchMessage(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()

	queryP := queryParams.Get("query")
	size := queryParams.Get("size")
	from := queryParams.Get("from")
	sort := queryParams.Get("sort")

	resource := resources.NewSearchMessageQuery(queryP, size, from, sort)
	query, err := transform.SearchQueryResourceToQuery(*resource)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(query)
	w.Write([]byte(`{ "Mensaje": "Hola Mundo" }`))
}
