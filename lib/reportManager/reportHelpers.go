package reportManager

import "github.com/AnnWann/pstu_finance_system/lib/models"

func getMemberPayments(registers []models.Register, members []models.Person) ([]models.Register, map[string]models.Person, float32, error) {
	membersAfterPaying := members
	memberPayments := []models.Register{}
	totalPayments := float32(0)

	for _, m := range membersAfterPaying { // para cada membro
		monthPayment := float32(0)
		for _, register := range registers { // para cada registro
			if register.Giver == m.Id && register.Type == "payment" { // quando o membro é o doador e o registro é de pagamento
				monthPayment = monthPayment + register.Value*float32(register.Amount)
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
	return memberPayments, membersAfterPayingMap, totalPayments, nil
}

func getJornalSales(registers []models.Register) ([]models.Register, float32, float32) {
	jornalSales := []models.Register{}
	totalJornalSales := float32(0)
	totalPartyShare := float32(0)

	for _, register := range registers {
		if register.Type == "jornal" {
			jornalSales = append(jornalSales, register)
			totalJornalSales = totalJornalSales + register.Value*float32(register.Amount)
			totalPartyShare = totalPartyShare + register.PartyShare*float32(register.Amount)
		}
	}

	return jornalSales, totalJornalSales, totalPartyShare
}

func getOtherSales(registers []models.Register) ([]models.Register, float32, float32) {
	otherSales := []models.Register{}
	totalOtherSales := float32(0)
	totalPartyShare := float32(0)

	for _, register := range registers {
		if register.Type == "other" {
			otherSales = append(otherSales, register)
			totalOtherSales = totalOtherSales + register.Value*float32(register.Amount)
			totalPartyShare = totalPartyShare + register.PartyShare*float32(register.Amount)
		}
	}

	return otherSales, totalOtherSales, totalPartyShare
}

func getExpenses(registers []models.Register) ([]models.Register, float32) {
	expenses := []models.Register{}
	totalExpenses := float32(0)

	for _, register := range registers {
		if register.Type == "expense" {
			expenses = append(expenses, register)
			totalExpenses = totalExpenses + register.Value*float32(register.Amount)
		}
	}

	return expenses, totalExpenses
}
