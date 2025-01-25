package models

type Person struct {
	Id             string  `json:"id"`
	Name           string  `json:"name"`
	Role           string  `json:"role"`
	Nucleo         int  `json:"nucleo"`
	MonthlyPayment float64 `json:"monthly_payment"`
	Credit         float64 `json:"credit"`
}
