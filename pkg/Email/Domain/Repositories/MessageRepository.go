package repositories

import (
	entities "mamuro_api/pkg/Email/Domain/Model/Entities"
	shared_value_objects "mamuro_api/pkg/Shared/Domain/Model/ValueObjects"
)

type MessageRepository interface {
	Search(query string, size int, from int64, sort string) ([]entities.EmailMessage, shared_value_objects.Pagination, error)
	Post(emailMessaage entities.EmailMessage) (entities.EmailMessage, error)
}
