package model

import "time"

type Enquiry struct {
	EnquiryId           string         `json:"enquiryId" validate:"required,len=32"`
	EnquiryTitle        string         `json:"enquiryTitle" validate:"required"`
	EnquiryDate         time.Time      `json:"enquiryDate" validate:"required"`
	EnquirerInfo        EnquirerInfo   `json:"enquirerInfo" validate:"required"`
	EnquiriedProperties []PropertyInfo `json:"properties" validate:"required,min=1,dive"`
}

type EnquirerInfo struct {
	FirstName    string `json:"firstName" validate:"required"`
	LastName     string `json:"lastName" validate:"required"`
	MobileNumber string `json:"mobileNumber" validate:"required,mobile"`
	EmailAddress string `json:"emailAddress" validate:"required,email"`
}
