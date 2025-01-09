package commands

import "time"

type PostMessageCommand struct {
	From    string    `json:"from" validate:"required"`
	To      []string  `json:"to" validate:"required"`
	Date    time.Time `json:"date" validate:"required"`
	Subject string    `json:"subject" validate:"required"`
	Content string    `json:"content" validate:"required"`
}
