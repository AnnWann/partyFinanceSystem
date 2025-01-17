package reportManager

import (
	"github.com/AnnWann/pstu_finance_system/lib/database"
	"github.com/AnnWann/pstu_finance_system/lib/models"
)

func GetMonthlyReport(month string, year string) (models.MonthReport, error) {

	registersOfTheMonth, err := database.GetRegisterByMonthAndYear(month, year)
	if err != nil {
		return models.MonthReport{}, err
	}

	members, err := database.GetMembers()
	if err != nil {
		return models.MonthReport{}, err
	}

	core, err := database.GetCore()
	if err != nil {
		return models.MonthReport{}, err
	}

	party, err := database.GetParty()
	if err != nil {
		return models.MonthReport{}, err
	}

	membersMap := make(map[string]models.Person)
	for _, m := range members {
		membersMap[m.Id] = m
	}
	memberPayments, membersAfterPaying, totalPayments, err := getMemberPayments(registersOfTheMonth, members)
	if err != nil {
		return models.MonthReport{}, err
	}

	jornalSales, totalJornalSales, journalPartyShare := getJornalSales(registersOfTheMonth)

	otherSales, totalOtherSales, othersPartyShare := getOtherSales(registersOfTheMonth)

	expenses, totalExpenses := getExpenses(registersOfTheMonth)

	reportId := month + "-" + year

	report := models.MonthReport{
		Id:                 reportId,
		Month:              month,
		Year:               year,
		Members:            membersMap,
		MembersAfterPaying: membersAfterPaying,
		Core:               core,
		Party:              party,
		MembersPayments: models.SubReport{
			Registers:  memberPayments,
			PartyShare: totalPayments,
			Total:      totalPayments,
		},
		Expenses: models.SubReport{
			Registers:  expenses,
			PartyShare: 0,
			Total:      totalExpenses,
		},
		Sales: models.Sales{
			Jornals: models.SubReport{
				Registers:  jornalSales,
				PartyShare: journalPartyShare,
				Total:      totalJornalSales,
			},
			Others: models.SubReport{
				Registers:  otherSales,
				PartyShare: othersPartyShare,
				Total:      totalOtherSales,
			},
			TotalSales: totalJornalSales + totalOtherSales,
		},
		TotalEarned: totalJornalSales + totalOtherSales + totalPayments,
		TotalLiquid: totalJornalSales + totalOtherSales + totalPayments - totalExpenses,
		PartyDebts:  journalPartyShare + othersPartyShare + totalPayments,
		CoreSurplus: (totalJornalSales + totalOtherSales + totalPayments + core.Credit - totalExpenses) -
			(journalPartyShare + othersPartyShare + totalPayments),
	}

	payday, err := database.GetPayday()
	if err != nil {
		return models.MonthReport{}, err
	}

	partyDebts := models.Register{
		Day:         payday,
		Month:       month,
		Year:        year,
		Type:        "partyDebts",
		Description: "Pagamento do partido",
		Value:       journalPartyShare + othersPartyShare + totalPayments,
		PartyShare:  journalPartyShare + othersPartyShare + totalPayments,
		Amount:      1,
	}

	err = database.InsertRegister(partyDebts)
	if err != nil {
		return models.MonthReport{}, err
	}

	err = database.InsertReport(report)

	return report, err
}

/* func GetYearlyReport(year string) (models.YearlyReport, error) {
	//TODO: Implement this
} */
