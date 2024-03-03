package validation

import (
	"slices"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// RegisterCustomValidations registers custom validation rules with the validator.
// use the pointer type here, in this case, the validator just needed to be defined once, then it can bind the validation rules with this validator in full life cycle
func RegisterCustomValidations(validate *validator.Validate, trans ut.Translator) {
	en_translations.RegisterDefaultTranslations(validate, trans)
	validate.RegisterValidation("postcode", validatePostCode)
	validate.RegisterValidation("state", validateState)
	validate.RegisterValidation("mobile", validateMobile)
	RegisterCustomTranslations(validate, trans)
}

// Custom validation function for PostCode
func validatePostCode(fl validator.FieldLevel) bool {
	postcode := fl.Field().Int()
	postcodeStart := []int64{2, 3, 4, 5, 6, 7}
	return postcode >= 2000 && postcode <= 7999 && slices.Contains(postcodeStart, postcode/1000)
}

// Custom validation function for PostCode
func validateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	return len(mobile) == 10 && strings.HasPrefix(mobile, "04")
}

// Custom validation function for State
func validateState(fl validator.FieldLevel) bool {
	state := fl.Field().String()
	states := []string{"VIC", "NSW", "ACT", "NT", "WA", "SA", "TAS", "QLD"}
	return slices.Contains(states, strings.ToUpper(state))
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

	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required and cannot be empty value", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	validate.RegisterTranslation("mobile", trans, func(ut ut.Translator) error {
		return ut.Add("mobile", "{0} must be valid Australian mobile number", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("mobile", fe.Field())
		return t
	})

	validate.RegisterTranslation("len", trans, func(ut ut.Translator) error {
		return ut.Add("len", "{0} should be exactly {1} characters", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("len", fe.Field(), fe.Param())
		return t
	})
}
