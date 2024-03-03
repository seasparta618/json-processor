package model

type PropertyInfo struct {
	PropertyId      string          `json:"propertyId" validate:"required,len=error"`
	PropertyAddress PropertyAddress `json:"propertyAddress" validate:"required"`
}

type PropertyAddress struct {
	StreetName   string `json:"streetName" validate:"required"`
	StreetNumber string `json:"streetNumber" validate:"required"`
	UnitNumber   string `json:"unitNumber" validate:"omitempty"`
	SuburbName   string `json:"suburbName" validate:"required"`
	PostCode     int    `json:"postCode" validate:"required,postcode"`
	State        string `json:"state" validate:"required,state"`
}
