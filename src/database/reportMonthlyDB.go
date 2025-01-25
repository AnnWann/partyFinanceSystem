package database

import (
	"database/sql"

	"github.com/AnnWann/pstu_finance_system/src/models"
)

type ReportDB struct {
	*sql.DB
}

func (db *DBWrapper) GetReportDB() *ReportDB {
	return &ReportDB{db.DB}
}

func (db *ReportDB) GetNextId() (string, error) {
	var id string
	err := db.QueryRow("SELECT COUNT(*) FROM reports").Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (rDb ReportDB) InsertReport(r models.MonthReport) error {

	DB := GetDB()
	PersonDB := DB.GetPersonDB()
	for _, m := range r.MembersAfterPaying {
		err := PersonDB.UpdateCredit(m.Id, m.Credit)
		if err != nil {
			return err
		}
	}

	err := DB.GetNucleoDB().UpdateCredit(r.Nucleo.Id, r.Nucleo.Credit)
	if err != nil {
		return err
	}

	membersIds := make([]string, len(r.Members))
	for _, value := range r.Members {
		membersIds = append(membersIds, value.Id)
	}

	paymentsIds := make([]int, len(r.MembersPayments.Registers))
	for i, payment := range r.MembersPayments.Registers {
		paymentsIds[i] = payment.Id
	}

	expensesIds := make([]int, len(r.Expenses.Registers))
	for i, expense := range r.Expenses.Registers {
		expensesIds[i] = expense.Id
	}

	var allSalesId []int
	for _, salesType := range r.Sales.EachType {
		for _, r := range salesType.Registers {
				allSalesId = append(allSalesId, r.Id)
		}
	}

	_, err = rDb.Exec(
		"INSERT INTO"+
			"relatorio (id, mes, ano, nucleo, membros, pagamentos_membros, despesas, vendas, total_ganho, total_liquido, dividendos_partido, superavit_nucleo)"+
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		r.Id, r.Month, r.Year, r.Nucleo.Id, membersIds,
		paymentsIds, expensesIds, allSalesId, r.TotalEarned,
		r.TotalLiquid, r.PartyDebts, r.CoreSurplus)
	if err != nil {
		return err
	}

	return nil
}

func (db ReportDB) ReportExists(nucleo int, month string, year string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM relatorio WHERE nucleo = ? AND mes = ? AND ano = ?", nucleo, month, year).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (db ReportDB) GetReport(id string) (models.CompactMonthReport, error) {
	var r models.CompactMonthReport
	err := db.QueryRow("SELECT * FROM relatorio WHERE id = ?", id).
		Scan(&r.Id, &r.Month, &r.Year, &r.Members, &r.MembersPayments,
			&r.Expenses, &r.Sales, &r.TotalEarned,
			&r.TotalLiquid, &r.PartyDebts, &r.CoreSurplus)
	if err != nil {
		return r, err
	}

	return r, nil
}
