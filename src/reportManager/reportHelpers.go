package reportManager

import "github.com/AnnWann/pstu_finance_system/src/models"

type RegistrosPorTipo map[int]models.SubRelatorio

func getRegistrosPorTipo(registers []models.Registro, types []models.Tipo_de_registro) RegistrosPorTipo {
	registersByType := make(RegistrosPorTipo)

	typesMap := make(map[int]models.Tipo_de_registro)
	for _, t := range types {
		typesMap[t.ID] = t
	}

	for _, r := range registers {
		if _, ok := registersByType[r.Tipo]; !ok {
			registersByType[r.Tipo] = models.SubRelatorio{Registros: []models.Registro{}, Tipo: typesMap[r.Tipo].Nome, Total: 0}
		}
		registersByType[r.Tipo] = models.SubRelatorio{
			Registros: append(registersByType[r.Tipo].Registros, r),
			Tipo:      typesMap[r.Tipo].Nome,
			Total:     registersByType[r.Tipo].Total + r.Valor*float64(r.Quantidade),
		}
	}
	return registersByType
}

func applyMemberPayments(membersReport *models.SubRelatorio, members []models.Membro) map[int]models.Membro {
	membersAfterPaying := members
	totalPayments := float64(0)

	for i, m := range membersAfterPaying { // para cada membro
		monthPayment := float64(0)
		for _, register := range membersReport.Registros { // para cada registro de pagamento
			if register.Pagante != m.ID {
				continue
			}
			monthPayment = monthPayment + register.Valor*float64(register.Quantidade)
		}

		new_credit := m.Credito + monthPayment - m.Contribuicao_mensal
		if new_credit >= 0 {
			totalPayments = totalPayments + m.Contribuicao_mensal
			m.Credito = new_credit
		} else {
			totalPayments = totalPayments + m.Credito + monthPayment
			m.Credito = 0
		}

		membersAfterPaying[i] = m
	}

	membersAfterPayingMap := make(map[int]models.Membro)
	for _, m := range membersAfterPaying {
		membersAfterPayingMap[m.ID] = m
	}

	membersReport.Total = totalPayments
	return membersAfterPayingMap
}

func extractRegistrosEspecificosDeNucleo(r RegistrosPorTipo) models.Registros_Especificos_Nucleo {
	especificos := models.Registros_Especificos_Nucleo{}
	especificos.Tipos = make(map[int]models.SubRelatorio)
	for t, s := range r {
		if t < 0 { //pula os tipos de registro gerais
			continue
		}
		especificos.Tipos[t] = s
		especificos.Total += s.Total
	}
	return especificos
}

func calcPartilhaPartidariaEspecifica(tipos []models.Tipo_de_registro, especificos map[int]models.SubRelatorio) float64 {
	partilhaPartidaria := float64(0)

	for _, t := range tipos {
		if t.ID < 0 {
			continue
		}
		for t_e, e := range especificos {
			if t_e != t.ID {
				continue
			}
			e_pp := t.Parcela_partidaria * float64(len(e.Registros))
			partilhaPartidaria = partilhaPartidaria + e_pp
		}
	}
	return partilhaPartidaria
}
