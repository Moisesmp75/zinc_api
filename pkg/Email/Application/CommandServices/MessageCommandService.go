package command_services

import (
	commands "mamuro_api/pkg/Email/Domain/Model/Commands"
	entities "mamuro_api/pkg/Email/Domain/Model/Entities"
	domainRepositories "mamuro_api/pkg/Email/Domain/Repositories"
	domainServices "mamuro_api/pkg/Email/Domain/Services"
)

type MessageCommandService struct {
	MessageRepository domainRepositories.MessageRepository
}

var _ domainServices.MessageCommandService = (*MessageCommandService)(nil)

func (m *MessageCommandService) Post(command commands.PostMessageCommand) (entities.EmailMessage, error) {
	panic("unimplemented")
}
