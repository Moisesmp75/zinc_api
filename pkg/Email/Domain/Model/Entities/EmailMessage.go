package entities

import (
	commands "mamuro_api/pkg/Email/Domain/Model/Commands"
	"time"
)

type EmailMessage struct {
	Message_ID                string
	Date                      time.Time
	From                      string
	To                        []string
	Subject                   string
	Cc                        []string
	Mime_version              string
	Content_Type              string
	Content_Transfer_Encoding string
	Bcc                       []string
	X_from                    string
	X_to                      []string
	X_cc                      []string
	X_bcc                     []string
	X_folder                  string
	X_origin                  string
	X_filename                string
	Content                   string
}

func EmailFromCommand(command commands.PostMessageCommand) *EmailMessage {
	return &EmailMessage{
		From:    command.From,
		To:      command.To,
		Subject: command.Subject,
		Date:    command.Date,
		Content: command.Content,
	}
}
