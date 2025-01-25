package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DBWrapper struct {
	*sql.DB
}

var DB *DBWrapper = nil

func GetDB() *DBWrapper {
	var db *sql.DB
	if DB == nil {
		DB = &DBWrapper{db}
	}
	return DB
}

func (db *DBWrapper) GetConnection() *sql.DB {
	return db.DB
}

func (db *DBWrapper) InitDB(file string) {
	var err error
	db.DB, err = sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}

	db.Exec("CREATE TABLE IF NOT EXISTS Partido (id INTEGER NOT NULL PRIMARY KEY," +
		"nome TEXT NOT NULL, " +
		"credito FLOAT NOT NULL)")

	db.Exec("CREATE TABLE Nucleo (id INTEGER NOT NULL PRIMARY KEY," + 
		"nome TEXT NOT NULL, " +
		"cidade TEXT NOT NULL, " +
		"estado TEXT NOT NULL, " +
		"credito FLOAT NOT NULL)")

	db.Exec("CREATE TABLE IF NOT EXISTS pessoas (id INTEGER NOT NULL PRIMARY KEY," +
		"nome TEXT NOT NULL, " +
		"papel TEXT NOT NULL, " +
		"nucleo INTEGER NOT NULL, " +
		"contribuicao_mensal REAL NOT NULL, " +
		"credito FLOAT NOT NULL), " + 
		"FOREIGN KEY (nucleo) REFERENCES nucleo(id)")

	db.Exec("CREATE TABLE IF NOT EXISTS tipo_de_registro (id INTEGER NOT NULL PRIMARY KEY," +
		"nome TEXT NOT NULL, " +
		"nucleo INTEGER NOT NULL, " +
		"descricao TEXT NOT NULL, " + 
		"partilha_partidaria FLOAT NOT NULL)" + 
		"FOREIGN KEY (nucleo) REFERENCES nucleo(id)")

	db.Exec("CREATE TABLE IF NOT EXISTS registro (id INTEGER NOT NULL PRIMARY KEY," +
		"dia TEXT NOT NULL, " +
		"mes TEXT NOT NULL, " +
		"ano TEXT NOT NULL, " +
		"tipo INTEGER NOT NULL, " +
		"nucleo INTEGER NOT NULL, " +
		"doador TEXT NOT NULL, " +
		"receptor TEXT NOT NULL, " +
		"quantidade INTEGER NOT NULL, " +
		"valor FLOAT NOT NULL, " +
		"descricao TEXT NOT NULL, " +
		"FOREIGN KEY (nucleo) REFERENCES nucleo(id), " +
		"FOREIGN KEY (doador) REFERENCES pessoas(id), " +
		"FOREIGN KEY (receptor) REFERENCES pessoas(id), " + 
		"FOREIGN KEY (tipo) REFERENCES tipo_de_registro(id))")

	db.Exec("CREATE TABLE IF NOT EXISTS relatorio (id INTEGER NOT NULL PRIMARY KEY," +
		"mes TEXT NOT NULL, " +
		"ano TEXT NOT NULL, " +
	  "nucleo INTEGER NOT NULL, " +
		"membros TEXT NOT NULL, " +
		"pagamentos_membros TEXT NOT NULL, " +
		"despesas TEXT NOT NULL, " +
		"vendas TEXT NOT NULL, " +
		"total_ganho FLOAT NOT NULL, " +
		"total_liquido FLOAT NOT NULL, " +
		"dividendos_partido FLOAT NOT NULL, " +
		"superavit_nucleo FLOAT NOT NULL, " +
		"path_to_pdf TEXT NOT NULL, " +
		"FOREIGN KEY (membros) REFERENCES pessoas(id), " +
		"FOREIGN KEY (pagamentos_membros) REFERENCES registro(id), " +
		"FOREIGN KEY (despesas) REFERENCES registro(id), " +
		"FOREIGN KEY (vendas) REFERENCES registro(id)" + 
		"FOREIGN KEY (nucleo) REFERENCES nucleo(id))")

	db.Exec("CREATE TABLE IF NOT EXISTS dia_de_pagamento (dia_de_pagamento TEXT NOT NULL)")
	db.Exec("INSERT OR IGNORE INTO dia_de_pagamento (dia_de_pagamento) VALUES (10)")
	db.Exec("INSERT OR IGNORE INTO pessoas (id, nome, papel, contribuicao_mensal, credito) VALUES ('nucleo', 'nucleo', 'nucleo', 0, 0)")
	db.Exec("INSERT OR IGNORE INTO pessoas (id, nome, papel, contribuicao_mensal, credito) VALUES ('partido', 'partido', 'partido', 0, 0)")
	db.Exec("INSERT OR IGNORE INTO pessoas (id, nome, papel, contribuicao_mensal, credito) VALUES ('externo', 'externo', 'externo', 0, 0)")
}
