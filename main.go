package main

import (
	"net/http"

	"listsystem/controllers"
	"listsystem/models"
)

func main() {
	handler := controllers.New()

	server := &http.Server {
		Addr:	"0.0.0.0:8008",
		Handler: handler,
	}

	models.ConnectDatabase()

	server.ListenAndServe()
}