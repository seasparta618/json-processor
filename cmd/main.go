package main

import (
	"flag"
	"fmt"
	apperrors "json-processor/internal/error"
	"json-processor/internal/model"
	"json-processor/internal/service"
	"json-processor/internal/util"
	"json-processor/internal/validation"

	"github.com/go-playground/locales/en"
	universal_translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
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
	uni := universal_translator.New(enLocale, enLocale)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)
	validation.RegisterCustomValidations(validate)

	// Create JSON service
	jsonService := service.JSONService[model.Enquiry]{
		Validator: validate,
	}

	// Create processor service
	processorService := service.NewProcessorService(&jsonService)

	// Read JSON input from the provided file path
	jsonData, err := util.ReadFile(jsonFilePath)
	if err != nil {
		fmt.Println(apperrors.ErrorFileNotFound, err)
		return
	}

	if len(jsonData) == 0 {
		fmt.Println(apperrors.ErrorEmptyContent)
		return
	}

	// Process data
	err = processorService.ProcessData([]byte(jsonData))
	if err != nil {
		fmt.Println("Error processing data:", "\n", err)
		return
	}

	fmt.Println("Data processed successfully!")
}
