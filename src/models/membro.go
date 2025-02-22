package models

type Membro struct {
	ID                  int     `json:"id"`
	Nome                string  `json:"nome"`
	Cargo               int     `json:"designacao"`
	Nucleo              int     `json:"nucleo"`
	Contribuicao_mensal float64 `json:"contribuicao_mensal"`
	Credito             float64 `json:"credito"`
}
