package test

import (
	"os"
	"testing"

	"github.com/AnnWann/pstu_finance_system/src/executors"
	test_helpers "github.com/AnnWann/pstu_finance_system/test"
)

func TestGetPartido(t *testing.T) {
	file := "partido1/get/test.db"
	db := test_helpers.BuildEnviroment(t, file, nil)

	defer db.Close()
	defer os.RemoveAll("partido1")

	_, err := executors.GetPartido()
	if err != nil {
		println(err.Error())
		t.Error("Erro ao obter partido - 1")
	}
}
