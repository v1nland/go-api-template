package Models

import (
	"github.com/v1nland/ms-apiservice/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllUsuario(u *[]Usuario) (err error) {
	if err = Config.DB.Find(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewUsuario(u *Usuario) (err error) {
	if err = Config.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneUsuario(u *Usuario, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneUsuario(u *Usuario, id string) (err error) {
	fmt.Println(u)
	Config.DB.Save(u)
	return nil
}

func DeleteUsuario(u *Usuario, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(u)
	return nil
}
