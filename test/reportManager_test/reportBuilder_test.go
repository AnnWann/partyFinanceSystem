package test

import (
	"os"
	"reflect"
	"testing"

	"github.com/AnnWann/pstu_finance_system/src/models"
	"github.com/AnnWann/pstu_finance_system/src/reportManager"
	test_helpers "github.com/AnnWann/pstu_finance_system/test"
)

func TestBuildRelatorioMensal(t *testing.T) {
	file := "reportManager1/buildRelatorioMensal/test.db"
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
		"insert into nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento, administrador) values (3, 'Nucleo 1', 'Cidade 1', 'Estado 1', 0, '01', 1)",
		"insert into membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) values (4, 'Membro 1', 3, -300, 10, 0)",
		"insert into membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) values (5, 'Membro 2', 3, -400, 10, 0)",
		"insert into membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) values (6, 'Membro 3', 3, -200, 10, 0)",
		"insert into membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) values (7, 'Membro 4', 3, -200, 10, 0)",
		"insert into membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) values (8, 'Membro 5', 3, -200, 10, 0)",
		"insert into membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) values (9, 'Membro 6', 3, -200, 10, 0)",
		"insert into membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) values (10, 'Membro 7', 3, -100, 10, 5)",
		"insert into membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) values (11, 'Membro 8', 3, -100, 10, 10)",
		"insert into tipos_de_registro (id, nome, nucleo, descricao, partilha_partidaria) values (1, 'Tipo 1', 3, 'Descricao 1', 1)",
		"insert into tipos_de_registro (id, nome, nucleo, descricao, partilha_partidaria) values (2, 'Tipo 2', 3, 'Descricao 2', 1)",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (1, '01', '01', '2020', -100, 3, 4, 3, 1, 10, 'Descricao 1')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (2, '01', '01', '2020', -100, 3, 5, 3, 1, 10, 'Descricao 2')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (3, '01', '01', '2020', -100, 3, 6, 3, 1, 10, 'Descricao 3')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (4, '01', '01', '2020', -100, 3, 7, 3, 1, 10, 'Descricao 4')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (5, '01', '01', '2020', -100, 3, 8, 3, 1, 10, 'Descricao 5')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (6, '01', '01', '2020', -100, 3, 9, 3, 1, 10, 'Descricao 6')",
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
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (17, '01', '01', '2020', 2, 3, 7, 3, 1, 10, 'Descricao 17')",
		"insert into registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) values (18, '01', '01', '2020', 2, 3, 7, 3, 1, 10, 'Descricao 18')",
	}

	db := test_helpers.BuildEnviroment(t, file, sqlCommands)

	defer db.Close()
	defer os.RemoveAll("reportManager1")

	expectedId := "3-01/2020"
	expectedMembros := map[int]models.Membro{
		4: {
			ID:                  4,
			Nome:                "Membro 1",
			Nucleo:              3,
			Cargo:               -300,
			Contribuicao_mensal: 10,
			Credito:             0,
		},
		5: {
			ID:                  5,
			Nome:                "Membro 2",
			Nucleo:              3,
			Cargo:               -400,
			Contribuicao_mensal: 10,
			Credito:             0,
		},
		6: {
			ID:                  6,
			Nome:                "Membro 3",
			Nucleo:              3,
			Cargo:               -200,
			Contribuicao_mensal: 10,
			Credito:             0,
		},
		7: {
			ID:                  7,
			Nome:                "Membro 4",
			Nucleo:              3,
			Cargo:               -200,
			Contribuicao_mensal: 10,
			Credito:             0,
		},
		8: {
			ID:                  8,
			Nome:                "Membro 5",
			Nucleo:              3,
			Cargo:               -200,
			Contribuicao_mensal: 10,
			Credito:             0,
		},
		9: {
			ID:                  9,
			Nome:                "Membro 6",
			Nucleo:              3,
			Cargo:               -200,
			Contribuicao_mensal: 10,
			Credito:             0,
		},
		10: {
			ID:                  10,
			Nome:                "Membro 7",
			Nucleo:              3,
			Cargo:               -100,
			Contribuicao_mensal: 10,
			Credito:             5,
		},
		11: {
			ID:                  11,
			Nome:                "Membro 8",
			Nucleo:              3,
			Cargo:               -100,
			Contribuicao_mensal: 10,
			Credito:             10,
		},
	}
	expectedMembros_apos_pagamentos := map[int]models.Membro{
		4: {
			ID:                  4,
			Nome:                "Membro 1",
			Nucleo:              3,
			Cargo:               -300,
			Contribuicao_mensal: 10,
			Credito:             0,
		},
		5: {
			ID:                  5,
			Nome:                "Membro 2",
			Nucleo:              3,
			Cargo:               -400,
			Contribuicao_mensal: 10,
			Credito:             0,
		},
		6: {
			ID:                  6,
			Nome:                "Membro 3",
			Nucleo:              3,
			Cargo:               -200,
			Contribuicao_mensal: 10,
			Credito:             0,
		},
		7: {
			ID:                  7,
			Nome:                "Membro 4",
			Nucleo:              3,
			Cargo:               -200,
			Contribuicao_mensal: 10,
			Credito:             0,
		},
		8: {
			ID:                  8,
			Nome:                "Membro 5",
			Nucleo:              3,
			Cargo:               -200,
			Contribuicao_mensal: 10,
			Credito:             0,
		},
		9: {
			ID:                  9,
			Nome:                "Membro 6",
			Nucleo:              3,
			Cargo:               -200,
			Contribuicao_mensal: 10,
			Credito:             0,
		},
		10: {
			ID:                  10,
			Nome:                "Membro 7",
			Nucleo:              3,
			Cargo:               -100,
			Contribuicao_mensal: 10,
			Credito:             0,
		},
		11: {
			ID:                  11,
			Nome:                "Membro 8",
			Nucleo:              3,
			Cargo:               -100,
			Contribuicao_mensal: 10,
			Credito:             0,
		},
	}
	expectedNucleo := models.Nucleo{
		ID:               3,
		Nome:             "Nucleo 1",
		Cidade:           "Cidade 1",
		Estado:           "Estado 1",
		Reserva:          0,
		Dia_de_Pagamento: "01",
		Administrador:    1,
	}
	expectedTotalMembros := 80
	expectedTotalJornal := 12
	expectedTotalGastos := 30
	expectedTotalRegistrosEspecificos1 := 30
	expectedTotalRegistrosEspecificos2 := 20
	ExpectedTotal_Ganho := expectedTotalMembros + expectedTotalJornal + expectedTotalRegistrosEspecificos1 + expectedTotalRegistrosEspecificos2
	ExpectedTotal_Liquido := ExpectedTotal_Ganho - expectedTotalGastos
	ExpectedPagamento_Partidario := expectedTotalMembros + 2*3 + 3 + 2
	ExpectedLucro_Nucleo := ExpectedTotal_Liquido - ExpectedPagamento_Partidario

	expected_Pagamento_partido := models.Registro{
		ID:         19,
		Dia:        "01",
		Mes:        "01",
		Ano:        "2020",
		Tipo:       -300,
		Nucleo:     3,
		Pagante:    3,
		Cobrante:   1,
		Quantidade: 1,
		Valor:      float64(ExpectedPagamento_Partidario),
		Descricao:  "Pagamento ao administrador",
	}

	result, pagamento_partido, err := reportManager.BuildRelatorioMensal("01", "2020", 3)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if !reflect.DeepEqual(expectedId, result.ID) {
		t.Errorf("Para ID -> Expected: %v\nGot: %v", expectedId, result.ID)
	}

	if !reflect.DeepEqual(expectedMembros, result.Membros) {
		t.Errorf("Para Membros -> Expected: %v\nGot: %v", expectedMembros, result.Membros)
	}

	if !reflect.DeepEqual(expectedMembros_apos_pagamentos, result.Membros_apos_pagamentos) {
		t.Errorf("Para Membros Após Pagamento -> Expected: %v\nGot: %v", expectedMembros_apos_pagamentos, result.Membros_apos_pagamentos)
	}

	if !reflect.DeepEqual(expectedNucleo, result.Nucleo) {
		t.Errorf("Para Nucleo -> Expected: %v\nGot: %v", expectedNucleo, result.Nucleo)
	}

	if expectedTotalMembros != int(result.Pagamentos_de_membros.Total) {
		t.Errorf("Para Total Membros -> Expected: %v\nGot: %v", expectedTotalMembros, result.Pagamentos_de_membros.Total)
	}

	if expectedTotalJornal != int(result.Vendas_jornal.Total) {
		t.Errorf("Para Total Jornal -> Expected: %v\nGot: %v", expectedTotalJornal, result.Vendas_jornal.Total)
	}

	if expectedTotalGastos != int(result.Gastos.Total) {
		t.Errorf("Para Total Gastos -> Expected: %v\nGot: %v", expectedTotalGastos, result.Gastos.Total)
	}

	if expectedTotalRegistrosEspecificos1 != int(result.Registros_especificos.Tipos[1].Total) {
		t.Errorf("Para Total Registros Especificos 1 -> Expected: %v\nGot: %v", expectedTotalRegistrosEspecificos1, result.Registros_especificos.Tipos[1].Total)
	}

	if expectedTotalRegistrosEspecificos2 != int(result.Registros_especificos.Tipos[2].Total) {
		t.Errorf("Para Total Registros Especificos 2 -> Expected: %v\nGot: %v", expectedTotalRegistrosEspecificos2, result.Registros_especificos.Tipos[2].Total)
	}

	if ExpectedTotal_Ganho != int(result.Total_Ganho) {
		t.Errorf("Para Total Ganho -> Expected: %v\nGot: %v", ExpectedTotal_Ganho, result.Total_Ganho)
	}

	if ExpectedTotal_Liquido != int(result.Total_Liquido) {
		t.Errorf("Para Total Liquido -> Expected: %v\nGot: %v", ExpectedTotal_Liquido, result.Total_Liquido)
	}

	if ExpectedPagamento_Partidario != int(result.Pagamento_Partidario) {
		t.Errorf("Para Pagamento Partidario -> Expected: %v\nGot: %v", ExpectedPagamento_Partidario, result.Pagamento_Partidario)
	}

	if ExpectedLucro_Nucleo != int(result.Lucro_Nucleo) {
		t.Errorf("Para Lucro Nucleo -> Expected: %v\nGot: %v", ExpectedLucro_Nucleo, result.Lucro_Nucleo)
	}

	if !reflect.DeepEqual(expected_Pagamento_partido, pagamento_partido) {
		t.Errorf("PagamentoPartido -> Expected: %v\nGot: %v", expected_Pagamento_partido, pagamento_partido)
	}
}
