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

func (db *NucleoDB) GetNextId() (int, error) {
	var id int
	err := db.QueryRow("SELECT MAX(id) FROM nucleo").Scan(&id)
	if err != nil {
		return 0, err
	}
	return id + 1, nil

}

func (db *NucleoDB) InsertNucleo(nucleo models.Nucleo) error {
	_, err := db.Exec("INSERT INTO nucleo (id, nucleo, cidade, estado, credito, dia_de_pagamento) VALUES (?, ?, ?, ?, ?, ?)", nucleo.Name, nucleo.City, nucleo.State, nucleo.Credit, nucleo.Payday)
	return err
}

func (db *NucleoDB) GetNucleo() ([]models.Nucleo, error) {
	query := "SELECT * FROM nucleo"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	nucleos := []models.Nucleo{}
	for rows.Next() {
		var nucleo models.Nucleo
		err = rows.Scan(&nucleo.Id, &nucleo.Name, &nucleo.City, &nucleo.State, &nucleo.Credit)
		if err != nil {
			return nil, err
		}
		nucleos = append(nucleos, nucleo)
	}

	return nucleos, nil
}

func (db *NucleoDB) GetNucleoById(id int) (models.Nucleo, error) {
	query := "SELECT * FROM nucleo WHERE id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		return models.Nucleo{}, err
	}
	defer rows.Close()

	var nucleo models.Nucleo
	for rows.Next() {
		err = rows.Scan(&nucleo.Id, &nucleo.Name, &nucleo.City, &nucleo.State, &nucleo.Credit)
		if err != nil {
			return models.Nucleo{}, err
		}
	}

	return nucleo, nil
}

func (db *NucleoDB) Counts() (int, error) {
	query := "SELECT COUNT(*) FROM nucleo"
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
	err := db.QueryRow("SELECT COUNT(*) FROM nucleo WHERE id = ?", id).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (db *NucleoDB) UpdateCredit(id int, credit float64) error {
	_, err := db.Exec("UPDATE nucleo SET credito = ? WHERE id = ?", credit, id)
	return err
}

func (db *NucleoDB) UpdatePayday(id int, dia string) error {
	_, err := db.Exec("UPDATE nucleo SET dia_de_pagamento = ? WHERE id = ?", dia, id)
	return err
}


func (db *NucleoDB) DeleteNucleo(id int) error {
	_, err := db.Exec("DELETE FROM nucleo WHERE id = ?", id)
	return err
}
