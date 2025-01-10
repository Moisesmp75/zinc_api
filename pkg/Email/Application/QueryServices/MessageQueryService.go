package query_services

import (
	"log"
	entities "mamuro_api/pkg/Email/Domain/Model/Entities"
	queries "mamuro_api/pkg/Email/Domain/Model/Queries"
	domainRepositories "mamuro_api/pkg/Email/Domain/Repositories"
	domainServices "mamuro_api/pkg/Email/Domain/Services"
)

type MessageQueryService struct {
	MessageRepository domainRepositories.MessageRepository
}

func NewMessageQueryService(messageRepository domainRepositories.MessageRepository) *MessageQueryService {
	return &MessageQueryService{
		MessageRepository: messageRepository,
	}
}

var _ domainServices.MessageQueryService = (*MessageQueryService)(nil)

func (m *MessageQueryService) Search(query queries.SearchQuery) ([]entities.EmailMessage, error) {
	result, err := m.MessageRepository.Search(query.Query, query.Size, query.From, query.Sort)
	if err != nil {
		log.Println(err.Error())
		return []entities.EmailMessage{}, err
	}
	return result, nil
}
