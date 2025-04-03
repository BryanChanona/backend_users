package domain

type SupervisorsModel struct {
	Id_supervisor int `json:"id_supervisor,omitempty"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Id_usuario int `json:"id_usuario"`
}

type SupervisorsResponse struct{
	Id_supervisor int `json:"id_supervisor"`
	Name string `json:"name"`
	Email string `json:"email"`
}