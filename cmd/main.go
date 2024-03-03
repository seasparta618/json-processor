package main

import (
	"flag"
	"fmt"
	"json-processor/internal/model"
	"json-processor/internal/service"
	"json-processor/internal/util"
	"json-processor/internal/validation"
	"net/http"
	"os"

	config "json-processor/config"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	var jsonFilePath string
	flag.StringVar(&jsonFilePath, "file-path", "", "Path to the JSON file")
	flag.Parse()

	if jsonFilePath == "" {
		fmt.Println("Usage: json-processor --file-path <path to JSON file>")
		return
	}

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	// Get API_HOST from environment variables
	apiHost := os.Getenv("API_HOST")
	if apiHost == "" {
		fmt.Println("Error: API_HOST is not set in .env file")
		return
	}
	// Initialize configuration
	cfg := &config.Config{
		APIHost: apiHost,
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
	httpClient := &http.Client{} // Use the standard http.Client
	enquiryService := service.NewEnquiryService(jsonService, httpClient, cfg)
	err = enquiryService.SaveEnquiry([]byte(jsonData))
	if err != nil {
		fmt.Println("Failed to save enquiry:", err)
		return
	}
	fmt.Println("Successfully saved enquiry")
}
