package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-api-template/Config"
	"go-api-template/Models"
	"go-api-template/Routers"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", "DBUSERNAME:DBPASSWORD@tcp(127.0.0.1:3306)/DBNAME?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Usuario{})

	r := Routers.SetupRouter()

	r.Run()
}
