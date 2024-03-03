package service

import (
	"encoding/json"
	"fmt"
	apperrors "json-processor/internal/error"

	"github.com/go-playground/validator/v10"
)

type IJSONService[T any] interface {
	ValidateJSON(jsonData []byte) (*T, error)
}

type JSONService[T any] struct {
	Validator *validator.Validate
}

func (js *JSONService[T]) ValidateJSON(jsonData []byte) (*T, error) {
	var data T
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", apperrors.ErrorJSONInputInvalid, err)
	}

	err = js.Validator.Struct(data)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", apperrors.ErrorFieldValidationFailed, err)
	}

	return &data, nil
}
