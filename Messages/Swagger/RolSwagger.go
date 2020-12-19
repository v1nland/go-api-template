package Swagger

import "go-api-template/Messages/Response"

type ListRolesSwagger struct {
	Status bool                       `json:"Status"`
	Meta   string                     `json:"Meta"`
	Data   Response.ListRolesResponse `json:"Data"`
}

type GetOneRolSwagger struct {
	Status bool                       `json:"Status"`
	Meta   string                     `json:"Meta"`
	Data   Response.GetOneRolResponse `json:"Data"`
}

type AddNewRolSwagger struct {
	Status bool                       `json:"Status"`
	Meta   string                     `json:"Meta"`
	Data   Response.AddNewRolResponse `json:"Data"`
}

type PutOneRolSwagger struct {
	Status bool                       `json:"Status"`
	Meta   string                     `json:"Meta"`
	Data   Response.PutOneRolResponse `json:"Data"`
}

type DeleteRolSwagger struct {
	Status bool                       `json:"Status"`
	Meta   string                     `json:"Meta"`
	Data   Response.DeleteRolResponse `json:"Data"`
}
