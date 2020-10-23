# Go API template

A simple RESTful API compatible with heroku using:

1. **Gin Framework**
2. **Gorm**
3. **Go 1.13**

## Installation & Run

```bash
# Download this project
$ go get github.com/v1nland/go-api-template

# Download Gin Framework
$ go get github.com/gin-gonic/gin

# Download GORM
$ go get github.com/jinzhu/gorm
```

Setting DB in main.go

```go
Config.DB, err = gorm.Open("mysql", "DBUSERNAME:DBPASSWORD@tcp(127.0.0.1:3306)/DBNAME?charset=utf8&parseTime=True&loc=Local")
```

## Structure

```
├── ApiHelpers: Response interfaces
├── Config: Database or another services configuration
├── Controllers: GIN controller
├── Formats: Format functions (input/output)
├── Models: Database models
├── Repositories: Database repositories
├── RequestMessages: Input structures
├── ResponseMessages: Output structures
├── Routers: Routing handler
└── main.go
```

## API

#### /usuario

-   `GET` : Get all usuario
-   `POST` : Create a new usuario

#### /usuario/:id

-   `GET` : Get a usuario
-   `PUT` : Update a usuario
-   `DELETE` : Delete a usuario

# Usuario POST Params (JSON)

```
{
	"usuario": "test_username",
	"nombre": "test_name",
	"apellido": "test_lastname"
}
```
