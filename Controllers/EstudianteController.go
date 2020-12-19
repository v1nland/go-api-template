package Controllers

import (
	"go-api-template/ApiHelpers"
	"go-api-template/Formats/Input"
	"go-api-template/Formats/Output"
	"go-api-template/Messages/Request"
	"go-api-template/Models"
	"go-api-template/Repositories"
	"go-api-template/Utils"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de estudiantes
// @Description Lista todos los estudiantes
// @Tags Estudiantes
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListEstudiantesSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes [get]
func ListEstudiantes(c *gin.Context) {
	// model container
	var container []Models.Estudiante

	// query
	err := Repositories.GetAllEstudiantes(&container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.GetEstudiantesOutput(container))
}

// @Summary Obtiene un estudiante
// @Description Obtiene un estudiante según su UUID
// @Tags Estudiantes
// @Accept  json
// @Produce  json
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a buscar"
// @Success 200 {object} Swagger.GetOneEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/{uuid_estudiante} [get]
func GetOneEstudiante(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Estudiante

	// query
	err := Repositories.GetOneEstudiante(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.GetOneEstudianteOutput(container))
}

// @Summary Agrega un nuevo estudiante
// @Description Genera un nuevo estudiante con los datos entregados
// @Tags Estudiantes
// @Accept  json
// @Produce  json
// @Param   input_estudiante     body    Request.AddNewEstudiantePayload     true        "Estudiante a agregar"
// @Success 200 {object} Swagger.AddNewEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes [post]
func AddNewEstudiante(c *gin.Context) {
	// input container
	var container Request.AddNewEstudiantePayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.AddNewEstudianteInput(&container)

	// generate model entity
	model_container := Models.Estudiante{
		Id_rol:                        container.Id_rol,
		Rut_estudiante:                container.Rut_estudiante,
		Nombres_estudiante:            container.Nombres_estudiante,
		Apellidos_estudiante:          container.Apellidos_estudiante,
		Hash_contrasena_estudiante:    container.Hash_contrasena_estudiante,
		Correo_electronico_estudiante: container.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      container.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   container.Telefono_celular_estudiante,
	}

	// query
	err := Repositories.AddNewEstudiante(&model_container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.AddNewEstudianteOutput(model_container))
}

// @Summary Modifica un estudiante
// @Description Modifica un estudiante con los datos entregados
// @Tags Estudiantes
// @Accept  json
// @Produce  json
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a modificar"
// @Param   input_actualiza_estudiante     body    Request.PutOneEstudiantePayload     true        "Estudiante a modificar"
// @Success 200 {object} Swagger.PutOneEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/{uuid_estudiante} [put]
func PutOneEstudiante(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container Request.PutOneEstudiantePayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.PutOneEstudianteInput(&container)

	// generate model entity
	var model_container Models.Estudiante

	// get query
	err := Repositories.GetOneEstudiante(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model_container = Models.Estudiante{
		ID:                            model_container.ID,
		Id_rol:                        Utils.CheckUpdatedInt(container.Id_rol, model_container.Id_rol),
		Rut_estudiante:                Utils.CheckUpdatedString(container.Rut_estudiante, model_container.Rut_estudiante),
		Nombres_estudiante:            Utils.CheckUpdatedString(container.Nombres_estudiante, model_container.Nombres_estudiante),
		Apellidos_estudiante:          Utils.CheckUpdatedString(container.Apellidos_estudiante, model_container.Apellidos_estudiante),
		Hash_contrasena_estudiante:    Utils.CheckUpdatedString(container.Hash_contrasena_estudiante, model_container.Hash_contrasena_estudiante),
		Correo_electronico_estudiante: Utils.CheckUpdatedString(container.Correo_electronico_estudiante, model_container.Correo_electronico_estudiante),
		Telefono_fijo_estudiante:      Utils.CheckUpdatedString(container.Telefono_fijo_estudiante, model_container.Telefono_fijo_estudiante),
		Telefono_celular_estudiante:   Utils.CheckUpdatedString(container.Telefono_celular_estudiante, model_container.Telefono_celular_estudiante),
	}

	// put query
	err = Repositories.PutOneEstudiante(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.PutOneEstudianteOutput(model_container))
}

// @Summary Elimina un estudiante
// @Description Elimina un estudiante con los datos entregados
// @Tags Estudiantes
// @Accept  json
// @Produce  json
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a eliminar"
// @Success 200 {object} Swagger.DeleteEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/{uuid_estudiante} [delete]
func DeleteEstudiante(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Estudiante

	// query
	err := Repositories.DeleteEstudiante(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.DeleteEstudianteOutput(container))
}
