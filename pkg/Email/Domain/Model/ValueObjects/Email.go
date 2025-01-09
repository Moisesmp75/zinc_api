package value_objects

import (
	"errors"
	"regexp"
	"strings"
)

type Email struct {
	value string
}

func (e *Email) GetDomain() string {
	parts := strings.Split(e.value, "@")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

func (e *Email) Get() string {
	return e.value
}

var emailPattern = regexp.MustCompile(`^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)

func NewEmail(email string) (*Email, error) {
	if email == "" {
		return nil, errors.New("el email es requerido")
	}

	email = strings.TrimSpace(email)
	if !isValidEmail(email) {
		return nil, errors.New("el email es inválido")
	}

	return &Email{value: email}, nil
}

func isValidEmail(email string) bool {
	return emailPattern.MatchString(email)
}
