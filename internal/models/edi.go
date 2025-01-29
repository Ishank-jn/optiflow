package models

import (
	"time"
)

// EDI represents an Electronic Data Interchange transaction
type EDI struct {
	ID        string    `json:"id" validate:"required,uuid4"` // Unique identifier for the EDI transaction
	Data      string    `json:"data" validate:"required"`     // EDI data (e.g., XML, JSON, or plain text)
	CreatedAt time.Time `json:"created_at"`                   // Timestamp when the EDI was created
	UpdatedAt time.Time `json:"updated_at"`                   // Timestamp when the EDI was last updated
}

// EDIRequest represents the request payload for creating or updating an EDI
type EDIRequest struct {
	Data string `json:"data" validate:"required"` // EDI data (e.g., XML, JSON, or plain text)
}

// EDIResponse represents the response payload for an EDI transaction
type EDIResponse struct {
	ID        string    `json:"id"`         // Unique identifier for the EDI transaction
	Data      string    `json:"data"`       // EDI data (e.g., XML, JSON, or plain text)
	CreatedAt time.Time `json:"created_at"` // Timestamp when the EDI was created
	UpdatedAt time.Time `json:"updated_at"` // Timestamp when the EDI was last updated
}

// ToResponse converts an EDI model to an EDIResponse
func (e *EDI) ToResponse() EDIResponse {
	return EDIResponse{
		ID:        e.ID,
		Data:      e.Data,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

// FromRequest converts an EDIRequest to an EDI model
func FromRequest(req EDIRequest) EDI {
	return EDI{
		ID:        generateUUID(), // Generate a unique UUID for the EDI
		Data:      req.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// generateUUID generates a unique UUID for the EDI
func generateUUID() string {
	// Use a proper UUID library in production (e.g., github.com/google/uuid)
	return "unique-id-" + time.Now().Format("20060102150405") // Placeholder for demonstration
}
