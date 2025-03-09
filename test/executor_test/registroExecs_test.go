package test

import (
	"os"
	"testing"

	"github.com/AnnWann/pstu_finance_system/src/database"
	"github.com/AnnWann/pstu_finance_system/src/executors"
	test_helpers "github.com/AnnWann/pstu_finance_system/test"
)

func TestAddRegistroSucceding(t *testing.T) {
	file := "register1/add/test.db"
	sqlCommands := []string{
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento, administrador) VALUES (3, 'nucleo1', 'cidade1', 'estado1', 0, '01', 1)",
		"INSERT INTO tipos_de_registro (id, nome, nucleo, descricao, partilha_partidaria) VALUES (1, 'tipo1', 1, 'descricao1', 0.5)",
		"INSERT INTO pessoas (id, classe) VALUES (3, 'nucleo')",
		"INSERT INTO pessoas (id, classe) VALUES (4, 'membro')",
		"INSERT INTO pessoas (id, classe) VALUES (5, 'membro')",
	}

	db := test_helpers.BuildEnviroment(t, file, sqlCommands)
	defer db.Close()
	defer os.RemoveAll("register1")

	id, err := executors.AddRegistro("01", "01", "2020", "3", "-100", "4", "5", "1", "1.0", "description")
	if err != nil {
		t.Fatal(err)
	}

	register, err := database.GetDB().GetRegisterDB().GetRegisterById(id)
	if err != nil {
		t.Fatal(err)
	}
	if register.ID != id {
		t.Errorf("Expected id %d, got %d", id, register.ID)
	}
	if register.Dia != "01" {
		t.Errorf("Expected day 01, got %s", register.Dia)
	}
	if register.Mes != "01" {
		t.Errorf("Expected month 01, got %s", register.Mes)
	}
	if register.Ano != "2020" {
		t.Errorf("Expected year 2020, got %s", register.Ano)
	}
	if register.Nucleo != 3 {
		t.Errorf("Expected nucleo 1, got %d", register.Nucleo)
	}
	if register.Tipo != -100 {
		t.Errorf("Expected type 1, got %d", register.Tipo)
	}
	if register.Pagante != 4 {
		t.Errorf("Expected Pago_por 4, got %d", register.Pagante)
	}
	if register.Cobrante != 5 {
		t.Errorf("Expected Cobrado_por 5, got %d", register.Cobrante)
	}
	if register.Quantidade != 1 {
		t.Errorf("Expected amount 1, got %d", register.Quantidade)
	}
	if register.Valor != 1.0 {
		t.Errorf("Expected value 1.0, got %f", register.Valor)
	}
	if register.Descricao != "description" {
		t.Errorf("Expected description description, got %s", register.Descricao)
	}
}
