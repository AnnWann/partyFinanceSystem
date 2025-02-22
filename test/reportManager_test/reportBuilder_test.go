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

	db := test_helpers.BuildEnviroment(t, file, sqlCommands)

	defer db.Close()
	defer os.RemoveAll("reportManager1")

	expectedId := "3-01/2020"
	expectedMembrosAposPagamento := map[int]models.Membro{
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
			Credito:             2,
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
	expectedPagamentosMembrosTotal := float64(78)
	expectedVendasJornalTotal := float64(12)
	expectedGastosTotal := float64(30)
	expectedRegistrosEspecificos1Total := float64(30)
	expectedRegistrosEspecificos2Total := float64(8)
	expectedRegistrosEspecificosTotal := float64(38)
	expectedTotalGanho := float64(128)
	expectedTotalLiquido := float64(98)
	expectedPagamentoPartidario := float64(89)
	expectedLucroNucleo := float64(9)

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
		Valor:      89,
		Descricao:  "Pagamento ao partido",
	}

	result, pagamento_partido, err := reportManager.BuildRelatorioMensal("01", "2020", 3)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if !reflect.DeepEqual(expected_Pagamento_partido, pagamento_partido) {
		t.Errorf("For pagamento partido. Expected: %v\nGot: %v", expected_Pagamento_partido, pagamento_partido)
	}

	if result.ID != expectedId {
		t.Errorf("For id. Expected. %v\nGot: %v", expectedId, result.ID)
	}

	if !reflect.DeepEqual(expectedMembrosAposPagamento, result.Membros_apos_pagamentos) {
		t.Errorf("For Mmembros_apos_pagamento. Expected: %v\nGot: %v", expectedMembrosAposPagamento, result.Membros_apos_pagamentos)
	}

	if result.Pagamentos_de_membros.Total != expectedPagamentosMembrosTotal {
		t.Errorf("For pagamentos_de_membros.Total. Expected: %v\nGot: %v", expectedPagamentosMembrosTotal, result.Pagamentos_de_membros.Total)
	}

	if result.Vendas_jornal.Total != expectedVendasJornalTotal {
		t.Errorf("For Vendas_jornal.Total. Expected: %v\nGot: %v", expectedVendasJornalTotal, result.Vendas_jornal.Total)
	}

	if result.Gastos.Total != expectedGastosTotal {
		t.Errorf("For Gastos.Total. Expected: %v\nGot: %v", expectedGastosTotal, result.Gastos.Total)
	}

	if result.Registros_especificos.Tipos[1].Total != expectedRegistrosEspecificos1Total {
		t.Errorf("For Registros_especificos.Tipos[1].Total. Expected: %v\nGot: %v", expectedRegistrosEspecificos1Total, result.Registros_especificos.Tipos[1].Total)
	}

	if result.Registros_especificos.Tipos[2].Total != expectedRegistrosEspecificos2Total {
		t.Errorf("For Registros_especificos.Tipos[2].Total. Expected: %v\nGot: %v", expectedRegistrosEspecificos2Total, result.Registros_especificos.Tipos[2].Total)
	}

	if result.Registros_especificos.Total != expectedRegistrosEspecificosTotal {
		t.Errorf("For Registros_especificos.Total. Expected: %v\nGot: %v", expectedRegistrosEspecificosTotal, result.Registros_especificos.Total)
	}

	if result.Total_Ganho != expectedTotalGanho {
		t.Errorf("For Total_Ganho. Expected: %v\nGot: %v", expectedTotalGanho, result.Total_Ganho)
	}

	if result.Total_Liquido != expectedTotalLiquido {
		t.Errorf("For Total_Liquido. Expected: %v\nGot: %v", expectedTotalLiquido, result.Total_Liquido)
	}

	if result.Pagamento_Partidario != expectedPagamentoPartidario {
		t.Errorf("For Pagamento_Partidario. Expected: %v\nGot: %v", expectedPagamentoPartidario, result.Pagamento_Partidario)
	}

	if result.Lucro_Nucleo != expectedLucroNucleo {
		t.Errorf("For Lucro_Nucleo. Expected: %v\nGot: %v", expectedLucroNucleo, result.Lucro_Nucleo)
	}

}
