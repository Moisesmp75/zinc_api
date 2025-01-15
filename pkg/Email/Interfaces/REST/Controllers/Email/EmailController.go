package controllers

import (
	domain_services "mamuro_api/pkg/Email/Domain/Services"
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
