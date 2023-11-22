package common

type ErrorType string

const (
	NotFound      ErrorType = "NOT_FOUND"
	InternalError ErrorType = "INTERNAL_ERROR"
	Invalid       ErrorType = "INVALID"
	Unauthorized  ErrorType = "UNAUTHORIZED"
	BadRequest    ErrorType = "BAD_REQUEST"
	Duplicate     ErrorType = "DUPLICATE"
)

type ErrorResponse struct {
	message   string
	errorType ErrorType
}

func NewErrorResponse(err_type ErrorType, msg string) *ErrorResponse {
	return &ErrorResponse{
		message:   msg,
		errorType: err_type,
	}
}

func (err ErrorResponse) Error() string {
	return err.message
}

func (err ErrorResponse) Type() ErrorType {
	return err.errorType
}

type HttpErrorResponse struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code,omitempty"`
}