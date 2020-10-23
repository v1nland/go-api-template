package Controllers

import (
	"github.com/gin-gonic/gin"
	"go-api-template/ApiHelpers"
	"go-api-template/Models"
	"go-api-template/Repositories"
)

func ListUsuario(c *gin.Context) {
	var usr []Models.Usuario
	err := Repositories.GetAllUsuario(&usr)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usr)
	} else {
		ApiHelpers.RespondJSON(c, 200, usr)
	}
}

func AddNewUsuario(c *gin.Context) {
	var usr Models.Usuario
	c.BindJSON(&usr)
	err := Repositories.AddNewUsuario(&usr)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usr)
	} else {
		ApiHelpers.RespondJSON(c, 200, usr)
	}
}

func GetOneUsuario(c *gin.Context) {
	id := c.Params.ByName("id")
	var usr Models.Usuario
	err := Repositories.GetOneUsuario(&usr, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usr)
	} else {
		ApiHelpers.RespondJSON(c, 200, usr)
	}
}

func PutOneUsuario(c *gin.Context) {
	var usr Models.Usuario
	id := c.Params.ByName("id")
	err := Repositories.GetOneUsuario(&usr, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usr)
	}
	c.BindJSON(&usr)
	err = Repositories.PutOneUsuario(&usr, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usr)
	} else {
		ApiHelpers.RespondJSON(c, 200, usr)
	}
}

func DeleteUsuario(c *gin.Context) {
	var usr Models.Usuario
	id := c.Params.ByName("id")
	err := Repositories.DeleteUsuario(&usr, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usr)
	} else {
		ApiHelpers.RespondJSON(c, 200, usr)
	}
}
