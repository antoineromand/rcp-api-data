package carcontroller

import (
	"encoding/json"
	"net/http"
	usecase "rcp-api-data/internal/usecase/data-collector"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type CarStatusController struct {
	DB *gorm.DB
}

func NewCarStatusController(db *gorm.DB) *CarStatusController {
	return &CarStatusController{
		DB: db,
	}
}

func (csc *CarStatusController) Controller() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "PUT":
			// Séparer l'URL pour obtenir le dernier segment
			urlParts := strings.Split(r.URL.Path, "/")
			action := urlParts[len(urlParts)-2] // Récupérer l'avant-dernier segment de l'URL
			param := urlParts[len(urlParts)-1]  // Récupérer le dernier segment de l'URL

			// Vérifier si l'action est valide
			if action != "activate" && action != "desactivate" {
				http.Error(w, "Invalid action", http.StatusBadRequest)
				return
			}

			// Vérifier si le paramètre est vide
			param = strings.TrimSpace(param)
			if param == "" {
				http.Error(w, "Invalid parameter", http.StatusBadRequest)
				return
			}

			paramUint, err := strconv.ParseUint(param, 10, 64)
			if err != nil {
				http.Error(w, "Invalid parameter", http.StatusBadRequest)
				return
			}

			// Appeler la fonction correspondante en fonction de l'action
			switch action {
			case "activate":
				csc.activate(uint(paramUint), w)
			case "desactivate":
				csc.desactivate(uint(paramUint), w)
			}
		}
	}
}

func (csc *CarStatusController) activate(id uint, w http.ResponseWriter) {
	usecase := usecase.NewActivateCarUserUseCase(csc.DB)
	response := usecase.ActivateCarUser(id)
	if !response.Success {
		http.Error(w, response.Message, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (csc *CarStatusController) desactivate(id uint, w http.ResponseWriter) {
	usecase := usecase.NewDesactivateCarUserUseCase(csc.DB)
	response := usecase.DesactivateCarUser(id)
	if !response.Success {
		http.Error(w, response.Message, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
