# Project Overview

### Description
This project is a Go application that retrieves structured JSON input, validates it, logs it to STDOUT, and sends a request to an HTTP-based API with the input as part of the request message.

### Objective
The main objective of this project is to demonstrate the ability to work with JSON data, perform validation, interact with HTTP APIs, and handle errors effectively in Go.

### How to use the application

Note: In order to avoid some issue by reading the input from console, this app accepts the file path input with json content. If you do not want to define a json by yourself, you can using the sample json files inside assets.

1. To use the application, firstly please insure you have installed all the go dependency.

2. Please clone this git repo, and enter into the root directory of the project.

3. Please run `go build cmd/main.go`, you should see a main.(suffix) based on your system

4. To run the application, please run the following command
```bash
./main --file-path your-file-path
```

### Using sample json input
#### using valid json example
You can run the following command:
```bash
./main --file-path assets/sample-valid.json
```

This should expect the following output in STDOUT, it will come with the standard json output in good format in STDOUT, then it should come with another output, which is the result of sending to the api.
```json
{
  "enquiryId": "123e4567e89b12d3a456426614174000",
  "enquiryTitle": "Property Enquiry from ....",
  "enquiryDate": "2023-04-07T12:00:00Z",
  "enquirerInfo": {
    "firstName": "John",
    "lastName": "Doe",
    "mobileNumber": "0412345678",
    "emailAddress": "john.doe@example.com"
  },
  "properties": [
    {
      "propertyId": "123e4567e89b12d3a456426614174000",
      "propertyAddress": {
        "streetName": "Main Street",
        "streetNumber": "123",
        "unitNumber": "2A",
        "suburbName": "Springfield",
        "postCode": 2000,
        "state": "NSW"
      }
    }
  ]
}
```
#### using invalid json example
You can run the following command:
```bash
./main --file-path assets/sample-invalid.json
```
You should expect to see the following result in the STDOUT
```bash
Failed to save enquiry: validation failed: field validation failed: {
  "Enquiry.EnquirerInfo.MobileNumber": "MobileNumber must be valid Australian mobile number",
  "Enquiry.EnquiryId": "EnquiryId should be exactly 32 characters",
  "Enquiry.EnquiryTitle": "EnquiryTitle is required and cannot be empty value"
}
```
This is because there are some translation of the errors from validation library.

From the requirement, it can be believed that one of the challenge for this assignment is assess your ability to validate json and provide friendly and error reasonable based on the input json.

## Challenges and Considerations

### JSON Validation: 
- Implementing comprehensive validation rules to ensure that only valid JSON input is processed.
- The validation process should be elegant and efficient, even for large JSON structures, with as few error returns as possible. 
- It is not hard to parse and unmarshal the json, but it might be challengable to build the json parsing progress friendly & transparently to the user and return reasonable and meaningful errors at the same time. It is not a good practice each time just return 1 error then return another one at next time, ideally, the parsing progress should be able to return all the errors at once.

### HTTP API Interaction: 
Handling the communication with an external HTTP API, including constructing the request, processing the response, and managing errors. The challenge is to ensure robust and reliable API interaction, especially in the face of network issues or API downtime.

### Logging: 
Deciding what information to log and how to format it for readability and usefulness. The challenge is to log enough information to be useful for debugging and monitoring, without overwhelming the log files with unnecessary details.

### Error Handling: 
Developing a strategy for handling and reporting errors throughout the application. The challenge is to handle errors gracefully, providing meaningful error messages, and avoiding application crashes.

### Testing: 
Writing tests to cover various scenarios and ensure the reliability of the application. The challenge is to create comprehensive and effective tests, especially for the parts of the application that interact with external APIs. 

### Code Maintainability: 
Writing clean, readable, and maintainable code. The challenge is to structure the code in a way that is easy to understand and modify in the future.