package Controllers

import (
	"github.com/gin-gonic/gin"
	"go-api-template/ApiHelpers"
	"go-api-template/Formats"
	"go-api-template/Models"
	"go-api-template/Repositories"
	"go-api-template/RequestMessages"
)

/*
	*
	*  FUNCIÓN ListUsuario
	*
    *
	*
	*
    *
*/
func ListUsuario(c *gin.Context) {
	// model container
	var usr []Models.Usuario

	// query
	err := Repositories.GetAllUsuario(&usr)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usr)
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, usr)
}

/*
	*
	*  FUNCIÓN AddNewUsuario
	*
    *
	*
	*
    *
*/
func AddNewUsuario(c *gin.Context) {
	// input container
	var usr RequestMessages.AddNewUsuarioPayload

	// input bind
	c.BindJSON(&usr)

	// format input
	Formats.AddNewUsuarioInput(&usr)

	// generate model entity
	new_usr := Models.Usuario{
		Usuario:  usr.Usuario,
		Nombre:   usr.Nombre,
		Apellido: usr.Apellido,
	}

	// query
	err := Repositories.AddNewUsuario(&new_usr)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usr)
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, usr)
}

/*
	*
	*  FUNCIÓN GetOneUsuario
	*
    *
	*
	*
    *
*/
func GetOneUsuario(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var usr Models.Usuario

	// query
	err := Repositories.GetOneUsuario(&usr, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usr)
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, usr)
}

/*
	*
	*  FUNCIÓN PutOneUsuario
	*
    *
	*
	*
    *
*/
func PutOneUsuario(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var usr Models.Usuario

	// get query
	err := Repositories.GetOneUsuario(&usr, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usr)
	}

	// input bind
	c.BindJSON(&usr)

	// put query
	err = Repositories.PutOneUsuario(&usr, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usr)
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, usr)
}

/*
	*
	*  FUNCIÓN DeleteUsuario
	*
    *
	*
	*
    *
*/
func DeleteUsuario(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var usr Models.Usuario

	// query
	err := Repositories.DeleteUsuario(&usr, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usr)
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, usr)
}
