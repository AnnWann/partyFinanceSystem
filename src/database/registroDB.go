package database

import (
	"database/sql"

	"github.com/AnnWann/pstu_finance_system/src/models"
)

type RegisterDB struct {
	*sql.DB
}

func (db *DBWrapper) GetRegisterDB() *RegisterDB {
	return &RegisterDB{db.DB}
}

func (db *RegisterDB) GetNextId() int {
	var id int
	err := db.QueryRow("SELECT MAX(id) FROM registros").Scan(&id)
	if err != nil || id < 1 {
		return 1
	}
	return id + 1
}

func (db *RegisterDB) InsertRegister(r models.Registro) error {
	_, err := db.Exec("INSERT INTO registros (id, dia, mes, ano, tipo, nucleo, pagante, cobrante, quantidade, valor, descricao) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		r.ID, r.Dia, r.Mes, r.Ano, r.Tipo, r.Nucleo, r.Pagante, r.Cobrante, r.Quantidade, r.Valor, r.Descricao)

	return err
}

func (db *RegisterDB) GetRegister() ([]models.Registro, error) {
	rows, err := db.Query("SELECT * FROM registros")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var r []models.Registro
	for rows.Next() {
		var register models.Registro
		err := rows.Scan(&register.ID, &register.Dia, &register.Mes, &register.Ano, &register.Tipo, &register.Nucleo, &register.Pagante, &register.Cobrante, &register.Quantidade, &register.Valor, &register.Descricao)
		if err != nil {
			return nil, err
		}
		r = append(r, register)
	}

	return r, err
}

func (db *RegisterDB) GetRegisterById(id int) (models.Registro, error) {
	var r models.Registro
	err := db.QueryRow("SELECT * FROM registros WHERE id = ?", id).Scan(&r.ID, &r.Dia, &r.Mes, &r.Ano, &r.Tipo, &r.Nucleo, &r.Pagante, &r.Cobrante, &r.Quantidade, &r.Valor, &r.Descricao)
	if err != nil {
		return models.Registro{}, err
	}

	return r, nil
}

func (db *RegisterDB) GetRegisterByMonthAndYear(mes string, ano string) ([]models.Registro, error) {
	rows, err := db.Query("SELECT * FROM registros WHERE mes = ? AND ano = ?", mes, ano)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var registers []models.Registro
	for rows.Next() {
		var r models.Registro
		err := rows.Scan(&r.ID, &r.Dia, &r.Mes, &r.Ano, &r.Tipo, &r.Nucleo, &r.Pagante, &r.Cobrante, &r.Quantidade, &r.Valor, &r.Descricao)
		if err != nil {
			return nil, err
		}
		registers = append(registers, r)
	}
	return registers, nil
}

func (db *RegisterDB) GetRegistersByYear(ano string) ([]models.Registro, error) {
	rows, err := db.Query("SELECT * FROM registros WHERE ano = ?", ano)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var registers []models.Registro
	for rows.Next() {
		var r models.Registro
		err := rows.Scan(&r.ID, &r.Dia, &r.Mes, &r.Ano, &r.Tipo, &r.Pagante, &r.Cobrante, &r.Quantidade, &r.Valor, &r.Descricao)
		if err != nil {
			return nil, err
		}
		registers = append(registers, r)
	}
	return registers, nil
}
