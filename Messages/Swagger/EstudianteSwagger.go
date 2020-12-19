package Swagger

import "go-api-template/Messages/Response"

type ListEstudiantesSwagger struct {
	Status bool                             `json:"Status"`
	Meta   string                           `json:"Meta"`
	Data   Response.ListEstudiantesResponse `json:"Data"`
}

type GetOneEstudianteSwagger struct {
	Status bool                              `json:"Status"`
	Meta   string                            `json:"Meta"`
	Data   Response.GetOneEstudianteResponse `json:"Data"`
}

type AddNewEstudianteSwagger struct {
	Status bool                              `json:"Status"`
	Meta   string                            `json:"Meta"`
	Data   Response.AddNewEstudianteResponse `json:"Data"`
}

type PutOneEstudianteSwagger struct {
	Status bool                              `json:"Status"`
	Meta   string                            `json:"Meta"`
	Data   Response.PutOneEstudianteResponse `json:"Data"`
}

type DeleteEstudianteSwagger struct {
	Status bool                              `json:"Status"`
	Meta   string                            `json:"Meta"`
	Data   Response.DeleteEstudianteResponse `json:"Data"`
}
