package test

import (
	"os"
	"testing"

	pdfMaker "github.com/AnnWann/pstu_finance_system/src/pdf"
	"github.com/AnnWann/pstu_finance_system/src/reportManager"
	test_helpers "github.com/AnnWann/pstu_finance_system/test"
	"github.com/joho/godotenv"
	"github.com/unidoc/unipdf/v3/common/license"
)

func TestPdf(t *testing.T) {
	file := "pdf/test.db"
	sqlCommands := []string{
		"insert into pessoas (id, classe) values (3, 'Nucleo')",
		"insert into pessoas (id, classe) values (4, 'Membro')",
		"insert into pessoas (id, classe) values (5, 'Membro')",
		"insert into pessoas (id, classe) values (6, 'Membro')",
		"insert into pessoas (id, classe) values (7, 'Membro')",
		"insert into pessoas (id, classe) values (8, 'Membro')",
		"insert into pessoas (id, classe) values (9, 'Membro')",
		"insert into pessoas (id, classe) values (10, 'Membro')",
		"insert into pessoas (id, classe) values (11, 'Membro')",
		"insert into nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) values (3, 'Nucleo 1', 'Cidade 1', 'Estado 1', 0, '01')",
		"insert into membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) values (4, 'Membro 1', 3, -300, 10, 0)",
		"insert into membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) values (5, 'Membro 2', 3, -400, 10, 0)",
		"insert into membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) values (6, 'Membro 3', 3, -200, 10, 0)",
		"insert into membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) values (7, 'Membro 4', 3, -200, 10, 0)",
		"insert into membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) values (8, 'Membro 5', 3, -200, 10, 0)",
		"insert into membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) values (9, 'Membro 6', 3, -200, 10, 3)",
		"insert into membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) values (10, 'Membro 7', 3, -100, 10, 7)",
		"insert into membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) values (11, 'Membro 8', 3, -100, 10, 10)",
		"insert into tipos_de_registro (id, nome, nucleo, descricao, partilha_partidaria) values (1, 'Tipo 1', 3, 'Descricao 1', 1)",
		"insert into tipos_de_registro (id, nome, nucleo, descricao, partilha_partidaria) values (2, 'Tipo 2', 3, 'Descricao 2', 1)",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (1, '01', '01', '2020', -100, 3, 4, 3, 1, 10, 'Descricao 1')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (2, '01', '01', '2020', -100, 3, 5, 3, 1, 10, 'Descricao 2')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (3, '01', '01', '2020', -100, 3, 6, 3, 1, 10, 'Descricao 3')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (4, '01', '01', '2020', -100, 3, 7, 3, 1, 10, 'Descricao 4')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (5, '01', '01', '2020', -100, 3, 8, 3, 1, 10, 'Descricao 5')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (6, '01', '01', '2020', -100, 3, 9, 3, 1, 5, 'Descricao 6')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (7, '01', '01', '2020', -100, 3, 10, 3, 1, 5, 'Descricao 7')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (8, '01', '01', '2020', -200, 3, 0, 4, 1, 4, 'Descricao 8')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (9, '01', '01', '2020', -200, 3, 0, 5, 1, 4, 'Descricao 9')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (10, '01', '01', '2020', -200, 3, 0, 6, 1, 4, 'Descricao 10')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (11, '01', '01', '2020', -400, 3, 4, 0, 1, 10, 'Descricao 11')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (12, '01', '01', '2020', -400, 3, 5, 0, 1, 10, 'Descricao 12')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (13, '01', '01', '2020', -400, 3, 6, 0, 1, 10, 'Descricao 13')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (14, '01', '01', '2020', 1, 3, 7, 3, 1, 10, 'Descricao 14')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (15, '01', '01', '2020', 1, 3, 7, 3, 1, 10, 'Descricao 15')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (16, '01', '01', '2020', 1, 3, 7, 3, 1, 10, 'Descricao 16')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (17, '01', '01', '2020', 2, 3, 7, 3, 1, 4, 'Descricao 8')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (18, '01', '01', '2020', 2, 3, 7, 3, 1, 4, 'Descricao 9')",
	}

	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal(err)
	}

	key := os.Getenv("UNIDOC_KEY")
	err = license.SetMeteredKey(key)
	if err != nil {
		t.Fatal(err)
	}

	db := test_helpers.BuildEnviroment(t, file, sqlCommands)

	defer db.Close()
	defer os.RemoveAll("pdf")

	report, _, err := reportManager.BuildRelatorioMensal("01", "2020", 3)
	if err != nil {
		t.Fatal(err)
	}

	folder := os.Getenv("PDF_FOLDER")

	err = pdfMaker.PrintPDFMonthlyReport(report, folder)
	if err != nil {
		t.Fatal(err)
	}
}
