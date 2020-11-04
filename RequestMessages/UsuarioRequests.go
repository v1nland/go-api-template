package RequestMessages

type ListUsuarioPayload struct {
}

type AddNewUsuarioPayload struct {
	Usuario  string `json:"usuario"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

type GetOneUsuarioPayload struct {
}

type PutOneUsuarioPayload struct {
	Usuario  string `json:"usuario"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

type DeleteUsuarioPayload struct {
}
