package handlers

import (
	"fmt"
	services "mamuro_api/pkg/Email/Domain/Services"
	"net/http"
)

type MessageController struct {
	MessageCommandService services.MessageCommandService
	MessageQueryService   services.MessageQueryService
}

func NewMessageController(
	messageCommandService services.MessageCommandService,
	messageQueryService services.MessageQueryService) *MessageController {
	return &MessageController{
		MessageCommandService: messageCommandService,
		MessageQueryService:   messageQueryService,
	}
}

func (c *MessageController) SearchMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hola")
	bodyRequest := make([]byte, r.ContentLength)
	r.Body.Read(bodyRequest)

	w.Write(bodyRequest)
}
