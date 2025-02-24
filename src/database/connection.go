package database

import (
	"database/sql"
	"os"
	"path/filepath"

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

func (db *DBWrapper) InitDB(file string) error {
	var err error
	if err = fileIsValid(file); err != nil {
		return err
	}

	db.DB, err = sql.Open("sqlite3", file)
	if err != nil {
		return err
	}

	PARTIDO := os.Getenv("PARTIDO")

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS pessoas (id INTEGER NOT NULL PRIMARY KEY,"+
		"classe TEXT NOT NULL);"+

		"CREATE TABLE IF NOT EXISTS partido (id INTEGER NOT NULL PRIMARY KEY,"+
		"nome TEXT NOT NULL, "+
		"reserva FLOAT NOT NULL, "+
		"FOREIGN KEY (id) REFERENCES pessoas(id));"+

		"CREATE TABLE IF NOT EXISTS nucleos (id INTEGER NOT NULL PRIMARY KEY,"+
		"nome TEXT NOT NULL, "+
		"cidade TEXT NOT NULL, "+
		"estado TEXT NOT NULL, "+
		"reserva FLOAT NOT NULL,"+
		"dia_de_pagamento TEXT NOT NULL,"+
		"administrador INTEGER NOT NULL,"+
		"FOREIGN KEY (administrador) REFERENCES pessoas(id),"+
		"FOREIGN KEY (id) REFERENCES pessoas(id));"+

		"CREATE TABLE IF NOT EXISTS cargos (id INTEGER NOT NULL PRIMARY KEY,"+
		"titulo TEXT NOT NULL, "+
		"nucleo INTEGER NOT NULL, "+
		"descricao TEXT NOT NULL, "+
		"FOREIGN KEY (nucleo) REFERENCES nucleos(id));"+

		"CREATE TABLE IF NOT EXISTS membros (id TEXT NOT NULL PRIMARY KEY,"+
		"nome TEXT NOT NULL, "+
		"nucleo INTEGER NOT NULL, "+
		"cargo INTEGER NOT NULL, "+
		"contribuicao_mensal FLOAT NOT NULL, "+
		"credito FLOAT NOT NULL, "+
		"FOREIGN KEY (nucleo) REFERENCES nucleos(id), "+
		"FOREIGN KEY (cargo) REFERENCES cargos(id), "+
		"FOREIGN KEY (id) REFERENCES pessoas(id));"+

		"CREATE TABLE IF NOT EXISTS tipos_de_registro (id INTEGER NOT NULL PRIMARY KEY,"+
		"nome TEXT NOT NULL, "+
		"nucleo INTEGER NOT NULL, "+
		"descricao TEXT NOT NULL, "+
		"partilha_partidaria FLOAT NOT NULL,"+
		"FOREIGN KEY (nucleo) REFERENCES nucleos(id)); "+

		"CREATE TABLE IF NOT EXISTS registros (id INTEGER NOT NULL PRIMARY KEY,"+
		"dia TEXT NOT NULL, "+
		"mes TEXT NOT NULL, "+
		"ano TEXT NOT NULL, "+
		"tipo INTEGER NOT NULL, "+
		"nucleo INTEGER NOT NULL, "+
		"pagante TEXT NOT NULL, "+
		"cobrante TEXT NOT NULL, "+
		"quantidade INTEGER NOT NULL, "+
		"valor FLOAT NOT NULL, "+
		"descricao TEXT NOT NULL, "+
		"FOREIGN KEY (nucleo) REFERENCES nucleos(id), "+
		"FOREIGN KEY (pagante) REFERENCES pessoas(id), "+
		"FOREIGN KEY (cobrante) REFERENCES pessoas(id), "+
		"FOREIGN KEY (tipo) REFERENCES tipo_de_registro(id)); "+

		"CREATE TABLE IF NOT EXISTS relatorios_mensais (id INTEGER NOT NULL PRIMARY KEY,"+
		"mes TEXT NOT NULL, "+
		"ano TEXT NOT NULL, "+
		"nucleo INTEGER NOT NULL, "+
		"membros TEXT NOT NULL, "+
		"registros TEXT NOT NULL, "+
		"total_ganho FLOAT NOT NULL, "+
		"total_liquido FLOAT NOT NULL, "+
		"pagamento_partido FLOAT NOT NULL, "+
		"lucro_nucleo FLOAT NOT NULL, "+
		"link_arquivo TEXT NOT NULL, "+
		"FOREIGN KEY (membros) REFERENCES pessoas(id), "+
		"FOREIGN KEY (registros) REFERENCES registro(id), "+
		"FOREIGN KEY (nucleo) REFERENCES nucleos(id)); "+

		"INSERT OR IGNORE INTO pessoas (id, classe) VALUES (0, 'externo');"+
		"INSERT OR IGNORE INTO pessoas (id, classe) VALUES (1, 'partido');"+
		"INSERT OR IGNORE INTO pessoas (id, classe) VALUES (2, 'nucleo_geral');"+
		"INSERT OR IGNORE INTO tipos_de_registro (id, nome, nucleo, descricao, partilha_partidaria) VALUES (-100, 'contribuicao', 2, 'Contribuição mensal do militante', 0.0); "+
		"INSERT OR IGNORE INTO tipos_de_registro (id, nome, nucleo, descricao, partilha_partidaria) VALUES (-200, 'jornal', 2, 'Venda de produtos', 2.0); "+
		"INSERT OR IGNORE INTO tipos_de_registro (id, nome, nucleo, descricao, partilha_partidaria) VALUES (-300, 'pagamento_partido', 2, 'pagamento do nucleo ao seu administrador', 0.0); "+
		"INSERT OR IGNORE INTO tipos_de_registro (id, nome, nucleo, descricao, partilha_partidaria) VALUES (-400, 'gasto', 2, 'Gasto do nucleo', 0.0); "+
		"INSERT OR IGNORE INTO tipos_de_registro (id, nome, nucleo, descricao, partilha_partidaria) VALUES (-500, 'abatimento', 2, 'Abatimento do nucleo', 0.0); "+
		"INSERT OR IGNORE INTO cargos (id, titulo, nucleo, descricao) VALUES (-100, 'aspirante', 2, 'Membro em treinamento'); "+
		"INSERT OR IGNORE INTO cargos (id, titulo, nucleo, descricao) VALUES (-200, 'militante', 2, 'Membro efetivo'); "+
		"INSERT OR IGNORE INTO cargos (id, titulo, nucleo, descricao) VALUES (-300, 'dirigente', 2, 'Membro com responsabilidades de liderança'); "+
		"INSERT OR IGNORE INTO cargos (id, titulo, nucleo, descricao) VALUES (-400, 'dirigente_financeiro', 2, 'Membro com responsabilidades de liderança financeira'); "+
		"INSERT OR IGNORE INTO partido (id, nome, reserva) VALUES (1, ?, 0);", PARTIDO)
	return err
}

func fileIsValid(file string) error {
	dir := filepath.Dir(file)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(file); os.IsNotExist(err) {
		_, err = os.Create(file)
		if err != nil {
			return err
		}
	}
	return nil
}
