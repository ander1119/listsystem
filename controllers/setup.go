package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	// the usage context is unreasonable, but still kept as basic api
	router.HandleFunc("/heads", GetAllHeads).Methods("GET")
	router.HandleFunc("/pages", GetAllPages).Methods("GET")

	router.HandleFunc("/head/{list_key}", GetHead).Methods("GET")
	router.HandleFunc("/page/{page_key}", GetPage).Methods("GET")

	router.HandleFunc("/head", InsertHead).Methods("POST")
	router.HandleFunc("/page", InsertPage).Methods("POST")

	router.HandleFunc("/resetHead", resetHead).Methods("POST")
	router.HandleFunc("/resetPage", resetPage).Methods("POST")
	
	router.HandleFunc("/deletePages", deletePages).Methods("POST")
	router.HandleFunc("/deleteHeads", deleteHeads).Methods("POST")

	return router
}