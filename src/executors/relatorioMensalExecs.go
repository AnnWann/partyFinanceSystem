package executors

import (
	"errors"
	"os"
	"strconv"

	"github.com/AnnWann/pstu_finance_system/src/database"
	"github.com/AnnWann/pstu_finance_system/src/models"
	pdfMaker "github.com/AnnWann/pstu_finance_system/src/pdf"
	"github.com/AnnWann/pstu_finance_system/src/reportManager"
	"github.com/joho/godotenv"
)

func AddRelatorioMensal(nucleo string, mes string, ano string) (string, string, error) {
	nucleoInt, err := strconv.Atoi(nucleo)
	if err != nil {
		return "", "", errors.New("invalid nucleo id")
	}

	nucleoExists := database.GetDB().GetNucleoDB().NucleoExists(nucleoInt)
	if !nucleoExists {
		return "", "", errors.New("nucleo not found")
	}

	db := database.GetDB().GetRelatorioMensalDB()
	reportExists := db.RelatorioExists(nucleoInt, mes, ano)

	if reportExists {
		return "", "", errors.New("report already exists")
	}

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
	}

	report, partyPayment, err := reportManager.BuildRelatorioMensal(mes, ano, nucleoInt)
	if err != nil {
		tx.Rollback()
		return "", "", err
	}

	err = database.GetDB().GetRegisterDB().InsertRegister(partyPayment)
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
	path_to_pdf := pdf_folder + "/" + nucleo + "_" + ano + "_" + mes + "_" + report.ID + ".pdf"

	err = db.InsertRelatorio(report)
	if err != nil {
		tx.Rollback()
		return "", "", err
	}

	err = godotenv.Load()
	if err != nil {
		tx.Rollback()
		return "", "", err
	}

	err = pdfMaker.PrintPDFMonthlyReport(report, path_to_pdf)
	if err != nil {
		tx.Rollback()
		return "", "", err
	}

	tx.Commit()
	return report.ID, path_to_pdf, nil
}

type RM_get_response struct {
	Mes string
	Ano string
	Url string
}

func GetRelatorioMensal(filterOptions map[string]string) ([]RM_get_response, error) {
	db := database.GetDB().GetRelatorioMensalDB()
	rm, err := db.GetRelatorioMensal()
	if err != nil {
		return nil, err
	}

	relatorios := filterRelatoriosMensal(rm, filterOptions)
	var relatoriosPaths []RM_get_response
	for _, r := range relatorios {
		response := RM_get_response{r.Mes, r.Ano, r.Link_Arquivo}
		relatoriosPaths = append(relatoriosPaths, response)
	}

	return relatoriosPaths, nil
}

func filterRelatoriosMensal(rm []models.Relatorio_mensal, filterOptions map[string]string) []models.Relatorio_mensal {
	if filterOptions == nil {
		return rm
	}

	var filteredRM []models.Relatorio_mensal
	for _, r := range rm {
		if filterRelatorioMensal(r, filterOptions) {
			filteredRM = append(filteredRM, r)
		}
	}

	return filteredRM
}

func filterRelatorioMensal(r models.Relatorio_mensal, filterOptions map[string]string) bool {
	isValid := false
	for key, value := range filterOptions {
		switch key {
		case "--nucleo":
			nucleo, err := strconv.Atoi(value)
			if err != nil {
				isValid = false
			} else {
				isValid = r.Nucleo == nucleo
			}
		case "--mes":
			isValid = r.Mes == value
		case "--ano":
			isValid = r.Ano == value
		}
	}
	return isValid
}

func GetPayday(nucleo string) (string, error) {
	nucleoInt, err := strconv.Atoi(nucleo)
	if err != nil {
		return "", errors.New("invalid nucleo id")
	}
	nucleoOBJ, err := database.GetDB().GetNucleoDB().GetNucleoById(nucleoInt)
	return nucleoOBJ.Dia_de_Pagamento, err
}
func SetPayday(nucleo string, payday string) error {
	nucleoInt, err := strconv.Atoi(nucleo)
	if err != nil {
		return errors.New("invalid nucleo id")
	}

	err = database.GetDB().GetNucleoDB().UpdateDiaDePagamento(nucleoInt, payday)
	return err
}
