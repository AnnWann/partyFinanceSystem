package models

type Register struct {
	Id          string
	Day         string
	Month       string
	Year        string
	Type        string
	Giver       string
	Receiver    string
	Amount      int16
	Value       float32
	PartyShare  float32
	Description string
}
