package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"go-api-template/Config"
	"go-api-template/Models"
	"go-api-template/Routers"
	"os"
)

var err error

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("status error: ", err)
	}

	db_service := os.Getenv("DB_SERVICE")
	connection_string := os.Getenv("DB_CONNECTION_STRING")
	shouldAutomigrate := os.Getenv("SHOULD_AUTOMIGRATE")

	Config.DB, err = gorm.Open(db_service, connection_string)
	if err != nil {
		fmt.Println("status: ", err)
	}
	defer Config.DB.Close()

	if shouldAutomigrate == "1" {
		Config.DB.AutoMigrate(&Models.Usuario{})
	}

	r := Routers.SetupRouter()

	r.Run()
}
