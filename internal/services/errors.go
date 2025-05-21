package services

import "errors"

var (
	InternalServerError    = errors.New("InternalServerError")
	AlreadyExistsError     = errors.New("AlreadyExistsError")
	InvalidParametersError = errors.New("InvalidParametersError")
	AccessDenied           = errors.New("Access denied")
)
