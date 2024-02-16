package exceptions

import (
	"errors"
	"net/http"
)

type ErrResponse struct {
	Code    int
	Message string
}

func HandleException(err error) ErrResponse {
	errRes, ok := err.(*Error)
	if !ok {
		return ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	switch errRes != nil {
	case
		errors.Is(errRes.CustomErr, ErrValidation):
		return ErrResponse{
			Code:    http.StatusBadRequest,
			Message: errRes.Error(),
		}
	case
		errors.Is(errRes.CustomErr, ErrBadData):
		return ErrResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: errRes.CustomErr.Error(),
		}
	default:
		return ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		}
	}
}
