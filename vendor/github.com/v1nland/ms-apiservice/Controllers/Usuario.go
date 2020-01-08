package Controllers

import (
	"github.com/v1nland/ms-apiservice/ApiHelpers"
	"github.com/v1nland/ms-apiservice/Models"
	"github.com/gin-gonic/gin"
)

func ListUsuario(c *gin.Context) {
	var usr []Models.Usuario
	err := Models.GetAllUsuario(&usr)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usr)
	} else {
		ApiHelpers.RespondJSON(c, 200, usr)
	}
}

func AddNewUsuario(c *gin.Context) {
	var usr Models.Usuario
	c.BindJSON(&usr)
	err := Models.AddNewUsuario(&usr)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usr)
	} else {
		ApiHelpers.RespondJSON(c, 200, usr)
	}
}

func GetOneUsuario(c *gin.Context) {
	id := c.Params.ByName("id")
	var usr Models.Usuario
	err := Models.GetOneUsuario(&usr, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usr)
	} else {
		ApiHelpers.RespondJSON(c, 200, usr)
	}
}

func PutOneUsuario(c *gin.Context) {
	var usr Models.Usuario
	id := c.Params.ByName("id")
	err := Models.GetOneUsuario(&usr, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usr)
	}
	c.BindJSON(&usr)
	err = Models.PutOneUsuario(&usr, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usr)
	} else {
		ApiHelpers.RespondJSON(c, 200, usr)
	}
}

func DeleteUsuario(c *gin.Context) {
	var usr Models.Usuario
	id := c.Params.ByName("id")
	err := Models.DeleteUsuario(&usr, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usr)
	} else {
		ApiHelpers.RespondJSON(c, 200, usr)
	}
}
