package executors

import (
	"errors"
	"os"
	"strconv"

	"github.com/AnnWann/pstu_finance_system/src/database"
	pdfMaker "github.com/AnnWann/pstu_finance_system/src/pdf"
	"github.com/AnnWann/pstu_finance_system/src/reportManager"
	"github.com/joho/godotenv"
)

func AddMonthlyReport(nucleo string, month string, year string) (string, string, error) {
	nucleoInt, err := strconv.Atoi(nucleo)
	if err != nil {
		return "", "", errors.New("invalid nucleo id")
	}

	nucleoExists := database.GetDB().GetNucleoDB().NucleoExists(nucleoInt)
	if !nucleoExists {
		return "", "", errors.New("nucleo not found")
	}

	db := database.GetDB().GetReportDB()
	reportExists := db.ReportExists(nucleoInt, month, year)

	if reportExists {
		return "", "", errors.New("report already exists")
	}

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
	}

	report, partyPayment, err := reportManager.BuildMonthlyReport(month, year, nucleoInt)
	if err != nil {
		tx.Rollback()
		return "", "", err
	}

	err = database.GetDB().GetRegisterDB().InsertRegister(partyPayment)
	if err != nil {
		tx.Rollback()
		return "", "", err
	}

	err = db.InsertReport(report)
	if err != nil {
		tx.Rollback()
		return "", "", err
	}

	err = godotenv.Load()
	if err != nil {
		tx.Rollback()
		return "", "", err
	}

	pdf_folder := os.Getenv("PDF_FOLDER")
	path_to_pdf := pdf_folder + "/" + nucleo + "_" + year + "_" + month + "_" + report.Id + ".pdf"
	err = pdfMaker.PrintPDFMonthlyReport(report, path_to_pdf)
	if err != nil {
		tx.Rollback()
		return "", "", err
	}

	tx.Commit()
	return report.Id, path_to_pdf, nil
}

func GetMonthlyReport(filterOptions map[string]string) error {
	return nil
}

func GetPayday(nucleo string) (string, error) {
	nucleoInt, err := strconv.Atoi(nucleo)
	if err != nil {
		return "", errors.New("invalid nucleo id")
	}
	payday, err := database.GetDB().GetPaydayDB().GetPayday(nucleoInt)
	return payday, err
}
func SetPayday(payday string) error {
	if payday == "" {
		return errors.New("invalid arguments. The correct format is 'set payday <day>'")
	}
	err := database.GetDB().GetPaydayDB().SetPayday(payday)
	return err
}
