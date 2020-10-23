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

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASS")
	shouldAutomigrate := os.Getenv("SHOULD_AUTOMIGRATE")

	Config.DB, err = gorm.Open("postgres", "host="+host+" port="+port+" user="+user+" dbname="+dbname+" password="+password)
	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config.DB.Close()

	if shouldAutomigrate == "1" {
		Config.DB.AutoMigrate(&Models.Usuario{})
	}

	r := Routers.SetupRouter()

	r.Run()
}
