package main

import (
	"ms-apiservice/Config"
	"ms-apiservice/Models"
	"ms-apiservice/Routers"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
 	Config.DB, err = gorm.Open("mysql", "b67537abb72605:1b4e72ef@(us-cdbr-iron-east-05.cleardb.net)/heroku_58752f0203bba17?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Book{})
	Config.DB.AutoMigrate(&Models.Usuario{})

	r := Routers.SetupRouter()

	r.Run()
}
