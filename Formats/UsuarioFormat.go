package Formats

import (
	"go-api-template/RequestMessages"
	"go-api-template/Utils"
	"strings"
)

func ListUsuarioInput(u RequestMessages.ListUsuarioPayload) {

}

func AddNewUsuarioInput(u *RequestMessages.AddNewUsuarioPayload) {
	u.Usuario = strings.TrimSpace(u.Usuario)
	u.Nombre = strings.TrimSpace(u.Nombre)
	u.Apellido = strings.TrimSpace(u.Apellido)

	u.Usuario = strings.ToUpper(u.Usuario)
	u.Nombre = strings.ToUpper(u.Nombre)
	u.Apellido = strings.ToUpper(u.Apellido)

	u.Usuario = Utils.RemoveAccents(u.Usuario)
	u.Nombre = Utils.RemoveAccents(u.Nombre)
	u.Apellido = Utils.RemoveAccents(u.Apellido)
}

func GetOneUsuarioInput(u RequestMessages.GetOneUsuarioPayload) {

}

func PutOneUsuarioInput(u RequestMessages.PutOneUsuarioPayload) {
	u.Usuario = strings.TrimSpace(u.Usuario)
	u.Nombre = strings.TrimSpace(u.Nombre)
	u.Apellido = strings.TrimSpace(u.Apellido)

	u.Usuario = strings.ToUpper(u.Usuario)
	u.Nombre = strings.ToUpper(u.Nombre)
	u.Apellido = strings.ToUpper(u.Apellido)

	u.Usuario = Utils.RemoveAccents(u.Usuario)
	u.Nombre = Utils.RemoveAccents(u.Nombre)
	u.Apellido = Utils.RemoveAccents(u.Apellido)
}

func DeleteUsuarioInput(u RequestMessages.DeleteUsuarioPayload) {

}
