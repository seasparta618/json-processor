package validation

import (
	"slices"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// RegisterCustomValidations registers custom validation rules with the validator.
func RegisterCustomValidations(validate *validator.Validate) {
	validate.RegisterValidation("postcode", validatePostCode)
	validate.RegisterValidation("state", validateState)
	// Add more custom rules here
}

// Custom validation function for PostCode
func validatePostCode(fl validator.FieldLevel) bool {
	postcode := fl.Field().Int()
	postcodeStart := []int64{2, 3, 4, 5, 6, 7}
	return postcode >= 2000 && postcode <= 7999 && slices.Contains(postcodeStart, postcode/1000)
}

// Custom validation function for State
func validateState(fl validator.FieldLevel) bool {
	state := fl.Field().String()
	states := []string{"VIC", "NSW", "ACT", "NT", "WA", "SA", "TAS", "QLD"}
	return slices.Contains(states, state)
}

func RegisterCustomTranslations(validate *validator.Validate, trans ut.Translator) {
	validate.RegisterTranslation("postcode", trans, func(ut ut.Translator) error {
		return ut.Add("postcode", "{0} must be a valid Australian postcode", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("postcode", fe.Field())
		return t
	})

	validate.RegisterTranslation("state", trans, func(ut ut.Translator) error {
		return ut.Add("state", "{0} must be a valid Australian state", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("state", fe.Field())
		return t
	})
}
