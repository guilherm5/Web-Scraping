package models

type EmpregareObject struct {
	Model struct {
		Dados []struct {
			Titulo              string `json:"titulo"`
			Salario             string `json:"salario"`
			TrabalhoRemoto      string `json:"trabalhoRemoto"`
			TrabalhoRemotoTexto string `json:"trabalhoRemotoTexto"`
		} `json:"dados"`
	} `json:"model"`
}
