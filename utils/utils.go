package utils

import (
    // "errors"
    "net/http"
    "regexp"
    "strings"
    "github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateInput validates struct fields using tags
func ValidateInput(input interface{}) error {
    return validate.Struct(input)
}

// SanitizeInput removes potentially harmful characters from input
func SanitizeInput(input string) string {
    input = strings.TrimSpace(input)
    input = strings.ToLower(input)
    return regexp.MustCompile(`[^a-zA-Z0-9\-_ ]+`).ReplaceAllString(input, "")
}

// HandleError writes an error response to the client
func HandleError(w http.ResponseWriter, statusCode int, message string) {
    w.WriteHeader(statusCode)
    w.Write([]byte(`{"error":"` + message + `"}`))
}

// IsValidID checks if an ID is in a valid format (e.g., UUID)
func IsValidID(id string) bool {
    return regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`).MatchString(id)
}