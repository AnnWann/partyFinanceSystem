package models

type MonthReport struct {
	Id                 string            `json:"id"`
	Month              string            `json:"month"`
	Year               string            `json:"year"`
	Members            map[string]Person `json:"members"`
	MembersAfterPaying map[string]Person `json:"members_after_paying"`
	Nucleo             Nucleo            `json:"nucleo"`
	Partido            Partido           `json:"partido"`
	MembersPayments    SubReport         `json:"members_payments"`
	Expenses           SubReport         `json:"expenses"`
	Sales              Sales             `json:"sales"`
	TotalEarned        float64           `json:"total_earned"`
	TotalLiquid        float64           `json:"total_liquid"`
	PartyDebts         float64           `json:"party_debts"`
	CoreSurplus        float64           `json:"core_surplus"`
}

type CompactMonthReport struct {
	Id              string   `json:"id"`
	Month           string   `json:"month"`
	Year            string   `json:"year"`
	Members         []Person `json:"members"`
	MembersPayments []string `json:"members_payments"`
	Expenses        []string `json:"expenses"`
	Sales           []string `json:"sales"`
	TotalEarned     float64  `json:"total_earned"`
	TotalLiquid     float64  `json:"total_liquid"`
	PartyDebts      float64  `json:"party_debts"`
	CoreSurplus     float64  `json:"core_surplus"`
}

func (r *CompactMonthReport) GetFullReport() (MonthReport, error) {
	//TODO: Implement this
	return MonthReport{}, nil
}

type SubReport struct {
	Registers []Register
	Type      string
	Total     float64
}

type Sales struct {
	EachType   map[TypeOfRegister]SubReport
	TotalSales float64
}
