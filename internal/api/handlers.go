package api

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"optiflow/internal/models"
	"optiflow/internal/services"
	"optiflow/internal/utils"
	"optiflow/internal/oauth"
)

func SetupRoutes(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/edi", createEDIHandler(db)).Methods("POST")
	router.HandleFunc("/edi/{id}", getEDIHandler(db)).Methods("GET")
}

func createEDIHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.EDIRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.HandleError(w, http.StatusBadRequest, "Invalid input")
			return
		}

		// Validate input
		if err := utils.ValidateInput(req); err != nil {
			utils.HandleError(w, http.StatusBadRequest, err.Error())
			return
		}

		// Convert request to EDI model
		edi := models.FromRequest(req)

		// Create EDI
		service := services.EDIService{DB: db}
		if err := service.CreateEDI(edi); err != nil {
			utils.HandleError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Return response
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(edi.ToResponse())
	}
}

func getEDIHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		// Validate ID
		if !utils.IsValidID(id) {
			utils.HandleError(w, http.StatusBadRequest, "Invalid ID format")
			return
		}

		// Fetch EDI from DB
		service := services.EDIService{DB: db}
		edi, err := service.GetEDI(id)
		if err != nil {
			if err.Error() == "EDI not found" {
				utils.HandleError(w, http.StatusNotFound, err.Error())
			} else {
				utils.HandleError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		// Return response
		json.NewEncoder(w).Encode(edi.ToResponse())
	}
}

func LoginHandler(db *sql.DB, secretKey string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var credentials struct {
            Username string `json:"username"`
            Password string `json:"password"`
        }
        if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }

        // Validate credentials
        user, err := auth.ValidateCredentials(db, credentials.Username, credentials.Password)
        if err != nil {
            http.Error(w, "Invalid username or password", http.StatusUnauthorized)
            return
        }

        // Generate a JWT token
        token, err := auth.GenerateToken(user, secretKey, 1*time.Hour)
        if err != nil {
            http.Error(w, "Failed to generate token", http.StatusInternalServerError)
            return
        }

        // Return the token
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "token": token,
        })
    }
}
