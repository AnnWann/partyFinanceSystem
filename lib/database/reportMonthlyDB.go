package database

import (
	"database/sql"

	"github.com/AnnWann/pstu_finance_system/lib/models"
)

type ReportDB struct {
	*sql.DB
}

func (db *DBWrapper) GetReportDB() *ReportDB {
	return &ReportDB{db.DB}
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

	err := PersonDB.UpdateCredit(r.Core.Id, r.Core.Credit)
	if err != nil {
		return err
	}

	membersIds := make([]string, len(r.Members))
	for _, value := range r.Members {
		membersIds = append(membersIds, value.Id)
	}

	paymentsIds := make([]string, len(r.MembersPayments.Registers))
	for i, payment := range r.MembersPayments.Registers {
		paymentsIds[i] = payment.Id
	}

	expensesIds := make([]string, len(r.Expenses.Registers))
	for i, expense := range r.Expenses.Registers {
		expensesIds[i] = expense.Id
	}

	allSales := append(r.Sales.Jornals.Registers, r.Sales.Others.Registers...)
	salesIds := make([]string, len(allSales))
	for i, sale := range allSales {
		salesIds[i] = sale.Id
	}

	_, err = rDb.Exec(
		"INSERT INTO"+
			"reports (id, month, year, members, membersPayments, expenses, sales, totalEarned, totalLiquid, partyDebts, CoreSurplus)"+
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		r.Id, r.Month, r.Year, membersIds,
		paymentsIds, expensesIds, salesIds, r.TotalEarned,
		r.TotalLiquid, r.PartyDebts, r.CoreSurplus)
	if err != nil {
		return err
	}

	return nil
}

func (db ReportDB)GetReport(id string) (models.CompactMonthReport, error) {
	var r models.CompactMonthReport
	err := db.QueryRow("SELECT * FROM reports WHERE id = ?", id).
		Scan(&r.Id, &r.Month, &r.Year, &r.Members, &r.MembersPayments,
			&r.Expenses, &r.Sales, &r.TotalEarned,
			&r.TotalLiquid, &r.PartyDebts, &r.CoreSurplus)
	if err != nil {
		return r, err
	}

	return r, nil
}
