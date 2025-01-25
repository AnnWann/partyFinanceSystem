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

func (db *RegisterDB) GetNextId() (int, error) {
	var id int
	err := db.QueryRow("SELECT MAX(id) FROM registers").Scan(&id)
	if err != nil {
		return 0, err
	}
	return id + 1, nil
}

func (db *RegisterDB) InsertRegister(r models.Register) error {
	_, err := db.Exec("INSERT INTO registers (id, day, month, year, type, giver, receiver, ammount, value, description) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		r.Id, r.Day, r.Month, r.Year, r.Type, r.Giver, r.Receiver, r.Amount, r.Value, r.Description)

	return err
}

func (db *RegisterDB) GetRegister() ([]models.Register, error) {
	rows, err := db.Query("SELECT * FROM registers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var r []models.Register
	for rows.Next() {
		var register models.Register
		err := rows.Scan(&register.Id, &register.Day, &register.Month, &register.Year, &register.Type, &register.Giver, &register.Receiver, &register.Amount, &register.Value, &register.Description)
		if err != nil {
			return nil, err
		}
		r = append(r, register)
	}

	return r, err
}

func (db *RegisterDB) GetRegisterById(id string) (models.Register, error) {
	var r models.Register
	err := db.QueryRow("SELECT * FROM registers WHERE id = ?", id).Scan(&r.Id, &r.Day, &r.Month, &r.Year, &r.Type, &r.Giver, &r.Receiver, &r.Amount, &r.Value, &r.Description)
	if err != nil {
		return models.Register{}, err
	}

	return r, nil
}

func (db *RegisterDB) GetRegisterByMonthAndYear(month string, year string) ([]models.Register, error) {
	rows, err := db.Query("SELECT * FROM registers WHERE month = ? AND year = ?", month, year)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var registers []models.Register
	for rows.Next() {
		var r models.Register
		err := rows.Scan(&r.Id, &r.Day, &r.Month, &r.Year, &r.Type, &r.Giver, &r.Receiver, &r.Amount, &r.Value, &r.Description)
		if err != nil {
			return nil, err
		}
		registers = append(registers, r)
	}
	return registers, nil
}

func (db *RegisterDB) GetRegistersByYear(year string) ([]models.Register, error) {
	rows, err := db.Query("SELECT * FROM registers WHERE year = ?", year)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var registers []models.Register
	for rows.Next() {
		var r models.Register
		err := rows.Scan(&r.Id, &r.Day, &r.Month, &r.Year, &r.Type, &r.Giver, &r.Receiver, &r.Amount, &r.Value, &r.Description)
		if err != nil {
			return nil, err
		}
		registers = append(registers, r)
	}
	return registers, nil
}
