# List System

List System is a RESTful API service that provides functionality to manage lists of articles. Each list contains a series of articles, which are stored in pages. 

## Getting Started
### Prerequisites
Before running the service, you will need the following:
- Go 1.16 or higher
- PostgreSQL 9.6 or higher
### Configuration
The service is configured using environment variables. The following variables must be set:
- `PGUSER`
- `PGPASSWORD`
- `PGDATABASE`
- `PGHOST`

### Usage
To use this service, clone this repository and run the following commands:
```
$ go mod tidy
$ go run main.go
```

This will start the server on `http://localhost:8008`, and once the service is running, you can use the following endpoints to interact with the system:

***All requests and responses are in JSON format.***

#### Head Endpoints
- `GET /heads`: get a list of all the heads
- `GET /head/{list_key}`: get the details of a specific head
- `POST /head`: create a new head, and its list key would be returned within response. Request body as below
```json
{
    "next_page_key": "other-uuid-string"
}
```
- `DELETE /deleteHeads?keep={sec}`: delete heads that are older than keep (in sec)
- `DELETE /resetHeads`: delete heads that are older than one day

#### Page Endpoints
- `GET /pages`: get a list of all the pages
- `GET /page/{page_key}`: get the details of a specific page
- `POST /page`: create a new page, and its page key would be returned within response. Request body as below
```json
{
    "articles": "hello world!!!",
    "next_page_key": "other-uuid-string"
}
```
- `DELETE /deletePages?keep={sec}`: delete pages that are older than keep (in sec)
- `DELETE /resetPages`: delete pages that are older than one day

### Unit Test
To run the tests, run the following command:
```
$ go test ./..
```

#### About
##### Reasons to Choose PostgreSQL
In my opinion, there's no much difference between choosing PostgreSQL and other database, since it's hard to simulate real world use case with tremendous users. But if we encounter bottleneck in future, PostgreSQL provides interface for us to conveniently analyze and catch out which transaction brought out overheads

##### How to efficiently delete numerous head/page
According to https://gorm.io/docs/delete.html and https://dba.stackexchange.com/questions/34864/most-efficient-way-of-bulk-deleting-rows-from-postgres, performing delete operation with primary key is much more faster. So within delete api I select list_key/page_key from head/page whose creat_at is older than boundary, then delete with primary key condition.