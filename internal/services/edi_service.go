package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"optiflow/cache"
	"optiflow/internal/models"
	"time"
)

type EDIService struct {
	DB *sql.DB
}

// CreateEDI creates a new EDI transaction
func (s *EDIService) CreateEDI(edi models.EDI) error {
	// Validate input
	if edi.Data == "" {
		return errors.New("EDI data cannot be empty")
	}

	// Insert into database
	query := `INSERT INTO edi (id, data, created_at) VALUES ($1, $2, $3)`
	_, err := s.DB.Exec(query, edi.ID, edi.Data, time.Now())
	if err != nil {
		log.Printf("Failed to create EDI: %v", err)
		return err
	}

	// Cache the result
	jsonData, _ := json.Marshal(edi)
	cache.Set("edi_"+edi.ID, jsonData, 10*time.Minute)

	return nil
}

// GetEDI retrieves an EDI transaction by ID
func (s *EDIService) GetEDI(id string) (*models.EDI, error) {
	// Check cache first
	cachedData, err := cache.Get("edi_" + id)
	if err == nil {
		var edi models.EDI
		if err := json.Unmarshal([]byte(cachedData), &edi); err == nil {
			return &edi, nil
		}
	}

	// Fetch from database
	var edi models.EDI
	query := `SELECT id, data, created_at FROM edi WHERE id = $1`
	row := s.DB.QueryRow(query, id)
	err = row.Scan(&edi.ID, &edi.Data, &edi.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("EDI not found")
		}
		log.Printf("Failed to fetch EDI: %v", err)
		return nil, err
	}

	// Cache the result
	jsonData, _ := json.Marshal(edi)
	cache.Set("edi_"+id, jsonData, 10*time.Minute)

	return &edi, nil
}

// UpdateEDI updates an existing EDI transaction
func (s *EDIService) UpdateEDI(id string, newData string) error {
	// Validate input
	if newData == "" {
		return errors.New("EDI data cannot be empty")
	}

	// Update database
	query := `UPDATE edi SET data = $1 WHERE id = $2`
	_, err := s.DB.Exec(query, newData, id)
	if err != nil {
		log.Printf("Failed to update EDI: %v", err)
		return err
	}

	// Invalidate cache
	cache.Delete("edi_" + id)

	return nil
}

// DeleteEDI deletes an EDI transaction by ID
func (s *EDIService) DeleteEDI(id string) error {
	// Delete from database
	query := `DELETE FROM edi WHERE id = $1`
	_, err := s.DB.Exec(query, id)
	if err != nil {
		log.Printf("Failed to delete EDI: %v", err)
		return err
	}

	// Invalidate cache
	cache.Delete("edi_" + id)

	return nil
}
