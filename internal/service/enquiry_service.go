package service

import (
	"bytes"
	"fmt"
	"json-processor/config"
	httpclient "json-processor/internal/http"
	"json-processor/internal/model"
	"net/http"
)

type EnquiryService struct {
	JSONValidator IJSONService[model.Enquiry]
	HTTPClient    httpclient.HttpClient
	Config        *config.Config
}

func NewEnquiryService(jsonValidator IJSONService[model.Enquiry], httpClient httpclient.HttpClient, config *config.Config) *EnquiryService {
	return &EnquiryService{
		JSONValidator: jsonValidator,
		HTTPClient:    httpClient,
		Config:        config,
	}
}

func (es *EnquiryService) SaveEnquiry(jsonData []byte) error {
	_, err := es.JSONValidator.ValidateJSON(jsonData)
	if err != nil {
		return err
	}

	apiEndpoint := "/enquiry/save/"
	uri := es.Config.APIHost + apiEndpoint

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := es.HTTPClient.Do(req)
	if err != nil {
		// should have specified error handling here, it includes retry or wrap the error, but requirement did not specify
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	return nil
}