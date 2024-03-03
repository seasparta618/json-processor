# Project Overview

### Description
This project is a Go application that retrieves structured JSON input, validates it, logs it to STDOUT, and sends a request to an HTTP-based API with the input as part of the request message.

### Objective
The main objective of this project is to demonstrate the ability to work with JSON data, perform validation, interact with HTTP APIs, and handle errors effectively in Go.

### How to run the application

Note: In order to avoid some issue by reading the input from console, this app accepts the file path input with json content. If you do not want to define a json by yourself, you can using the sample json files inside assets.

1. Please clone this git repo, and enter into the root directory of the project.

2. To use the application, firstly please insure you have installed all the go dependency.

3. Please run `go build cmd/main.go`, you should see a runnable compiled main based on your system

4. As this application will `validate json then send it to an http api`, please clone this `gin-webapp` repo as well https://github.com/seasparta618/gin-webapp. Please spin it up to make sure the endpoint is reachable.

5. Please make sure the `API_HOST` in .env is correct with the spinned up `gin-webapp`, please try to assign a new `API_TOKEN` into .env if the application outputs `401 Unauthorized` in STDOUT.

6. Now should able to run the application to read the json, please run the following command
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

### Error Handling: 
Developing a strategy for handling and reporting errors throughout the application. The challenge is to handle errors gracefully, providing meaningful error messages, and avoiding application crashes.

### Testing: 
Writing tests to cover various scenarios and ensure the reliability of the application. The challenge is to create comprehensive and effective tests, especially for the parts of the application that interact with external APIs. 

### Code Maintainability: 
Writing clean, readable, and maintainable code. The challenge is to structure the code in a way that is easy to understand and modify in the future.

## Solution of the code challenge:

### Ambiguity in Requirements:
One of the initial challenges I faced with this code challenge was the ambiguity in the requirements, as it was not specified whether this was intended to be a Go web app or a Go command-line app. These two approaches have significant differences, especially when it comes to handling JSON input and validation. For a web app, reading and validating JSON input from a POST request body would be more straightforward, avoiding many of the IO errors that can occur with command-line input. However, with command-line input, issues such as EOF or line breaks are common. Therefore, I chose to accept file input to streamline the process and reduce the likelihood of encountering such errors.

### Project Structure and Design:
Dependency Injection:
In this project, I heavily utilized dependency injection, particularly evident in enquiry_service.go and json_service.go. This approach greatly facilitates testing, as it allows for the isolation of components and the injection of mock dependencies during test execution.

### Custom Validation Rules and User-Friendly Error Messages:
I implemented custom validation rules to ensure the integrity of the JSON input. Additionally, I focused on providing user-friendly error messages that are more understandable than the default "failed on xxx tag" messages. For example, errors are returned in a clear and concise format:
```json
{
  "Enquiry.EnquirerInfo.MobileNumber": "MobileNumber must be valid Australian mobile number",
  "Enquiry.EnquiryId": "EnquiryId should be exactly 32 characters",
  "Enquiry.EnquiryTitle": "EnquiryTitle is required and cannot be empty value"
}
```

### MVC Pattern:
Although this project does not follow the traditional web app structure, it still adheres to the Model-View-Controller (MVC) pattern. In this context, the main function acts as a controller, processing the incoming JSON data and delegating it to the service layer for further handling. The service layer, in turn, manages the business logic and interacts with the model layer, which represents the data structure. Finally, any errors encountered are simply printed to STDOUT, serving as the response to the user.

### Handling Environment Variables:
In general practice, the .env file should not be committed to version control as it often contains sensitive information such as API tokens, database credentials, and other secrets. Instead, it is advisable to provide a .env.template file with placeholder values, which can then be copied to a local .env file and filled with actual development or production values.

However, for the purpose of this code challenge, and to facilitate easier testing for evaluators, I have included the .env file with the necessary information directly in the repository. It is important to note that in a real-world scenario, sensitive information like API tokens should never be exposed in this manner and should be securely managed using appropriate tools and practices.