package models

type Register struct {
	Id          int     `json:"id"`
	Day         string  `json:"day"`
	Month       string  `json:"month"`
	Year        string  `json:"year"`
	Nucleo      int     `json:"nucleo"`
	Type        int     `json:"type"`
	Giver       string  `json:"giver"`
	Receiver    string  `json:"receiver"`
	Amount      int     `json:"amount"`
	Value       float64 `json:"value"`
	Description string  `json:"description"`
}
