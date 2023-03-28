package controllers

import (
	"encoding/json"
	"net/http"
	
	"github.com/gorilla/mux"
	"listsystem/models"
	"listsystem/utils"
)

func GetAllHeads(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var heads []models.Head
	models.DB.Find(&heads)

	json.NewEncoder(w).Encode(heads)
}

func GetAllPages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var pages []models.Page
	models.DB.Find(&pages)

	json.NewEncoder(w).Encode(pages)
}

func GetHead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	key := mux.Vars(r)["key"]
	var head models.Head

	if err := models.DB.Where("key = ?", key).First(&head).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Query not found")
		return
	}

	json.NewEncoder(w).Encode(head)
}

func GetPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	key := mux.Vars(r)["key"]
	var page models.Page

	if err := models.DB.Where("key = ?", key).First(&page).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Query not found")
		return
	}

	json.NewEncoder(w).Encode(page)
}


