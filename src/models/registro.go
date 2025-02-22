package models

type Registro struct {
	ID         int     `json:"id"`
	Dia        string  `json:"dia"`
	Mes        string  `json:"mes"`
	Ano        string  `json:"ano"`
	Nucleo     int     `json:"nucleo"`
	Tipo       int     `json:"tipo"`
	Pagante    int     `json:"pagante"`
	Cobrante   int     `json:"cobrante"`
	Quantidade int     `json:"quantidade"`
	Valor      float64 `json:"valor"`
	Descricao  string  `json:"descricao"`
}
