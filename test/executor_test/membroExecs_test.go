package test

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/AnnWann/pstu_finance_system/src/database"
	"github.com/AnnWann/pstu_finance_system/src/executors"
	test_helpers "github.com/AnnWann/pstu_finance_system/test"
)

func TestAddMembroSucceding(t *testing.T) {
	file := "person1/add/succeding/test.db"
	sqlCommands := []string{
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (1, 'nucleo1', 'cidade1', 'estado1', 0, '01')",
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (2, 'nucleo2', 'cidade2', 'estado2', 0, '02')",
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (3, 'nucleo3', 'cidade3', 'estado3', 0, '03')"}

	db := test_helpers.BuildEnviroment(t, file, sqlCommands)

	defer db.Close()
	defer os.RemoveAll("person1")

	membrosEsperados := []struct {
		Signature int
		Name      string
		Nucleo    string
		cargo     int
	}{
		{1, "person1", "1", database.GetDB().GetCargoDB().GetDirigenteId()},
		{2, "person2", "2", database.GetDB().GetCargoDB().GetDirigenteId()},
		{3, "person3", "3", database.GetDB().GetCargoDB().GetDirigenteId()},
		{4, "person4", "2", database.GetDB().GetCargoDB().GetDirigenteFinanceiroId()},
		{5, "person5", "1", database.GetDB().GetCargoDB().GetDirigenteFinanceiroId()},
		{6, "person6", "1", database.GetDB().GetCargoDB().GetAspiranteId()},
		{7, "person7", "1", database.GetDB().GetCargoDB().GetAspiranteId()},
	}

	for _, membro := range membrosEsperados {
		id, err := executors.AddMembro(membro.Name, membro.Nucleo)
		if err != nil {
			t.Errorf("At signature %d, could not add person", membro.Signature)
			fmt.Printf("Sig %d: %s\n", membro.Signature, err)
			continue
		}

		membroOBJ, err := db.GetMembroDB().GetMembroById(id)
		if err != nil {
			t.Errorf("At signature %d, could not get person", membro.Signature)
			fmt.Printf("Sig %d: %s\n", membro.Signature, err)
			continue
		}

		if membroOBJ.Nome != membro.Name {
			t.Errorf("At signature %d, expected name %s, got %s", membro.Signature, membro.Name, membroOBJ.Nome)
			continue
		}

		if membroOBJ.Cargo != membro.cargo {
			t.Errorf("At signature %d, expected role %d, got %d", membro.Signature, membro.cargo, membroOBJ.Cargo)
			continue
		}
	}

}

func TestAddPersonFailing(t *testing.T) {
	file := "person2/add/failing/test.db"
	sqlCommands := []string{
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (1, 'nucleo1', 'cidade1', 'estado1', 0, '01')",
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (2, 'nucleo2', 'cidade2', 'estado2', 0, '02')",
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (3, 'nucleo3', 'cidade3', 'estado3', 0, '03')"}

	db := test_helpers.BuildEnviroment(t, file, sqlCommands)

	defer db.Close()
	defer os.RemoveAll("person2")

	expectedmembros := []struct {
		Signature int
		Name      string
		Nucleo    string
	}{
		{1, "person1", "4"},
		{2, "person2", "8"},
		{3, "person3", "5"},
	}

	for _, person := range expectedmembros {
		_, err := executors.AddMembro(person.Name, person.Nucleo)
		if err == nil {
			t.Errorf("At signature %d, should not add person", person.Signature)
			continue
		}
	}
}

func TestGetMembro(t *testing.T) {
	file := "person3/get/test.db"
	sqlCommands := []string{
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (1, 'nucleo1', 'cidade1', 'estado1', 0, '01')",
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (2, 'nucleo2', 'cidade2', 'estado2', 0, '02')",
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (3, 'nucleo3', 'cidade3', 'estado3', 0, '03')",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('1', 'person1', 1, -300, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('2', 'person2', 2, -300, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('3', 'person3', 3, -300, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('4', 'person4', 2, -400, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('5', 'person5', 1, -400, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('6', 'person6', 1, -100, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('7', 'person7', 1, -200, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('8', 'person8', 1, -200, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('9', 'person9', 1, -100, 0, 0)"}

	db := test_helpers.BuildEnviroment(t, file, sqlCommands)

	defer db.Close()
	defer os.RemoveAll("person3")

	membrosEsperados := []struct {
		Signature int
		Id        int
		Name      string
		Nucleo    string
		Role      int
	}{
		{'1', 1, "person1", "1", -300},
		{'2', 2, "person2", "2", -300},
		{'3', 3, "person3", "3", -300},
		{'4', 4, "person4", "2", -400},
		{'5', 5, "person5", "1", -400},
		{'6', 6, "person6", "1", -100},
		{'7', 7, "person7", "1", -200},
		{'8', 8, "person8", "1", -200},
		{'9', 9, "person9", "1", -100},
	}

	membros, err := executors.GetMembro(nil)
	if err != nil {
		t.Fatal("Could not get membros")
	}

	if len(membros) != len(membrosEsperados) {
		t.Fatalf("Expected %d membros, got %d", len(membrosEsperados), len(membros))
	}

	for _, membro := range membrosEsperados {
		found := false
		for _, p := range membros {
			if p.ID == membro.Id {
				found = true
				if p.Nome != membro.Name {
					t.Errorf("At signature %d, expected name %s, got %s", membro.Signature, membro.Name, p.Nome)
				}
				if p.Cargo != membro.Role {
					t.Errorf("At signature %d, expected role %d, got %d", membro.Signature, membro.Role, p.Cargo)
				}
				break
			}
		}
		if !found {
			t.Errorf("At signature %d, could not find person", membro.Signature)
		}
	}

	membrosDoNucleo1QueSaoMilitantes, err := executors.GetMembro(map[string]string{"--nucleo": "1", "--cargo": "-200"})
	if err != nil {
		t.Fatal("Could not get membros from nucleo 1 who are also militants")
	}

	if len(membrosDoNucleo1QueSaoMilitantes) != 2 {
		t.Fatalf("Expected 2 membros, got %d", len(membrosDoNucleo1QueSaoMilitantes))
	}

	for _, membro := range membrosEsperados {
		if membro.Nucleo == "1" && membro.Role == -200 {
			found := false
			for _, m := range membrosDoNucleo1QueSaoMilitantes {
				if m.ID == membro.Id {
					found = true
					if m.Nome != membro.Name {
						t.Errorf("At signature %d, expected name %s, got %s", membro.Signature, membro.Name, m.Nome)
					}
					if m.Cargo != membro.Role {
						t.Errorf("At signature %d, expected role %d, got %d", membro.Signature, membro.Role, m.Cargo)
					}
					break
				}
			}
			if !found {
				t.Errorf("At signature %d, could not find person", membro.Signature)
			}
		}
	}
}

func TestUpdateMembro(t *testing.T) {
	file := "person4/update/test.db"
	sqlCommands := []string{
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (1, 'nucleo1', 'cidade1', 'estado1', 0, '01')",
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (2, 'nucleo2', 'cidade2', 'estado2', 0, '02')",
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (3, 'nucleo3', 'cidade3', 'estado3', 0, '03')",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('1', 'person1', 1, -300, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('2', 'person2', 2, -300, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('3', 'person3', 3, -300, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('4', 'person4', 2, -400, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('5', 'person5', 1, -400, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('6', 'person6', 1, -100, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('7', 'person7', 1, -200, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('8', 'person8', 1, -200, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('9', 'person9', 1, -100, 0, 0)"}

	db := test_helpers.BuildEnviroment(t, file, sqlCommands)

	defer db.Close()
	defer os.RemoveAll("person4")

	membrosEsperados := []struct {
		Signature           int
		Id                  int
		Name                string
		Nucleo              string
		Cargo               int
		Contribuicao_mensal float64
	}{
		{'1', 1, "person1", "1", -300, 0},
		{'2', 2, "person2", "2", -300, 30},
		{'3', 3, "person3", "3", -300, 10},
		{'4', 4, "person4", "2", -400, 20},
		{'5', 5, "person5", "1", -400, 0},
		{'6', 6, "person6", "2", -100, 10},
		{'7', 7, "person7", "2", -200, 0},
		{'8', 8, "person8", "2", -200, 0},
		{'9', 9, "person9", "1", -100, 0},
	}

	err := executors.UpdateMembro("2", "", "30")
	if err != nil {
		t.Errorf("Could not update person of signature 2")
	}

	err = executors.UpdateMembro("3", "", "10")
	if err != nil {
		t.Errorf("Could not update person of signature 3")
	}

	err = executors.UpdateMembro("4", "", "20")
	if err != nil {
		t.Errorf("Could not update person of signature 4")
	}

	err = executors.UpdateMembro("6", "2", "10")
	if err != nil {
		t.Errorf("Could not update person of signature 6")
	}

	err = executors.UpdateMembro("7", "2", "0")
	if err != nil {
		t.Errorf("Could not update person of signature 7")
	}

	err = executors.UpdateMembro("8", "2", "0")
	if err != nil {
		t.Errorf("Could not update person of signature 8")
	}

	membros, err := executors.GetMembro(nil)
	if err != nil {
		t.Fatal("Could not get membros, Error: ", err)
	}

	if len(membros) != len(membrosEsperados) {
		t.Fatalf("Expected %d membros, got %d", len(membrosEsperados), len(membros))
	}

	for _, membro := range membrosEsperados {
		found := false
		for _, p := range membros {
			if p.ID == membro.Id {
				found = true
				if p.Nome != membro.Name {
					t.Errorf("At signature %d, expected name %s, got %s", membro.Signature, membro.Name, p.Nome)
				}
				if p.Cargo != membro.Cargo {
					t.Errorf("At signature %d, expected role %d, got %d", membro.Signature, membro.Cargo, p.Cargo)
				}
				if p.Contribuicao_mensal != membro.Contribuicao_mensal {
					t.Errorf("At signature %d, expected party contribution %f, got %f", membro.Signature, membro.Contribuicao_mensal, p.Contribuicao_mensal)
				}
				break
			}
		}
		if !found {
			t.Errorf("At signature %d, could not find person", membro.Signature)
		}
	}

}

func TestPromoteMembro(t *testing.T) {
	file := "person5/promote/test.db"
	sqlCommands := []string{
		"INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (1, 'nucleo1', 'cidade1', 'estado1', 0, '01')",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('1', 'person1', 1, -300, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('2', 'person2', 1, -400, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('3', 'person3', 1, -100, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('4', 'person4', 1, -200, 0, 0)",
		"INSERT INTO membros (id, nome, nucleo, cargo, contribuicao_mensal, credito) VALUES ('5', 'person5', 1, -100, 0, 0)"}

	db := test_helpers.BuildEnviroment(t, file, sqlCommands)

	defer db.Close()
	defer os.RemoveAll("person5")

	designicacaoDB := database.GetDB().GetCargoDB()

	expectedMembro := []struct {
		Signature           int
		Id                  int
		Name                string
		cargo               int
		DesignicacaoTentada int
		cargoNova           int
	}{
		{1, 1, "person1", designicacaoDB.GetDirigenteId(), 0, designicacaoDB.GetMilitanteId()},                                                             //demovido para militante
		{2, 2, "person2", designicacaoDB.GetDirigenteFinanceiroId(), designicacaoDB.GetDirigenteFinanceiroId(), designicacaoDB.GetDirigenteFinanceiroId()}, //promovido para dirigente_financeiro, nada acontece
		{3, 3, "person3", designicacaoDB.GetAspiranteId(), designicacaoDB.GetMilitanteId(), designicacaoDB.GetMilitanteId()},                               //promovido para militante
		{4, 4, "person4", designicacaoDB.GetMilitanteId(), designicacaoDB.GetDirigenteId(), designicacaoDB.GetDirigenteId()},                               //promovido para dirigente
		{5, 5, "person5", designicacaoDB.GetAspiranteId(), designicacaoDB.GetDirigenteId(), designicacaoDB.GetAspiranteId()},                               //promovido para dirigente mas é proibido, então nada muda
	}

	for _, membro := range expectedMembro {
		if membro.DesignicacaoTentada == 0 {
			continue
		}
		err := executors.Promote(strconv.Itoa(membro.Id), strconv.Itoa(membro.DesignicacaoTentada))
		if err != nil {
			if membro.Signature != 5 && membro.Signature != 2 {
				t.Errorf("At signature %d, could not promote person\nError: %s", membro.Signature, err)
			}
		}
	}

	for _, membro := range expectedMembro {
		membroOBJ, err := db.GetMembroDB().GetMembroById(membro.Id)
		if err != nil {
			t.Errorf("At signature %d, could not get person", membro.Signature)
		}

		if membroOBJ.Cargo != membro.cargoNova {
			t.Errorf("At signature %d, expected role %d, got %d", membro.Signature, membro.cargoNova, membroOBJ.Cargo)
		}
	}

	p1, err := db.GetMembroDB().GetMembroById(1)
	if err != nil {
		t.Errorf("At signature 1, could not get person")
	}

	if p1.Cargo != -200 {
		t.Errorf("At signature 1, expected role -200, got %d", p1.Cargo)
	}
}
