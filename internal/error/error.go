package apperrors

import "errors"

var (
	ErrorFileNotFound          = errors.New("file not found")
	ErrorEmptyContent          = errors.New("input content is empty")
	ErrorJSONInputInvalid      = errors.New("input is not valid json")
	ErrorFieldValidationFailed = errors.New("field validation failed")
)
