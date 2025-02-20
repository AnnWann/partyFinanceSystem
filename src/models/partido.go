package models

type Partido struct {
	ID      int     `json:"id"`
	Nome    string  `json:"nome"`
	Reserva float64 `json:"reserva"`
}
