package Models

import (
	"github.com/jinzhu/gorm"
)

type Usuario struct {
	gorm.Model
	Usuario		string	`json:"usuario"`
	Nombre		string	`json:"nombre"`
	Apellido	string	`json:"apellido"`
}

func (u *Usuario) TableName() string {
	return "usuarios"
}
