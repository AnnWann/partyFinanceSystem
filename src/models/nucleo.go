package models

type Nucleo struct {
	ID               int     `json:"id"`
	Nome             string  `json:"nome"`
	Cidade           string  `json:"cidade"`
	Estado           string  `json:"estado"`
	Reserva          float64 `json:"reserva"`
	Dia_de_Pagamento string  `json:"dia_de_pagamento"`
}
