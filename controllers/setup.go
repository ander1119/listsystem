package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	// the usage context is unreasonable, but still kept as basic api
	router.HandleFunc("/heads", GetAllHeads).Method("GET")
	router.HandleFunc("/pages", GetAllPages).Method("GET")

	router.HandleFunc("/head/{key}", GetHead).Method("GET")
	router.HandleFunc("/page/{key}", GetPage).Method("GET")

	// router.HandleFunc("")

	return router
}