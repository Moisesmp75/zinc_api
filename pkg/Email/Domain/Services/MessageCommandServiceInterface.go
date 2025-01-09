package services

import (
	commands "mamuro_api/pkg/Email/Domain/Model/Commands"
	entities "mamuro_api/pkg/Email/Domain/Model/Entities"
)

type MessageCommandServiceInterface interface {
	Post(command commands.PostMessageCommand) entities.EmailMessage
}
