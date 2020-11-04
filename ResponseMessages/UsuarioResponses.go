package ResponseMessages

type ListUsuarioResponse struct {
}

type AddNewUsuarioResponse struct {
	Usuario  string `json:"usuario"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

type GetOneUsuarioResponse struct {
}

type PutOneUsuarioResponse struct {
	Usuario  string `json:"usuario"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

type DeleteUsuarioResponse struct {
}
