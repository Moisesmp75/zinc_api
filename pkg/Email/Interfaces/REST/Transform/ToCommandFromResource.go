package transform

import (
	commands "mamuro_api/pkg/Email/Domain/Model/Commands"
	resources "mamuro_api/pkg/Email/Interfaces/REST/Resources"
)

func CreateMessageResourceToCommand(resource resources.CreateMessageResource) *commands.PostMessageCommand {
	return &commands.PostMessageCommand{
		From:    resource.From,
		To:      resource.To,
		Date:    resource.Date,
		Subject: resource.Subject,
		Content: resource.Content,
	}

}
