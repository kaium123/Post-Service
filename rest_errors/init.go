package rest_errors

import "errors"

// Map for errors with http code
var ResponseCode = make(map[string]int, 0)

func ResponseMap() map[string]int {
	return ResponseCode
}

func NewError(message string, httpCode int) error {
	_, available := ResponseCode[message]
	if !available {
		ResponseCode[message] = httpCode
	}
	return errors.New(message)
}
