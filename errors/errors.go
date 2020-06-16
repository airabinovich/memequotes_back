package errors

type UnauthorizedError struct {
	Message string
}

func NewUnauthorizedError(message string) UnauthorizedError {
	return UnauthorizedError{
		Message: message,
	}
}

func (err UnauthorizedError) Error() string {
	return err.Message
}
