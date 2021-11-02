package apperror

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// The error message
// swagger:response errorResponse
type errorResponse struct {
	// in: body
	Body Error
}

// Error model
type Error struct {
	Err error `json:"-"`
	// Response status code
	StatusCode int `json:"status_code"`
	// User message
	Message string `json:"message"`
	// Developer message
	DevMessage string `json:"dev_message"`
}

func NewError(err error, statusCode int, message string, devMessage string) *Error {
	if err == nil {
		err = fmt.Errorf(message)
	}
	return &Error{
		Err:        err,
		StatusCode: statusCode,
		Message:    message,
		DevMessage: devMessage,
	}
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

func (e *Error) Send(rw http.ResponseWriter) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(e.StatusCode)
	rw.Write(e.Marshal())
}

func teaPotError(err error) *Error {
	return NewError(err, 418, "Internal system error", err.Error())
}
