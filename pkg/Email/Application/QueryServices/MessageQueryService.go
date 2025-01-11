package query_services

import (
	entities "mamuro_api/pkg/Email/Domain/Model/Entities"
	queries "mamuro_api/pkg/Email/Domain/Model/Queries"
	domain_repositories "mamuro_api/pkg/Email/Domain/Repositories"
	domain_services "mamuro_api/pkg/Email/Domain/Services"
	shared_value_objects "mamuro_api/pkg/Shared/Domain/Model/ValueObjects"
)

type MessageQueryService struct {
	MessageRepository domain_repositories.MessageRepository
}

func NewMessageQueryService(messageRepository domain_repositories.MessageRepository) *MessageQueryService {
	return &MessageQueryService{
		MessageRepository: messageRepository,
	}
}

var _ domain_services.MessageQueryService = (*MessageQueryService)(nil)

func (m *MessageQueryService) Search(query queries.SearchQuery) ([]entities.EmailMessage, shared_value_objects.Pagination, error) {
	result, pagination, err := m.MessageRepository.Search(query.Query, query.Size, query.From, query.Sort)
	if err != nil {
		return []entities.EmailMessage{}, shared_value_objects.Pagination{}, err
	}
	return result, pagination, nil
}
