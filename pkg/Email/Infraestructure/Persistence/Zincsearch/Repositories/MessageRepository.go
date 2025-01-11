package repositories

import (
	domain_entities "mamuro_api/pkg/Email/Domain/Model/Entities"
	domain_repositories "mamuro_api/pkg/Email/Domain/Repositories"
	zinc_entities "mamuro_api/pkg/Email/Infraestructure/Persistence/Zincsearch/Entities"
	share_value_objects "mamuro_api/pkg/Shared/Domain/Model/ValueObjects"
	shared_repositories "mamuro_api/pkg/Shared/Infraestructure/Persistence/Zincsearch/Repositories"
)

type MessageRepository struct {
	ZincsearchRepository shared_repositories.ZincsearchRepository[zinc_entities.Source]
}

var _ domain_repositories.MessageRepository = (*MessageRepository)(nil)

func NewMessageRepositoryInjection(zincsearchRepository shared_repositories.ZincsearchRepository[zinc_entities.Source]) *MessageRepository {
	return &MessageRepository{
		ZincsearchRepository: zincsearchRepository,
	}
}

func NewMessageRepository(baseUrl string, username string, password string) *MessageRepository {
	zincsearchRepository := shared_repositories.ZincsearchRepository[zinc_entities.Source]{
		BaseUrl:  baseUrl,
		Username: username,
		Password: password,
		Index:    "email",
	}

	return &MessageRepository{
		ZincsearchRepository: zincsearchRepository,
	}
}

func zincsearchResourceToDomainEntity(entity zinc_entities.Source) domain_entities.EmailMessage {
	return domain_entities.EmailMessage{
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

func (r *MessageRepository) Search(query string, size int, from int64, sort string) ([]domain_entities.EmailMessage, share_value_objects.Pagination, error) {
	result, pagination, err := r.ZincsearchRepository.Search(query, size, from, sort)
	if err != nil {
		return []domain_entities.EmailMessage{}, share_value_objects.Pagination{}, err
	}

	messages := make([]domain_entities.EmailMessage, len(result.Hits.Hits))

	for i, data := range result.Hits.Hits {
		messages[i] = zincsearchResourceToDomainEntity(data.Source)
	}

	return messages, pagination, nil
}

func (r *MessageRepository) Post(emailMessage domain_entities.EmailMessage) (domain_entities.EmailMessage, error) {
	return domain_entities.EmailMessage{}, nil
}
