package command_services

import (
	commands "mamuro_api/pkg/Email/Domain/Model/Commands"
	entities "mamuro_api/pkg/Email/Domain/Model/Entities"
	domain_repositories "mamuro_api/pkg/Email/Domain/Repositories"
	domain_services "mamuro_api/pkg/Email/Domain/Services"
)

type MessageCommandService struct {
	MessageRepository domain_repositories.MessageRepository
}

func NewMessageCommandService(messageRepository domain_repositories.MessageRepository) *MessageCommandService {
	return &MessageCommandService{
		MessageRepository: messageRepository,
	}
}

var _ domain_services.MessageCommandService = (*MessageCommandService)(nil)

func (m *MessageCommandService) Post(command commands.PostMessageCommand) (entities.EmailMessage, error) {
	entity := entities.EmailFromCommand(command)
	result, err := m.MessageRepository.Post(*entity)
	if err != nil {
		return entities.EmailMessage{}, err
	}
	return result, nil
}
