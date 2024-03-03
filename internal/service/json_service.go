package service

import (
	"encoding/json"
	"fmt"
	apperrors "json-processor/internal/error"

	ut "github.com/go-playground/universal-translator"

	"github.com/go-playground/validator/v10"
)

type IJSONService[T any] interface {
	ValidateJSON(jsonData []byte) (*T, error)
}

type JSONService[T any] struct {
	Validator *validator.Validate
	Trans     ut.Translator
}

func NewJSONService[T any](validator *validator.Validate, trans ut.Translator) *JSONService[T] {
	return &JSONService[T]{Validator: validator, Trans: trans}
}

func (js *JSONService[T]) ValidateJSON(jsonData []byte) (*T, error) {
	var data T
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", apperrors.ErrorJSONInputInvalid, err)
	}

	err = js.Validator.Struct(data)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make(map[string]string)
			for _, validationErr := range validationErrors {
				errorMessages[validationErr.Namespace()] = validationErr.Translate(js.Trans)
			}
			jsonErrorMessages, _ := json.MarshalIndent(errorMessages, "", "  ")
			return nil, fmt.Errorf("%w: %s", apperrors.ErrorFieldValidationFailed, jsonErrorMessages)
		}
	}

	return &data, nil
}
