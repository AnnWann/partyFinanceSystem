package test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/AnnWann/pstu_finance_system/src/database"
	test_helpers "github.com/AnnWann/pstu_finance_system/test"
	"github.com/joho/godotenv"
)

func TestConnection(t *testing.T) {
	db := database.GetDB()
	if db == nil {
		t.Fatal("Could not get DBWrapper obj")
	}

	godotenv.Load("../../.env")

	file := "connection/test.db"
	defer os.RemoveAll("connection")
	err := db.InitDB(file)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Could not init database")
	}

	expectedTables := []string{"pessoas", "partido", "nucleos", "designacao", "membros", "tipos_de_registro", "registros", "relatorios_mensais"}

	for _, table := range expectedTables {
		if !test_helpers.TableExists(db, table) {
			t.Errorf("Table %s not found", table)
		}
	}
}

func TestConnectionWithManyDirectories(t *testing.T) {
	db := database.GetDB()

	if db == nil {
		t.Fatal("Could not get DBWrapper obj")
	}

	file := "connection1/many/directories/test.db"
	defer os.RemoveAll("connection1")
	err := db.InitDB(file)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Could not init database")
	}

	expectedTables := []string{"pessoas", "partido", "nucleos", "designacao", "membros", "tipos_de_registro", "registros", "relatorios_mensais"}

	for _, table := range expectedTables {
		if !test_helpers.TableExists(db, table) {
			t.Errorf("Table %s not found", table)
		}
	}
}

func TestConnectionWithDirectoryCreated(t *testing.T) {
	db := database.GetDB()

	if db == nil {
		t.Fatal("Could not get DBWrapper obj")
	}

	file := "connection2/with/directory/created/test.db"
	defer os.RemoveAll("connection2")
	directories := filepath.Dir(file)
	err := os.MkdirAll(directories, os.ModePerm)
	if err != nil {
		t.Fatal("Could not create directories")
	}

	err = db.InitDB(file)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Could not init database")
	}

	expectedTables := []string{"pessoas", "partido", "nucleos", "designacao", "membros", "tipos_de_registro", "registros", "relatorios_mensais"}

	for _, table := range expectedTables {
		if !test_helpers.TableExists(db, table) {
			t.Errorf("Table %s not found", table)
		}
	}
}

func TestConnectionWithFileCreated(t *testing.T) {
	db := database.GetDB()

	if db == nil {
		t.Fatal("Could not get DBWrapper obj")
	}

	file := "connection3/with/file/created/test.db"
	defer os.RemoveAll("connection3")
	directories := filepath.Dir(file)
	err := os.MkdirAll(directories, os.ModePerm)
	if err != nil {
		t.Fatal("Could not create directories")
	}

	_, err = os.Create(file)
	if err != nil {
		t.Fatal("Could not create file")
	}

	err = db.InitDB(file)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Could not init database")
	}

	expectedTables := []string{"pessoas", "partido", "nucleos", "designacao", "membros", "tipos_de_registro", "registros", "relatorios_mensais"}

	for _, table := range expectedTables {
		if !test_helpers.TableExists(db, table) {
			t.Errorf("Table %s not found", table)
		}
	}
}

func TestReConnection(t *testing.T) {
	db := database.GetDB()

	if db == nil {
		t.Fatal("Could not get DBWrapper obj")
	}

	file := "connection4/re/connection/test.db"
	defer os.RemoveAll("connection4")

	err := db.InitDB(file)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Could not init database")
	}

	expectedTables := []string{"pessoas", "partido", "nucleos", "designacao", "membros", "tipos_de_registro", "registros", "relatorios_mensais"}

	db.Close()
	err = db.InitDB(file)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		t.Fatal("Could not re-init database")
	}

	for _, table := range expectedTables {
		if !test_helpers.TableExists(db, table) {
			t.Errorf("Table %s not found", table)
		}
	}
}