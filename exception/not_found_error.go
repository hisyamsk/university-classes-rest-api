package exception

type NotFoundError struct {
	ErrorMessage string
}

func (e *NotFoundError) Error() string {
	return e.ErrorMessage
}

func NewNotFoundError(error string) *NotFoundError {
	return &NotFoundError{
		ErrorMessage: error,
	}
}
