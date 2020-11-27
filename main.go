package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	swagger_files "github.com/swaggo/files"
	gin_swagger "github.com/swaggo/gin-swagger"
	"go-api-template/Config"
	"go-api-template/Models"
	"go-api-template/Routers"
	"go-api-template/docs"
	"os"
)

var err error

// @title MedRoom API
// @version 1.0
// @description Swagger definition for MedRoom backend.
// @termsOfService https://swagger.io/terms/

// @contact.name Martín Saavedra
// @contact.url https://www.swagger.io/support
// @contact.email martin.saavedra@mail.udp.cl

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("status error: ", err)
	}

	connection_string := os.Getenv("POSTGRESQL_CONNECTION_STRING")
	should_automigrate := os.Getenv("SHOULD_AUTOMIGRATE")
	swagger_protocol := os.Getenv("SWAGGER_PROTOCOL")
	swagger_host := os.Getenv("SWAGGER_HOST")

	Config.DB, err = gorm.Open("postgres", connection_string)
	if err != nil {
		fmt.Println("status error: ", err)
	}
	defer Config.DB.Close()

	if should_automigrate == "1" {
		// // Inicialización de tablas en DB
		Config.DB.AutoMigrate(&Models.Rol{}, &Models.Estudiante{})

		// // Las FK de las relaciones N - N con datos adicionales se deben aplicar manualmente de la siguiente forma
		// // Para relaciones 1 - N y 1 - 1, se deben especificar únicamente en el Model.go
		// Config.DB.Model(&Models.PedidoProducto{}).AddForeignKey("pedido_id", "public.pedidos(id)", "CASCADE", "CASCADE")
		// Config.DB.Model(&Models.PedidoProducto{}).AddForeignKey("producto_id", "public.productos(id)", "CASCADE", "CASCADE")
	}

	r := Routers.SetupRouter()

	// auto swagger configuration
	docs.SwaggerInfo.Host = swagger_host
	url := gin_swagger.URL(swagger_protocol + "://" + swagger_host + "/docs/v1/doc.json")
	r.GET("/docs/v1/*any", gin_swagger.WrapHandler(swagger_files.Handler, url))

	r.Run()
}
