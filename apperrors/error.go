package apperrors

type JournalAppError struct {
	ErrCode
	Message string
	Err     error `json:"-"`
}

func (e *JournalAppError) Error() string {
	return e.Err.Error()
}

func (e *JournalAppError) Unwrap() error {
	return e.Err
}
