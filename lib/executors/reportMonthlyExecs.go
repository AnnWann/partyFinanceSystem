package executors

import (
	"errors"

	"github.com/AnnWann/pstu_finance_system/lib/database"
	pdfMaker "github.com/AnnWann/pstu_finance_system/lib/pdf"
	"github.com/AnnWann/pstu_finance_system/lib/reportManager"
)

func GetMonthlyReport(month string, year string) error {
	if month == "" || year == "" {
		return errors.New("invalid arguments. The correct format is 'get report <month> <year>'")
	}

	db := database.GetDB().GetReportDB()
	reportSummary, err := db.GetReport(month + "-" + year)
	if err != nil {
		return err
	}
	if reportSummary.Id != "" {
		report, err := reportSummary.GetFullReport()
		if err != nil {
			return err
		}
		err = pdfMaker.MakePDFReport(report)
		if err != nil {
			return err
		}
	} else {
		tx, err := db.Begin()
		if err != nil {
			tx.Rollback()
			return err
		}
		report, err := reportManager.GetMonthlyReport(month, year)
		if err != nil {
			tx.Rollback()
			return err
		}
		err = pdfMaker.MakePDFReport(report)
		if err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
	}

	return nil
}

func GetPayday() (string, error) {
	payday, err := database.GetDB().GetPaydayDB().GetPayday()
	return payday, err
}
func SetPayday(payday string) error {
	if payday == "" {
		return errors.New("invalid arguments. The correct format is 'set payday <day>'")
	}
	err := database.GetDB().GetPaydayDB().SetPayday(payday)
	return err
}
