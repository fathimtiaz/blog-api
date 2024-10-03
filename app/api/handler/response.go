package handler

import "errors"

type response struct {
	Success      bool        `json:"success"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"data"`
}

var (
	ErrBadRequest = errors.New("bad request")
	ErrServer     = errors.New("server error")
)

func setResponse(err error, data interface{}) response {
	var success = true
	var message = ""

	if err != nil {
		success = false
		message = err.Error()
	}

	return response{
		Success:      success,
		ErrorMessage: message,
		Data:         data,
	}
}
