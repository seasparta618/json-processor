package validation

import (
	"testing"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func TestValidatePostCode(t *testing.T) {
	validate := validator.New()
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	trans, _ := uni.GetTranslator("en")
	RegisterCustomValidations(validate, trans)

	tests := []struct {
		postCode int
		want     bool
	}{
		{2000, true},
		{1500, false},
		{3000, true},
		{8000, false},
	}

	for _, tt := range tests {
		err := validate.Var(tt.postCode, "postcode")
		got := err == nil
		if got != tt.want {
			t.Errorf("validatePostCode(%d) = %v, want %v", tt.postCode, got, tt.want)
		} else {
			t.Logf("validated (%d) with result %v passed", tt.postCode, tt.want)
		}
	}
}

func TestValidateState(t *testing.T) {
	validate := validator.New()
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	trans, _ := uni.GetTranslator("en")
	RegisterCustomValidations(validate, trans)

	tests := []struct {
		state string
		want  bool
	}{
		{"VIC", true},
		{"NSW", true},
		{"XYZ", false},
		{"ACT", true},
		{"NSS", false},
	}

	for _, tt := range tests {
		err := validate.Var(tt.state, "state")
		got := err == nil
		if got != tt.want {
			t.Errorf("validateState(%q) = %v, want %v", tt.state, got, tt.want)
		} else {
			t.Logf("validated (%s) with result %v passed", tt.state, tt.want)
		}
	}
}
