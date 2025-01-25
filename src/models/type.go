package models

type TypeOfRegister struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Nucleo      string  `json:"nucleo"`
	Description string  `json:"description"`
	PartyShare  float64 `json:"party_share"`
}
