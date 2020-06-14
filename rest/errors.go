package rest


import (
	"net/http"
	"strings"
)

const (
	// BadRequestMessage is the default message when the input parameters on a request are wrong or it is malformed.
	BadRequestMessage = "Invalid request parameters."
	// ResourceNotFoundMessage is the default message when a requested resource is not available.
	ResourceNotFoundMessage = "Resource not found."
	// MethodNotAllowedMessage is the default message when a HTTP verb is forbidden on a resource.
	MethodNotAllowedMessage = "Method not allowed on the current resource."
	// InternalServerErrorMessage is the default message when an unexpected condition occurs.
	InternalServerErrorMessage = "Internal Server Error."
)

// APIError represents the standard error structure for the HTTP responses.
type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Err     string `json:"error"`
}

// newAPIError creates and initializes an APIError.
func newAPIError(code int, message string, err string) *APIError {
	return &APIError{
		Status:  code,
		Message: message,
		Err:     err,
	}
}

// NewBadRequest creates an API Error for an invalid or malformed request.
func NewBadRequest(messages ...string) *APIError {
	message := BadRequestMessage
	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}

	return newAPIError(http.StatusBadRequest, message, "bad_request")
}

// NewResourceNotFound creates an API Error for an unexisting resource.
func NewResourceNotFound(messages ...string) *APIError {
	message := ResourceNotFoundMessage
	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}

	return newAPIError(http.StatusNotFound, message, "not_found")
}

// NewMethodNotAllowed creates an API Error for a forbidden verb on a resource.
func NewMethodNotAllowed(messages ...string) *APIError {
	message := MethodNotAllowedMessage
	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}

	return newAPIError(http.StatusMethodNotAllowed, message, "method_not_allowed")
}

// NewInternalServerError creates an API Error for an unexpected condition.
func NewInternalServerError(messages ...string) *APIError {
	message := InternalServerErrorMessage
	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}

	return newAPIError(http.StatusInternalServerError, message, "internal_error")
}
