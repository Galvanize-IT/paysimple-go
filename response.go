package paysimple

import (
	"fmt"
	"strings"
)

type ErrorCode string

const (
	InvalidInput       ErrorCode = "InvalidInput"
	InvalidPermissions ErrorCode = "InvalidPermissions"
	NotFound           ErrorCode = "NotFound"
	UnexpectedError    ErrorCode = "UnexpectedError"
)

type Meta struct {
	Errors         *Errors
	HttpStatus     string
	HttpStatusCode int64
	PagingDetails  struct {
		TotalItems   int64
		Page         int64
		ItemsPerPage int64
	}
}

type Errors struct {
	ErrorCode     ErrorCode
	ErrorMessages []ErrorMessage
	TraceCode     string
}

var _ error = Errors{}

// Error implements the built-in interface for Errors
func (err Errors) Error() string {
	fields := make([]string, len(err.ErrorMessages))
	for i, msg := range err.ErrorMessages {
		fields[i] = fmt.Sprintf("(%s)", msg)
	}
	return fmt.Sprintf(
		"%s - %s: %s",
		err.ErrorCode,
		err.TraceCode,
		strings.Join(fields, " "),
	)
}

type ErrorMessage struct {
	Field   string
	Message string
}

func (msg ErrorMessage) String() string {
	return fmt.Sprintf("%s: %s", msg.Field, msg.Message)
}

// Empty responses are returned when detail operations fail
type Empty struct {
	Meta Meta
	// Ignore response, it is irrelevant
}
