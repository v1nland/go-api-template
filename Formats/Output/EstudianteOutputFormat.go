package Output

import (
	"go-api-template/Messages/Response"
	"go-api-template/Models"
)

func GetEstudiantesOutput(u []Models.Estudiante) (output []Response.ListEstudiantesResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListEstudiantesResponse{
			Rol_estudiante:                u[i].Rol_estudiante,
			Rut_estudiante:                u[i].Rut_estudiante,
			Nombres_estudiante:            u[i].Nombres_estudiante,
			Apellidos_estudiante:          u[i].Apellidos_estudiante,
			Correo_electronico_estudiante: u[i].Correo_electronico_estudiante,
			Telefono_fijo_estudiante:      u[i].Telefono_fijo_estudiante,
			Telefono_celular_estudiante:   u[i].Telefono_celular_estudiante,
		})
	}

	return output
}

func GetOneEstudianteOutput(u Models.Estudiante) (output Response.GetOneEstudianteResponse) {
	return Response.GetOneEstudianteResponse{
		Rol_estudiante:                u.Rol_estudiante,
		Rut_estudiante:                u.Rut_estudiante,
		Nombres_estudiante:            u.Nombres_estudiante,
		Apellidos_estudiante:          u.Apellidos_estudiante,
		Correo_electronico_estudiante: u.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      u.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   u.Telefono_celular_estudiante,
	}
}

func AddNewEstudianteOutput(u Models.Estudiante) (output Response.AddNewEstudianteResponse) {
	return Response.AddNewEstudianteResponse{
		Rol_estudiante:                u.Rol_estudiante,
		Rut_estudiante:                u.Rut_estudiante,
		Nombres_estudiante:            u.Nombres_estudiante,
		Apellidos_estudiante:          u.Apellidos_estudiante,
		Correo_electronico_estudiante: u.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      u.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   u.Telefono_celular_estudiante,
	}
}

func PutOneEstudianteOutput(u Models.Estudiante) (output Response.PutOneEstudianteResponse) {
	return Response.PutOneEstudianteResponse{
		Rol_estudiante:                u.Rol_estudiante,
		Rut_estudiante:                u.Rut_estudiante,
		Nombres_estudiante:            u.Nombres_estudiante,
		Apellidos_estudiante:          u.Apellidos_estudiante,
		Hash_contrasena_estudiante:    u.Hash_contrasena_estudiante,
		Correo_electronico_estudiante: u.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      u.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   u.Telefono_celular_estudiante,
	}
}

func DeleteEstudianteOutput(u Models.Estudiante) (output Response.DeleteEstudianteResponse) {
	return Response.DeleteEstudianteResponse{
		Rol_estudiante:                u.Rol_estudiante,
		Rut_estudiante:                u.Rut_estudiante,
		Nombres_estudiante:            u.Nombres_estudiante,
		Apellidos_estudiante:          u.Apellidos_estudiante,
		Correo_electronico_estudiante: u.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      u.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   u.Telefono_celular_estudiante,
	}
}
