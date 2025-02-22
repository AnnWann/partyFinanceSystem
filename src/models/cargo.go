package models

type Cargo struct {
	ID        int    `json:"id"`
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	Nucleo    int    `json:"nucleo"`
}
