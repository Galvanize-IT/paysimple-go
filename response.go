package paysimple

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

type ErrorMessage struct {
	Field   string
	Message string
}
