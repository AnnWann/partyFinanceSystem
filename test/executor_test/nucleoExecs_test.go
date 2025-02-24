package test

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/AnnWann/pstu_finance_system/src/database"
	"github.com/AnnWann/pstu_finance_system/src/executors"
	"github.com/AnnWann/pstu_finance_system/src/models"
	test_helpers "github.com/AnnWann/pstu_finance_system/test"
)

func TestAddNucleo(t *testing.T) {
	file := "nucleo1/add/test.db"

	db := test_helpers.BuildEnviroment(t, file, nil)
	defer db.Close()
	defer os.RemoveAll("nucleo1")

	expectedNucleos := []models.Nucleo{
		{
			ID:               3,
			Nome:             "nucleo1",
			Cidade:           "cidade1",
			Estado:           "estado1",
			Reserva:          0,
			Dia_de_Pagamento: "01",
			Administrador:    1,
		},
		{
			ID:               4,
			Nome:             "nucleo2",
			Cidade:           "cidade2",
			Estado:           "estado2",
			Reserva:          0,
			Dia_de_Pagamento: "02",
			Administrador:    1,
		},
		{
			ID:               5,
			Nome:             "nucleo3",
			Cidade:           "cidade3",
			Estado:           "estado3",
			Reserva:          0,
			Dia_de_Pagamento: "03",
			Administrador:    1,
		},
	}

	for _, nucleo := range expectedNucleos {
		id, err := executors.AddNucleo(nucleo.Nome, nucleo.Cidade, nucleo.Estado, nucleo.Dia_de_Pagamento, strconv.Itoa(nucleo.Administrador))
		if err != nil {
			t.Fatal(err)
		}

		if id != nucleo.ID {
			t.Fatalf("Expected id %d, got %d", nucleo.ID, id)
		}

		n, err := database.GetDB().GetNucleoDB().GetNucleoById(id)
		if err != nil {
			fmt.Println(err)
			t.Fatal(err)
		}

		if !reflect.DeepEqual(n, nucleo) {
			t.Fatalf("Expected nucleo %+v, got %+v", nucleo, n)
		}
	}
}

func TestGetNucleo(t *testing.T) {
	file := "nucleo2/get/test.db"
	sqlCommands := []string{
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (1, 'nucleo1', 'cidade1', 'estado1', 0, '01');",
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (2, 'nucleo2', 'cidade2', 'estado2', 0, '02');",
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (3, 'nucleo3', 'cidade3', 'estado3', 0, '03');",
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (4, 'nucleo4', 'cidade2', 'estado2', 0, '04');",
	}

	db := test_helpers.BuildEnviroment(t, file, sqlCommands)
	defer db.Close()
	defer os.RemoveAll("nucleo2")

	expectedNucleos := []models.Nucleo{
		{
			ID:               1,
			Nome:             "nucleo1",
			Cidade:           "cidade1",
			Estado:           "estado1",
			Reserva:          0,
			Dia_de_Pagamento: "01",
		},
		{
			ID:               2,
			Nome:             "nucleo2",
			Cidade:           "cidade2",
			Estado:           "estado2",
			Reserva:          0,
			Dia_de_Pagamento: "02",
		},
		{
			ID:               3,
			Nome:             "nucleo3",
			Cidade:           "cidade3",
			Estado:           "estado3",
			Reserva:          0,
			Dia_de_Pagamento: "03",
		},
		{
			ID:               4,
			Nome:             "nucleo4",
			Cidade:           "cidade2",
			Estado:           "estado2",
			Reserva:          0,
			Dia_de_Pagamento: "04",
		},
	}

	nucleos, err := executors.GetNucleo(nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(nucleos) != len(expectedNucleos) {
		t.Fatalf("Expected %d nucleos, got %d", len(expectedNucleos), len(nucleos))
	}

	for i, nucleo := range expectedNucleos {
		if !reflect.DeepEqual(nucleos[i], nucleo) {
			t.Fatalf("Expected nucleo %+v, got %+v", nucleo, nucleos[i])
		}
	}

	nucleos, err = executors.GetNucleo(map[string]string{"--cidade": "cidade2"})
	if err != nil {
		t.Fatal(err)
	}

	if len(nucleos) != 2 {
		t.Fatalf("Expected 2 nucleos, got %d", len(nucleos))
	}

	expectedNucleosWithout1And3 := []models.Nucleo{
		{
			ID:               2,
			Nome:             "nucleo2",
			Cidade:           "cidade2",
			Estado:           "estado2",
			Reserva:          0,
			Dia_de_Pagamento: "02",
		},

		{
			ID:               4,
			Nome:             "nucleo4",
			Cidade:           "cidade2",
			Estado:           "estado2",
			Reserva:          0,
			Dia_de_Pagamento: "04",
		},
	}

	for i, nucleo := range expectedNucleosWithout1And3 {
		if !reflect.DeepEqual(nucleos[i], nucleo) {
			t.Fatalf("Expected nucleo %+v, got %+v", nucleo, nucleos[i])
		}
	}
}

func TestDeleteNucleo(t *testing.T) {
	file := "nucleo3/delete/test.db"
	sqlCommands := []string{
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (1, 'nucleo1', 'cidade1', 'estado1', 0, '01');",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES (1, 'pessoa1', 1, 1, 100, 0);",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES (2, 'pessoa2', 1, 1, 200, 0);",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES (3, 'pessoa3', 1, 1, 300, 0);",
	}

	db := test_helpers.BuildEnviroment(t, file, sqlCommands)
	defer db.Close()
	defer os.RemoveAll("nucleo3")

	err := executors.DeleteNucleo("1")
	if err != nil {
		t.Fatal(err)
	}

	nucleo, err := database.GetDB().GetNucleoDB().GetNucleoById(1)
	if err != nil {
		t.Fatal(err)
	}
	if nucleo != (models.Nucleo{}) {
		t.Fatalf("Expected empty nucleo, got %+v", nucleo)
	}

	p1, err := database.GetDB().GetMembroDB().GetMembroById(1)
	if err != nil {
		t.Fatal(err)
	}
	if p1 != (models.Membro{}) {
		t.Fatalf("Expected empty person, got %+v", p1)
	}

	p2, err := database.GetDB().GetMembroDB().GetMembroById(2)
	if err != nil {
		t.Fatal(err)
	}
	if p2 != (models.Membro{}) {
		t.Fatalf("Expected empty person, got %+v", p2)
	}

	p3, err := database.GetDB().GetMembroDB().GetMembroById(3)
	if err != nil {
		t.Fatal(err)
	}
	if p3 != (models.Membro{}) {
		t.Fatalf("Expected empty person, got %+v", p3)
	}

}

func TestUpdatePayday(t *testing.T) {
	file := "nucleo4/update/test.db"
	sqlCommands := []string{
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (1, 'nucleo1', 'cidade1', 'estado1', 0, '01');",
	}

	db := test_helpers.BuildEnviroment(t, file, sqlCommands)
	defer db.Close()
	defer os.RemoveAll("nucleo4")

	err := executors.UpdatePayday("1", "02")
	if err != nil {
		t.Fatal(err)
	}

	nucleo, err := database.GetDB().GetNucleoDB().GetNucleoById(1)
	if err != nil {
		t.Fatal(err)
	}

	if nucleo.Dia_de_Pagamento != "02" {
		t.Fatalf("Expected payday 02, got %s", nucleo.Dia_de_Pagamento)
	}

}
