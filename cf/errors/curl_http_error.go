package errors

import (
	. "github.com/LukasHeimann/cloudfoundrycli/v8/cf/i18n"
)

type CurlHTTPError struct {
	StatusCode int
}

func NewCurlHTTPError(statusCode int) error {
	return &CurlHTTPError{StatusCode: statusCode}
}

func (err CurlHTTPError) Error() string {
	return T("The requested URL returned error: {{.StatusCode}}", map[string]interface{}{"StatusCode": err.StatusCode})
}
