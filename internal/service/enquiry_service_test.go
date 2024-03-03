package service

import (
	"bytes"
	"io"
	"json-processor/config"
	apperrors "json-processor/internal/error"
	"json-processor/internal/model"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockJSONService is a mock implementation of IJSONService for testing
type MockJSONService struct {
	mock.Mock
}

func (m *MockJSONService) ValidateJSON(jsonData []byte) (*model.Enquiry, error) {
	args := m.Called(jsonData)
	return nil, args.Error(1) // Return nil as the first value when there is an error
}

// MockHttpClient is a mock implementation of HttpClient for testing
type MockHttpClient struct {
	mock.Mock
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

// mock a test without any error
func TestEnquiryService_SaveEnquiry_ValidationSuccessful(t *testing.T) {
	// Set up mocks
	mockJSONValidator := new(MockJSONService)
	mockJSONValidator.On("ValidateJSON", mock.Anything).Return(&model.Enquiry{}, nil)

	mockHTTPClient := new(MockHttpClient)
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString("")),
	}
	mockHTTPClient.On("Do", mock.Anything).Return(mockResponse, nil)

	// Set up service
	cfg := &config.Config{
		APIHost: "https://example.com",
	}
	enquiryService := NewEnquiryService(mockJSONValidator, mockHTTPClient, cfg)

	// Test SaveEnquiry
	err := enquiryService.SaveEnquiry([]byte(`{"key":"value"}`))
	assert.NoError(t, err)

	// Assert that the mocks were called as expected
	mockJSONValidator.AssertExpectations(t)
	mockHTTPClient.AssertExpectations(t)
}

func TestEnquiryService_SaveEnquiry_ValidationFailed(t *testing.T) {
	// Set up mocks
	mockJSONValidator := new(MockJSONService)
	validationErr := apperrors.ErrorFieldValidationFailed
	mockJSONValidator.On("ValidateJSON", mock.Anything).Return((*model.Enquiry)(nil), validationErr)

	mockHTTPClient := new(MockHttpClient) // No need to set up expectations, as it should not be called

	// Set up service
	cfg := &config.Config{
		APIHost: "https://example.com",
	}
	enquiryService := NewEnquiryService(mockJSONValidator, mockHTTPClient, cfg)

	// Test SaveEnquiry
	err := enquiryService.SaveEnquiry([]byte(`{"key":"value"}`))
	assert.Error(t, err)
	assert.Equal(t, validationErr, err)

	// Assert that the mocks were called as expected
	mockJSONValidator.AssertExpectations(t)
	mockHTTPClient.AssertNotCalled(t, "Do", mock.Anything)
}
