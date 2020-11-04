package Repositories

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-api-template/Config"
	"go-api-template/Models"
)

func GetAllUsuario(u *[]Models.Usuario) (err error) {
	if err = Config.DB.Find(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewUsuario(u *Models.Usuario) (err error) {
	if err = Config.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneUsuario(u *Models.Usuario, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneUsuario(u *Models.Usuario, id string) (err error) {
	fmt.Println(u)
	Config.DB.Save(u)
	return nil
}

func DeleteUsuario(u *Models.Usuario, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(u)
	return nil
}
