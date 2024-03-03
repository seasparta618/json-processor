package service

import (
	apperrors "json-processor/internal/error"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"required,min=18"`
	Email string `json:"email" validate:"required,email"`
}

func TestValidateJSON(t *testing.T) {
	validate := validator.New()
	jsonService := JSONService[TestStruct]{Validator: validate}

	validJSON := []byte(`{"name":"John Doe","age":30,"email":"john@example.com"}`)

	result, err := jsonService.ValidateJSON(validJSON)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "John Doe", result.Name)
	assert.Equal(t, 30, result.Age)
	assert.Equal(t, "john@example.com", result.Email)

	invalidJSON := []byte(`{"name":"Jane Doe","age":17}`)

	_, err = jsonService.ValidateJSON(invalidJSON)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), apperrors.ErrorFieldValidationFailed.Error())
}
