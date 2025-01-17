package database

import (
	"github.com/AnnWann/pstu_finance_system/lib/models"
)

func InsertRegister(r models.Register) error {
	_, err := DB.Exec("INSERT INTO registers (id, day, month, year, type, giver, receiver, ammount, value, description) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		r.Id, r.Day, r.Month, r.Year, r.Type, r.Giver, r.Receiver, r.Amount, r.Value, r.Description)

	return err
}

func GetRegister(id string) (models.Register, error) {
	var r models.Register
	err := DB.QueryRow("SELECT * FROM registers WHERE id = ?", id).Scan(
		&r.Id, &r.Day, &r.Month, &r.Year, &r.Type, &r.Giver, &r.Receiver, &r.Amount, &r.Value, &r.Description)

	return r, err
}

func GetRegisters() ([]models.Register, error) {
	rows, err := DB.Query("SELECT * FROM registers")
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

func GetRegisterByMonthAndYear(month string, year string) ([]models.Register, error) {
	rows, err := DB.Query("SELECT * FROM registers WHERE month = ? AND year = ?", month, year)
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

func GetRegistersByYear(year string) ([]models.Register, error) {
	rows, err := DB.Query("SELECT * FROM registers WHERE year = ?", year)
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

