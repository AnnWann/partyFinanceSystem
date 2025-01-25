package reportManager

import "github.com/AnnWann/pstu_finance_system/src/models"

func getMemberPayments(registers []models.Register, members []models.Person, typeId int) (models.SubReport, map[string]models.Person, error) {
	membersAfterPaying := members
	memberPayments := []models.Register{}
	totalPayments := float64(0)

	for _, m := range membersAfterPaying { // para cada membro
		monthPayment := float64(0)
		for _, register := range registers { // para cada registro
			if register.Giver == m.Id && register.Type == typeId { // quando o membro é o doador e o registro é de pagamento
				monthPayment = monthPayment + register.Value*float64(register.Amount)
				memberPayments = append(memberPayments, register)
			}
		}

		m.Credit = m.Credit + monthPayment
		if m.Credit+monthPayment > m.MonthlyPayment {
			m.Credit = m.Credit + monthPayment - m.MonthlyPayment
			totalPayments = totalPayments + m.MonthlyPayment
		} else {
			totalPayments = totalPayments + m.Credit + monthPayment
			m.Credit = 0
		}
	}

	membersAfterPayingMap := make(map[string]models.Person)
	for _, m := range membersAfterPaying {
		membersAfterPayingMap[m.Id] = m
	}

	return models.SubReport{Registers: memberPayments, Type: "pagamento", Total: totalPayments}, membersAfterPayingMap, nil
}

func getSales(registers []models.Register, saleTypeId int) ([]models.Register, float64) {
	sales := []models.Register{}
	totalSales := float64(0)

	for _, register := range registers {
		if register.Type == saleTypeId {
			sales = append(sales, register)
			totalSales = totalSales + register.Value*float64(register.Amount)
		}
	}

	return sales, totalSales
}

func getExpenses(registers []models.Register, expensesId int) models.SubReport {
	expenses := []models.Register{}
	totalExpenses := float64(0)

	for _, register := range registers {
		if register.Type == expensesId {
			expenses = append(expenses, register)
			totalExpenses = totalExpenses + register.Value*float64(register.Amount)
		}
	}

	return models.SubReport{
		Registers: expenses,
		Type:      "despesa",
		Total:     totalExpenses,
	}
}

func calcPartyShare(sales map[models.TypeOfRegister]models.SubReport) float64 {
	PartyShare := float64(0)
	for t, s := range sales {
		s_ps := float64(0)
		for _, r := range s.Registers {
			s_ps = s_ps + t.PartyShare*float64(r.Amount)
		}
	}
	return PartyShare
}
