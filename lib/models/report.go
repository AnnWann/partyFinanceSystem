package models

type MonthReport struct {
	Id                 string
	Month              string
	Year               string
	Members            map[string]Person
	MembersAfterPaying map[string]Person
	Core               Person
	Party              Person
	MembersPayments    SubReport
	Expenses           SubReport
	Sales              Sales
	TotalEarned        float32
	TotalLiquid        float32
	PartyDebts         float32
	CoreSurplus        float32
}

type CompactMonthReport struct {
	Id              string
	Month           string
	Year            string
	Members         []Person
	MembersPayments []string
	Expenses        []string
	Sales           []string
	TotalEarned     float32
	TotalLiquid     float32
	PartyDebts      float32
	CoreSurplus     float32
}

type SubReport struct {
	Registers  []Register
	PartyShare float32
	Total      float32
}

type Sales struct {
	Jornals    SubReport
	Others     SubReport
	TotalSales float32
}
