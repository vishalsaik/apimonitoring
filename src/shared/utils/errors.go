package utils

import "net/http"

type AppError struct {
	Message       string
	StatusCode    int
	Errors        interface{}
	IsOperational bool
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(message string, statusCode int, errors interface{}) *AppError {
	return &AppError{
		Message:       message,
		StatusCode:    statusCode,
		Errors:        errors,
		IsOperational: true,
	}
}

func NewNotFoundError(message string) *AppError {
	return NewAppError(message, http.StatusNotFound, nil)
}

func NewUnauthorizedError(message string) *AppError {
	return NewAppError(message, http.StatusUnauthorized, nil)
}

func NewForbiddenError(message string) *AppError {
	return NewAppError(message, http.StatusForbidden, nil)
}
