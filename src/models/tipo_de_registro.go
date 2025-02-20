package models

type Tipo_de_registro struct {
	ID                 int     `json:"id"`
	Nome               string  `json:"nome"`
	Nucleo             int  `json:"nucleo"`
	Descricao          string  `json:"descricao"`
	Parcela_partidaria float64 `json:"parcela_partidaria"`
}
