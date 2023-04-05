# ListSystem
ListSystem is a RESTful API for managing lists of articles. It allows users to create, read, update, and delete lists and articles.

## Setup
1. Clone this repository to your local machine.
2. Install Go and set up your GOPATH.
3. Install the necessary dependencies by running go mod download.
4. Set up a PostgreSQL database and update the DATABASE_URL environment variable in the .env file with the connection details.
5. Start the server by running go run main.go.

## API Endpoints
### List Endpoints
`GET /lists` 
Retrieves a list of all lists.

# List System

List System is a RESTful API service that provides functionality to manage lists of articles. Each list contains a series of articles, which are stored in pages. 

## Getting Started

To get started, clone this repository and run the following commands:
```
$ go mod tidy
$ go run main.go
```

This will start the server on `http://localhost:8008`.

## API Documentation

The List System API (all with content-type: application/json) provides the following endpoints:

### `Get /heads`
Retrieves a list of all heads.
#### Response
```json
[
  {
    "list_key": "list1",
    "next_page_key": "page1",
    "create_at": 1630515839,
    "update_at": 1630515839
  },
  {
    "list_key": "list2",
    "next_page_key": null,
    "create_at": 1630515845,
    "update_at": 1630515845
  }
]
```
### `POST /head`
Creates a new head.


