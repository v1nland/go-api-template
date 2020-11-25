# go-api-template

A simple RESTful API compatible with heroku using:

1. **Gin Framework**
2. **Gorm**
3. **Go 1.15**

## Installation & Run

Download this project

```bash
$ git clone github.com/v1nland/go-api-template
```

To run locally, generate .env file and modify its variables

```bash
cp .env.example .env
```

Run and install dependencies

```bash
$ go run *.go
```

## API

#### /estudiantes

-   `GET` : Get all estudiantes
-   `POST` : Create a new estudiante

#### /estudiantes/:uuid

-   `GET` : Get a estudiante
-   `PUT` : Update a estudiante
-   `DELETE` : Delete a estudiante

#### /roles

-   `GET` : Get all rols
-   `POST` : Create a new rol

#### /roles/:uuid

-   `GET` : Get a rol
-   `PUT` : Update a rol
-   `DELETE` : Delete a rol
