package conn_test

import (
	"testing"
	"reflect"

	"github.com/AnnWann/pstu_finance_system/src/database"
)

func TestConnect(t *testing.T) {
	file := "test/database_test/testConn.db"
	conn, err := database.GetDB().InitDB(file)
	if err != nil {
		t.Error(err)
	}

	tables := []string{"partido", "nucleo", "pessoas", "tipo_de_registro", "registro", "relatorio", "dia_de_pagamento"}
