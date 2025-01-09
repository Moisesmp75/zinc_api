package services

import (
	entities "mamuro_api/pkg/Email/Domain/Model/Entities"
	queries "mamuro_api/pkg/Email/Domain/Model/Queries"
)

type MessageQueryServiceInterface interface {
	Search(query queries.SearchQuery) entities.EmailMessage
}
