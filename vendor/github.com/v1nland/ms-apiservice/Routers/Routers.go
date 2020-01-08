package Routers

import (
	"github.com/v1nland/ms-apiservice/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// segment API by version
	v1 := r.Group("/v1")
	{
		v1.GET("book", Controllers.ListBook)
		v1.POST("book", Controllers.AddNewBook)
		v1.GET("book/:id", Controllers.GetOneBook)
		v1.PUT("book/:id", Controllers.PutOneBook)
		v1.DELETE("book/:id", Controllers.DeleteBook)

		v1.GET("usuario", Controllers.ListUsuario)
		v1.POST("usuario", Controllers.AddNewUsuario)
		v1.GET("usuario/:id", Controllers.GetOneUsuario)
		v1.PUT("usuario/:id", Controllers.PutOneUsuario)
		v1.DELETE("usuario/:id", Controllers.DeleteUsuario)
	}

	return r
}
