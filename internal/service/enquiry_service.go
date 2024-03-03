package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"json-processor/config"
	httpclient "json-processor/internal/http"
	"json-processor/internal/model"
	"json-processor/internal/util"
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
	enquiry, err := es.JSONValidator.ValidateJSON(jsonData)
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	jsonOutput, err := json.MarshalIndent(enquiry, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	fmt.Println(string(jsonOutput))

	apiEndpoint := "/api/enquiry/save/"
	uri := es.Config.APIHost + apiEndpoint

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+es.Config.APIToken)
	resp, err := es.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	return util.HandleHTTPResponse(resp)
}
