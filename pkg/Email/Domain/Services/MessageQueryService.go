package services

import (
	entities "mamuro_api/pkg/Email/Domain/Model/Entities"
	queries "mamuro_api/pkg/Email/Domain/Model/Queries"
	shared_value_objects "mamuro_api/pkg/Shared/Domain/Model/ValueObjects"
)

type MessageQueryService interface {
	Search(query queries.SearchQuery) ([]entities.EmailMessage, shared_value_objects.Pagination, error)
}
