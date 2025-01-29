package api

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"optiflow/internal/models"
	"optiflow/internal/services"
	"optiflow/internal/utils"
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
