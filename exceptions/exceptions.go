package exceptions

import "errors"

var (
	ErrValidation        error
	ErrBadData           = errors.New("user: unprocessable data received, failure with json payload")
	ErrUserAlreadyExists = errors.New("user: user with sent id already exists")
)
