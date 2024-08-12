package CustomError

import "runtime"

type CError struct {
	Message string
	File    string
	Line    int
}

func NewError(message string) error {
	_, file, line, _ := runtime.Caller(1)
	return &CError{
		Message: message,
		File:    file,
		Line:    line,
	}
}

func (e CError) Error() string {
	return e.Message
}
