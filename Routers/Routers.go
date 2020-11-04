package Routers

import (
	"github.com/gin-gonic/gin"
	"go-api-template/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// segment API by version
	r.Group("/v1")
	{
		usuario := r.Group("/usuario")
		{
			usuario.GET("", Controllers.ListUsuario)
			usuario.POST("", Controllers.AddNewUsuario)
			usuario.GET(":id", Controllers.GetOneUsuario)
			usuario.PUT(":id", Controllers.PutOneUsuario)
			usuario.DELETE(":id", Controllers.DeleteUsuario)
		}
	}

	return r
}
