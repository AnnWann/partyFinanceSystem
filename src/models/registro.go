package models

type Registro struct {
	ID          int     `json:"id"`
	Dia         string  `json:"dia"`
	Mes         string  `json:"mes"`
	Ano         string  `json:"ano"`
	Nucleo      int     `json:"nucleo"`
	Tipo        int     `json:"tipo"`
	Pago_por    int     `json:"pago_por"`
	Cobrado_por int     `json:"cobrado_por"`
	Quantidade  int     `json:"quantidade"`
	Valor       float64 `json:"valor"`
	Descricao   string  `json:"descricao"`
}
