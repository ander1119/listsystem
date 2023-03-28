package main

import (
	"net/http"

	"github.com/ander1119/listsystem/controllers"
	"github.com/ander1119/listsystem/models"
)

func main() {
	handler := controllers.New()

	server := &http.Server {
		Addr:	"127.0.0.1:5432",
		Handler: handler,
	}

	models.ConnectDatabase()

	server.ListenAndServe()
}