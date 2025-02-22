package reportManager

import (
	"strconv"

	"github.com/AnnWann/pstu_finance_system/src/database"
	"github.com/AnnWann/pstu_finance_system/src/models"
)

func BuildRelatorioMensal(month string, year string, nucleoId int) (models.Relatorio_mensal_complexo, models.Registro, error) {

	db := database.GetDB()

	registrosDoMes, err := db.GetRegisterDB().GetRegisterByMonthAndYear(month, year)
	if err != nil {
		return models.Relatorio_mensal_complexo{}, models.Registro{}, err
	}

	membroDB := db.GetMembroDB()
	membros, err := membroDB.GetMembroByNucleo(nucleoId)
	if err != nil {
		return models.Relatorio_mensal_complexo{}, models.Registro{}, err
	}

	nucleo, err := db.GetNucleoDB().GetNucleoById(nucleoId)
	if err != nil {
		return models.Relatorio_mensal_complexo{}, models.Registro{}, err
	}

	partido, err := db.GetPartidoDB().GetPartido()
	if err != nil {
		return models.Relatorio_mensal_complexo{}, models.Registro{}, err
	}

	tiposDeRegistros, err := db.GetTiposDeRegistroDB().GetTipoPorNucleo(nucleoId)
	if err != nil {
		return models.Relatorio_mensal_complexo{}, models.Registro{}, err
	}

	membrosMAP := make(map[int]models.Membro)
	for _, m := range membros {
		membrosMAP[m.ID] = m
	}

	registrosPorTipo := getRegistrosPorTipo(registrosDoMes, tiposDeRegistros)

	contribuicaoId := db.GetTiposDeRegistroDB().GetContribuicaoId()
	gastosId := db.GetTiposDeRegistroDB().GetGastosId()

	contribuicao := registrosPorTipo[contribuicaoId]
	membrosAposPagamento := applyMemberPayments(&contribuicao, membros)
	registrosPorTipo[contribuicaoId] = contribuicao

	registrosEspecificos := extractRegistrosEspecificosDeNucleo(registrosPorTipo)

	jornalId := db.GetTiposDeRegistroDB().GetJornalId()
	var parcela_partidaria_jornal float64
	for _, t := range tiposDeRegistros {
		if t.ID == jornalId {
			parcela_partidaria_jornal = t.Parcela_partidaria
			break
		}
	}

	reportId := strconv.Itoa(nucleoId) + "-" + month + "/" + year

	total_ganho := registrosEspecificos.Total + registrosPorTipo[contribuicaoId].Total + registrosPorTipo[jornalId].Total
	total_liquido := total_ganho - registrosPorTipo[gastosId].Total

	partilha_partidaria_especifica := calcPartilhaPartidariaEspecifica(tiposDeRegistros, registrosEspecificos.Tipos)
	partilha_partidaria := partilha_partidaria_especifica + registrosPorTipo[contribuicaoId].Total + float64(len(registrosPorTipo[jornalId].Registros))*parcela_partidaria_jornal

	lucro_nucleo := total_liquido - partilha_partidaria

	report := models.Relatorio_mensal_complexo{
		ID:                      reportId,
		Mes:                     month,
		Ano:                     year,
		Membros:                 membrosMAP,
		Membros_apos_pagamentos: membrosAposPagamento,
		Nucleo:                  nucleo,
		Partido:                 partido,
		Pagamentos_de_membros:   registrosPorTipo[contribuicaoId],
		Vendas_jornal:           registrosPorTipo[db.GetTiposDeRegistroDB().GetJornalId()],
		Gastos:                  registrosPorTipo[gastosId],
		Registros_especificos:   registrosEspecificos,
		Total_Ganho:             total_ganho,
		Total_Liquido:           total_liquido,
		Pagamento_Partidario:    partilha_partidaria,
		Lucro_Nucleo:            lucro_nucleo,
		Link_Arquivo:            "",
	}

	registro_partilha_partidaria := db.GetRegisterDB().GetNextId()

	partyDebtsRegister := models.Registro{
		ID:         registro_partilha_partidaria,
		Dia:        nucleo.Dia_de_Pagamento,
		Mes:        month,
		Ano:        year,
		Nucleo:     nucleoId,
		Pagante:    nucleoId,
		Cobrante:   db.GetPartidoDB().GetPartidoId(),
		Tipo:       database.GetDB().GetTiposDeRegistroDB().GetPagamentoPartido(),
		Descricao:  "Pagamento ao partido",
		Valor:      partilha_partidaria,
		Quantidade: 1,
	}

	return report, partyDebtsRegister, err
}

/* func GetYearlyReport(year string) (models.YearlyReport, error) {
	//TODO: Implement this
} */
