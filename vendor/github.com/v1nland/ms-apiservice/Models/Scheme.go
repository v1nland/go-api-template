package Models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name     string `json:"name"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

func (b *Book) TableName() string {
	return "book"
}

type Usuario struct {
	gorm.Model
	Usuario		string	`json:"usuario"`
	Nombre		string	`json:"nombre"`
	Apellido	string	`json:"apellido"`
}

func (u *Usuario) TableName() string {
	return "usuarios"
}
