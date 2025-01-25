package reportManager

import (
	"github.com/AnnWann/pstu_finance_system/src/database"
	"github.com/AnnWann/pstu_finance_system/src/models"
)

func BuildMonthlyReport(month string, year string, nucleoId int) (models.MonthReport, models.Register, error) {

	db := database.GetDB()

	registersOfTheMonth, err := db.GetRegisterDB().GetRegisterByMonthAndYear(month, year)
	if err != nil {
		return models.MonthReport{}, models.Register{}, err
	}

	personDB := db.GetPersonDB()
	members, err := personDB.GetPersonByNucleo(nucleoId)
	if err != nil {
		return models.MonthReport{}, models.Register{}, err
	}

	nucleo, err := db.GetNucleoDB().GetNucleoById(nucleoId)
	if err != nil {
		return models.MonthReport{}, models.Register{}, err
	}

	partido, err := db.GetPartidoDB().GetPartido()
	if err != nil {
		return models.MonthReport{}, models.Register{}, err
	}

	typesOfRegisters, err := db.GetTypeOfRegisterDB().GetTypesByNucleo(nucleoId)
	if err != nil {
		return models.MonthReport{}, models.Register{}, err
	}

	membersMap := make(map[string]models.Person)
	for _, m := range members {
		membersMap[m.Id] = m
	}

	var paymentsId int
	for _, t := range typesOfRegisters {
		if t.Name == "pagamento" {
			paymentsId = t.Id
			break
		}
	}

	membersReport, membersAfterPaying, err := getMemberPayments(registersOfTheMonth, members, paymentsId)
	if err != nil {
		return models.MonthReport{}, models.Register{}, err
	}

	Sales := models.Sales{}
	var expensesId int
	for _, t := range typesOfRegisters {
		if t.Name == "pagamento" {
			continue
		}
		if t.Name == "despesa" {
			expensesId = t.Id
			continue
		}

		registers, totalSales := getSales(registersOfTheMonth, t.Id)

		Sales.EachType[t] = models.SubReport{
			Registers: registers,
			Type:      t.Name,
			Total:     totalSales,
		}

		Sales.TotalSales += totalSales
	}

	expensesReport := getExpenses(registersOfTheMonth, expensesId)

	reportId := nucleo.Name + "-" + month + "/" + year

	totalEarned := Sales.TotalSales + membersReport.Total
	totalLiquid := totalEarned - expensesReport.Total

	SalesPartyShare := calcPartyShare(Sales.EachType)
	PartyDebts := SalesPartyShare + membersReport.Total

	CoreSurplus := totalLiquid - PartyDebts

	report := models.MonthReport{
		Id:                 reportId,
		Month:              month,
		Year:               year,
		Members:            membersMap,
		MembersAfterPaying: membersAfterPaying,
		Nucleo:             nucleo,
		Partido:            partido,
		MembersPayments:    membersReport,
		Expenses:           expensesReport,
		Sales:              Sales,
		TotalEarned:        totalEarned,
		TotalLiquid:        totalLiquid,
		PartyDebts:         PartyDebts,
		CoreSurplus:        CoreSurplus,
	}

	payday, err := db.GetPaydayDB().GetPayday(nucleoId)
	if err != nil {
		return models.MonthReport{}, models.Register{}, err
	}

	partyDebtsRegisterId, err := db.GetRegisterDB().GetNextId()
	if err != nil {
		return models.MonthReport{}, models.Register{}, err
	}

	partyDebtsRegister := models.Register{
		Id:          partyDebtsRegisterId,
		Day:         payday,
		Month:       month,
		Year:        year,
		Nucleo:      nucleoId,
		Giver:       "",
		Receiver:    "",
		Type:        paymentsId,
		Description: "Pagamento ao partido",
		Value:       PartyDebts,
		Amount:      1,
	}

	return report, partyDebtsRegister, err
}

/* func GetYearlyReport(year string) (models.YearlyReport, error) {
	//TODO: Implement this
} */
