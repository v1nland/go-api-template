package ApiHelpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api-template/Utils"
)

type ResponseData struct {
	Status int
	Meta   interface{}
	Data   interface{}
}

func RespondJSON(w *gin.Context, status int, payload interface{}) {
	// http status code
	fmt.Println("status ", status)

	// response container
	var res ResponseData

	// set values
	res.Status = status
	res.Meta = Utils.StatusMessage(status)
	res.Data = payload

	// response
	w.JSON(200, res)
}
