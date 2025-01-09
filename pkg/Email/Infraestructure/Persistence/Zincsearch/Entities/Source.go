package entities

import (
	shared "mamuro_api/pkg/Shared/Infraestructure/Persistence/Zincsearch/Entities"
	"time"
)

type Source struct {
	Timestamp               string                         `json:"@timestamp"`
	Bcc                     []string                       `json:"bcc"`
	Cc                      []string                       `json:"cc"`
	Content                 string                         `json:"content"`
	ContentTransferEncoding shared.ContentTransferEncoding `json:"content_transfer_encoding"`
	ContentType             shared.ContentType             `json:"content_type"`
	Date                    time.Time                      `json:"date"`
	From                    string                         `json:"from"`
	MessageID               string                         `json:"message_id"`
	MIMEVersion             shared.MIMEVersion             `json:"mime_version"`
	Subject                 string                         `json:"subject"`
	To                      []string                       `json:"to"`
	XBcc                    interface{}                    `json:"x_bcc"`
	XCc                     interface{}                    `json:"x_cc"`
	XFilename               string                         `json:"x_filename"`
	XFolder                 string                         `json:"x_folder"`
	XFrom                   string                         `json:"x_from"`
	XOrigin                 string                         `json:"x_origin"`
	XTo                     []string                       `json:"x_to"`
}
