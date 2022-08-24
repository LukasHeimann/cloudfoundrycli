package errors

import (
	. "github.com/LukasHeimann/cloudfoundrycli/v8/cf/i18n"
)

type InvalidTokenError struct {
	description string
}

func NewInvalidTokenError(description string) error {
	return &InvalidTokenError{description: description}
}

func (err *InvalidTokenError) Error() string {
	return T("Invalid auth token: ") + err.description
}
