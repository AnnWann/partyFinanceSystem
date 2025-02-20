package test_helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/AnnWann/pstu_finance_system/src/database"
)

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func TableExists(db *database.DBWrapper, tableName string) bool {
	var name string
	query := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' AND name='%s';", tableName)
	err := db.QueryRow(query).Scan(&name)
	return err == nil && name == tableName
}

func BuildEnviroment(t *testing.T, file string, sqlCommands []string) *database.DBWrapper {
	db := database.GetDB()
	err := db.InitDB(file)
	if err != nil {
		fmt.Println(err)
		os.RemoveAll(getBaseDir(file))
		t.Fatal("Could not init database")
		// remove the base directory

	}

	for _, command := range sqlCommands {
		_, err = db.Exec(command)
		if err != nil {
			fmt.Println(err)
			os.RemoveAll(getBaseDir(file))
			t.Fatalf("Could not execute command: %s\nerror: %s", command, err)
		}

	}

	return db
}

func getBaseDir(file string) string {
	for {
		f := filepath.Dir(file)
		if f == "." {
			return file
		} else {
			file = f
		}
	}
}
