package database

import (
	"database/sql"

	"github.com/AnnWann/pstu_finance_system/lib/models"
)

type RegisterDB struct {
	*sql.DB
}

func (db *DBWrapper) GetRegisterDB() *RegisterDB {
	return &RegisterDB{db.DB}
}

func (db RegisterDB) InsertRegister(r models.Register) error {
	_, err := db.Exec("INSERT INTO registers (id, day, month, year, type, giver, receiver, ammount, value, description) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		r.Id, r.Day, r.Month, r.Year, r.Type, r.Giver, r.Receiver, r.Amount, r.Value, r.Description)

	return err
}

func (db RegisterDB) GetRegister(id string) (models.Register, error) {
	var r models.Register
	err := db.QueryRow("SELECT * FROM registers WHERE id = ?", id).Scan(
		&r.Id, &r.Day, &r.Month, &r.Year, &r.Type, &r.Giver, &r.Receiver, &r.Amount, &r.Value, &r.Description)

	return r, err
}

func (db RegisterDB) GetRegisters() ([]models.Register, error) {
	rows, err := db.Query("SELECT * FROM registers")
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

func (db RegisterDB) GetRegisterByMonthAndYear(month string, year string) ([]models.Register, error) {
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

func (db RegisterDB) GetRegistersByYear(year string) ([]models.Register, error) {
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

