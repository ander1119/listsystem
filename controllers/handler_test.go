package controllers

import (
	"encoding/json"
	"fmt"
	"listsystem/models"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func TestGetAllHeads(t *testing.T) {
	// connect to database
	models.ConnectDatabase()

    // create a new http request
    req, err := http.NewRequest("GET", "/heads", nil)
    if err != nil {
        t.Fatal(err)
    }

    // create a new ResponseRecorder to record the http response
    rr := httptest.NewRecorder()

	// Call the GetAllHeads function with the mock request and response
	router := mux.NewRouter()
	router.HandleFunc("/heads", GetAllHeads).Methods("GET")
	router.ServeHTTP(rr, req)

    // check the status code returned by the router
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("router returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // check the content type header returned by the router
    if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
        t.Errorf("router returned wrong content type: got %v want %v",
            contentType, "application/json")
    }

    // check the json response body returned by the router
    expected := []models.Head{}
    json.Unmarshal(rr.Body.Bytes(), &expected)

    var actual []models.Head
    models.DB.Find(&actual)

    if !reflect.DeepEqual(expected, actual) {
        t.Errorf("router returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestGetAllPages(t *testing.T) {
	// connect to database
	models.ConnectDatabase()

    // create a new http request
    req, err := http.NewRequest("GET", "/pages", nil)
    if err != nil {
        t.Fatal(err)
    }

    // create a new ResponseRecorder to record the http response
    rr := httptest.NewRecorder()

	// Call the GetAllPages function with the mock request and response
	router := mux.NewRouter()
	router.HandleFunc("/pages", GetAllPages).Methods("GET")
	router.ServeHTTP(rr, req)

    // check the status code returned by the router
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("router returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // check the content type header returned by the router
    if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
        t.Errorf("router returned wrong content type: got %v want %v",
            contentType, "application/json")
    }

    // check the json response body returned by the router
    expected := []models.Page{}
    json.Unmarshal(rr.Body.Bytes(), &expected)

    var actual []models.Page
    models.DB.Find(&actual)

    if !reflect.DeepEqual(expected, actual) {
        t.Errorf("router returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestGetHead(t *testing.T) {
	// connect to database
	models.ConnectDatabase()

	// Create a new test head and insert it into the database
	testHead := models.Head{
		ListKey: uuid.NewString(), 
		NextPageKey: uuid.NewString(),
		CreateAt: time.Now().Unix(),
		UpdateAt: time.Now().Unix(),
	}

	if err := models.DB.Create(&testHead).Error; err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the test head's key as a parameter
	req, err := http.NewRequest("GET", fmt.Sprintf("/head/%s", testHead.ListKey), nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Call the GetHead function with the mock request and response
	router := mux.NewRouter()
	router.HandleFunc("/head/{list_key}", GetHead).Methods("GET")
	router.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("router returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response content type
	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("router returned wrong content type: got %v want %v",
			contentType, "application/json")
	}

	// Check the response body
	var responseHead models.Head
	if err := json.Unmarshal(rr.Body.Bytes(), &responseHead); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(testHead, responseHead) {
        t.Errorf("router returned unexpected body: got %v want %v",
            rr.Body.String(), responseHead)
    }
}

func TestGetPage(t *testing.T) {
	// connect to database
	models.ConnectDatabase()

	// Create a new test head and insert it into the database
	testPage := models.Page{
		PageKey: uuid.NewString(), 
		Articles: "hello world",
		NextPageKey: nil,
		CreateAt: time.Now().Unix(),
		UpdateAt: time.Now().Unix(),
	}

	if err := models.DB.Create(&testPage).Error; err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the test page's key as a parameter
	req, err := http.NewRequest("GET", fmt.Sprintf("/page/%s", testPage.PageKey), nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Call the GetPage function with the mock request and response
	router := mux.NewRouter()
	router.HandleFunc("/page/{page_key}", GetPage).Methods("GET")
	router.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("router returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response content type
	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("router returned wrong content type: got %v want %v",
			contentType, "application/json")
	}

	// Check the response body
	var responsePage models.Page
	if err := json.Unmarshal(rr.Body.Bytes(), &responsePage); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(testPage, responsePage) {
        t.Errorf("router returned unexpected body: got %v want %v",
            rr.Body.String(), responsePage)
    }
}