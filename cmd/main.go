package main

import (
	"flag"
	"fmt"
	"json-processor/internal/model"
	"json-processor/internal/service"
	"json-processor/internal/util"
	"json-processor/internal/validation"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func main() {
	var jsonFilePath string
	flag.StringVar(&jsonFilePath, "file-path", "", "Path to the JSON file")
	flag.Parse()

	if jsonFilePath == "" {
		fmt.Println("Usage: json-processor --file-path <path to JSON file>")
		return
	}

	// Set up validator and translator
	validate := validator.New()
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	trans, _ := uni.GetTranslator("en")

	// Register custom validations
	validation.RegisterCustomValidations(validate, trans)

	// Create JSON service with translator
	jsonService := service.NewJSONService[model.Enquiry](validate, trans)

	// Read JSON input from the provided file path
	jsonData, err := util.ReadFile(jsonFilePath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Validate JSON
	_, err = jsonService.ValidateJSON([]byte(jsonData))
	if err != nil {
		fmt.Println("Error validating JSON:", err)
		return
	}

	fmt.Println("JSON validation passed!")
}
