package repositories

import (
	domainEntities "mamuro_api/pkg/Email/Domain/Model/Entities"
	domainRepositories "mamuro_api/pkg/Email/Domain/Repositories"
	zincEntities "mamuro_api/pkg/Email/Infraestructure/Persistence/Zincsearch/Entities"
	sharedRepositories "mamuro_api/pkg/Shared/Infraestructure/Persistence/Zincsearch/Repositories"
)

type MessageRepository struct {
	ZincsearchRepository sharedRepositories.ZincsearchRepository[zincEntities.Source]
}

var _ domainRepositories.MessageRepoInterface = (*MessageRepository)(nil)

func (r *MessageRepository) Search(query string, size int, from int64, sort string) (domainEntities.EmailMessage, error) {
	return domainEntities.EmailMessage{}, nil
}

func (r *MessageRepository) Post(emailMessage domainEntities.EmailMessage) (domainEntities.EmailMessage, error) {
	return domainEntities.EmailMessage{}, nil
}
