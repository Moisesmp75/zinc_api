package repositories

import (
	entities "mamuro_api/pkg/Email/Domain/Model/Entities"
)

type MessageRepository interface {
	Search(query string, size int, from int64, sort string) ([]entities.EmailMessage, error)
	Post(emailMessaage entities.EmailMessage) (entities.EmailMessage, error)
}
