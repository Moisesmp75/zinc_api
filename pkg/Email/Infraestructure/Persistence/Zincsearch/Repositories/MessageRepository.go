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

func NewMessageRepositoryInjection(zincsearchRepository sharedRepositories.ZincsearchRepository[zincEntities.Source]) *MessageRepository {
	return &MessageRepository{
		ZincsearchRepository: zincsearchRepository,
	}
}

func NewMessageRepository(baseUrl string, username string, password string) *MessageRepository {
	zincsearchRepository := sharedRepositories.ZincsearchRepository[zincEntities.Source]{
		BaseUrl:  baseUrl,
		Username: username,
		Password: password,
		Index:    "email",
	}

	return &MessageRepository{
		ZincsearchRepository: zincsearchRepository,
	}
}

func zincsearchResourceToDomainEntity(entity zincEntities.Source) domainEntities.EmailMessage {
	return domainEntities.EmailMessage{
		Message_ID:                entity.MessageID,
		Date:                      entity.Date,
		From:                      entity.From,
		To:                        entity.To,
		Subject:                   entity.Subject,
		Cc:                        entity.Cc,
		Mime_version:              string(entity.MIMEVersion),
		Content_Type:              string(entity.ContentType),
		Content_Transfer_Encoding: string(entity.ContentTransferEncoding),
		Bcc:                       entity.Bcc,
		X_from:                    entity.XFrom,
		X_to:                      entity.XTo,
		X_folder:                  entity.XFolder,
		X_origin:                  entity.XOrigin,
		X_filename:                entity.XFilename,
		Content:                   entity.Content,
		// X_cc:                      entity.XCc,
		// X_bcc:                     entity.XBcc,
	}
}

func (r *MessageRepository) Search(query string, size int, from int64, sort string) ([]domainEntities.EmailMessage, error) {
	result, err := r.ZincsearchRepository.Search(query, size, from, sort)
	if err != nil {
		return []domainEntities.EmailMessage{}, nil
	}

	messages := make([]domainEntities.EmailMessage, len(result.Hits.Hits))

	for i, data := range result.Hits.Hits {
		messages[i] = zincsearchResourceToDomainEntity(data.Source)
	}

	return messages, nil
}

func (r *MessageRepository) Post(emailMessage domainEntities.EmailMessage) (domainEntities.EmailMessage, error) {
	return domainEntities.EmailMessage{}, nil
}
