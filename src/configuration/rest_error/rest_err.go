package rest_error

import "net/http"

type Err struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *Err) Error() string {
	return r.Message
}

func NewError(message, err string, code int, causes []Causes) *Err {
	return &Err{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string, causes []Causes) *Err {
	return &Err{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *Err {
	return &Err{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string, err error) *Err {
	return &Err{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *Err {
	return &Err{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *Err {
	return &Err{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}

func NewUnauthorizedError(message string) *Err {
	return &Err{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}
