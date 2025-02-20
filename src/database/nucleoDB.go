package database

import (
	"database/sql"

	"github.com/AnnWann/pstu_finance_system/src/models"
)

type NucleoDB struct {
	*sql.DB
}

func (db *DBWrapper) GetNucleoDB() *NucleoDB {
	return &NucleoDB{db.DB}
}

func (db *NucleoDB) InsertNucleo(nucleo models.Nucleo) (int, error) {
	if db.NucleoExistsByName(nucleo.Nome) {
		return -1, nil
	}
	id, err := GetDB().GetPessoasDB().InsertPessoa("nucleo")
	if err != nil {
		return 0, err
	}
	_, err = db.Exec("INSERT INTO nucleos (id, nome, cidade, estado, reserva, dia_de_pagamento) VALUES (?, ?, ?, ?, ?, ?)", id, nucleo.Nome, nucleo.Cidade, nucleo.Estado, nucleo.Reserva, nucleo.Dia_de_Pagamento)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db *NucleoDB) GetNucleo() ([]models.Nucleo, error) {
	query := "SELECT * FROM nucleos"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	nucleos := []models.Nucleo{}
	for rows.Next() {
		var nucleo models.Nucleo
		err = rows.Scan(&nucleo.ID, &nucleo.Nome, &nucleo.Cidade, &nucleo.Estado, &nucleo.Reserva, &nucleo.Dia_de_Pagamento)
		if err != nil {
			return nil, err
		}
		nucleos = append(nucleos, nucleo)
	}

	return nucleos, nil
}

func (db *NucleoDB) GetNucleoById(id int) (models.Nucleo, error) {
	query := "SELECT * FROM nucleos WHERE id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		return models.Nucleo{}, err
	}
	defer rows.Close()

	var nucleo models.Nucleo
	for rows.Next() {
		err = rows.Scan(&nucleo.ID, &nucleo.Nome, &nucleo.Cidade, &nucleo.Estado, &nucleo.Reserva, &nucleo.Dia_de_Pagamento)
		if err != nil {
			return models.Nucleo{}, err
		}
	}

	return nucleo, nil
}

func (db *NucleoDB) Counts() (int, error) {
	query := "SELECT COUNT(*) FROM nucleos"
	rows, err := db.DB.Query(query)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}

func (db *NucleoDB) NucleoExists(id int) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM nucleos WHERE id = ?", id).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (db *NucleoDB) NucleoExistsByName(nome string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM nucleos WHERE nome = ?", nome).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (db *NucleoDB) UpdateReserva(id int, reserva float64) error {
	_, err := db.Exec("UPDATE nucleos SET reserva = ? WHERE id = ?", reserva, id)
	return err
}

func (db *NucleoDB) UpdateDiaDePagamento(id int, dia string) error {
	_, err := db.Exec("UPDATE nucleos SET dia_de_pagamento = ? WHERE id = ?", dia, id)
	return err
}

func (db *NucleoDB) DeleteNucleo(id int) error {
	_, err := db.Exec("DELETE FROM nucleos WHERE id = ?", id)
	return err
}
