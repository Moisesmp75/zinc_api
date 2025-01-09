package repositories

import (
	entities "mamuro_api/pkg/Email/Domain/Model/Entities"
)

type MessageRepoInterface interface {
	Search(query string, size int, from int64, sort string) entities.EmailMessage
	Post(emailMessaage entities.EmailMessage) entities.EmailMessage
}
