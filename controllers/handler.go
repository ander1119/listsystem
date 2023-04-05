package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"listsystem/models"
	"listsystem/utils"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetAllHeads(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var heads []models.Head
	models.DB.Find(&heads)

	// return http response with struct value (json)
	json.NewEncoder(w).Encode(heads)
}

func GetAllPages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var pages []models.Page
	models.DB.Find(&pages)

	// return http response with struct value (json)
	json.NewEncoder(w).Encode(pages)
}

func GetHead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	key := mux.Vars(r)["list_key"]
	var head models.Head
	if err := models.DB.Where("list_key = ?", key).First(&head).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Query not found")
		return
	}

	// return http response with struct value (json)
	json.NewEncoder(w).Encode(head)
}

func GetPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	key := mux.Vars(r)["page_key"]
	var page models.Page

	if err := models.DB.Where("page_key = ?", key).First(&page).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Query not found")
		return
	}

	// return http response with struct value (json)
	json.NewEncoder(w).Encode(page)
}

func InsertHead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input models.HeadInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)
	validate := validator.New()
 	
	// preventing missing fields from head
	if err := validate.Struct(input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return 
	}

	// preventing no corresponding page in page table
	if err := models.DB.First(&models.Page{}, "page_key = ?", input.NextPageKey).Error; err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return 
	}

	head := &models.Head{
		ListKey: input.ListKey,
		NextPageKey: input.NextPageKey,
		CreateAt: time.Now().Unix(),
		UpdateAt: time.Now().Unix(),
	}
	
	result := models.DB.Create(head)

	if result.Error != nil {
		utils.RespondWithError(w, http.StatusBadRequest, result.Error.Error())
		return
	}

	// return http response with struct value (json)
	json.NewEncoder(w).Encode(head) 
}

func InsertPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input models.PageInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)
	validate := validator.New()

	if err := validate.Struct(input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation Error" + err.Error())
		return 
	}

	var nextPageKey *string = nil
	if input.NextPageKey != "" {
		nextPageKey = &input.NextPageKey
	}

	page := models.Page{
		PageKey:     uuid.NewString(),
		Articles:    input.Articles,
		NextPageKey: nextPageKey,
		CreateAt:    time.Now().Unix(),
		UpdateAt:    time.Now().Unix(),
	}

	if input.NextPageKey != "" {
		err := models.DB.First(&models.Page{}, "page_key = ?", input.NextPageKey).Error
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return 
		}
		page.NextPageKey = &input.NextPageKey
	}

	if err := models.DB.Create(page).Error; err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return 
	} 

	// return http response with struct value (json)
	json.NewEncoder(w).Encode(page) 
}

func resetHead(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")
	before := time.Now().Round(time.Hour * 24).Add(-time.Hour * 24).Unix()

	// var toDeleteHeads []models.Head
	// if err := models.DB.Where("create_at <= ?", ber)

	result := models.DB.Where("create_at <= ?", before).Delete(&models.Head{})

	if result.Error != nil {
		utils.RespondWithError(w, http.StatusBadRequest, result.Error.Error())
		return
	}

	// return http response with struct value (json)
  	json.NewEncoder(w).Encode(result.RowsAffected)
}

func resetPage(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")

	before := time.Now().Round(time.Hour * 24).Add(-time.Hour * 24).Unix()

	result := models.DB.Where("create_at <= ?", before).Delete(&models.Page{})

	if result.Error != nil {
		utils.RespondWithError(w, http.StatusBadRequest, result.Error.Error())
		return
	}

	// return http response with struct value (json)
  	json.NewEncoder(w).Encode(result.RowsAffected)
}

func deletePages(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")

	keepSeconds, err := strconv.Atoi(r.FormValue("keep"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "unvalid data type for keep, should be int")
		return
	}

	before := time.Now().Add(-time.Second * time.Duration(keepSeconds)).Unix()
	
	result := models.DB.Where("create_at <= ?", before).Delete(&models.Page{})

	if result.Error != nil {
		utils.RespondWithError(w, http.StatusBadRequest, result.Error.Error())
		return
	}

	// return http response with struct value (json)
  	json.NewEncoder(w).Encode(map[string]int{"affectedRows": int(result.RowsAffected)})
}

func deleteHeads(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")

	keepSeconds, err := strconv.Atoi(r.FormValue("keep"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "unvalid data type for keep, should be int")
		return
	}

	before := time.Now().Add(-time.Second * time.Duration(keepSeconds)).Unix()
	
	result := models.DB.Where("create_at <= ?", before).Delete(&models.Head{})

	if result.Error != nil {
		utils.RespondWithError(w, http.StatusBadRequest, result.Error.Error())
		return
	}

	// return http response with struct value (json)
  	json.NewEncoder(w).Encode(map[string]int{"affectedRows": int(result.RowsAffected)})
}
