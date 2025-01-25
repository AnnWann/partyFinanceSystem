package database

import (
	"database/sql"

	"github.com/AnnWann/pstu_finance_system/src/models"
)

type TypeOfRegisterDB struct {
	*sql.DB
}

func (db *DBWrapper) GetTypeOfRegisterDB() *TypeOfRegisterDB {
	return &TypeOfRegisterDB{db.DB}
}

func (db *TypeOfRegisterDB) GetNextId() (int, error) {
	var id int
	err := db.QueryRow("SELECT MAX(id) FROM tipo_de_registro").Scan(&id)
	if err != nil {
		return 0, err
	}

	return id + 1, nil
}

func (db *TypeOfRegisterDB) InsertType(t models.TypeOfRegister) error {
	_, err := db.Exec("INSERT INTO tipo_de_registro (id, nome, descricao, partilha_partidaria) VALUES (?, ?, ?, ?)", t.Id, t.Name, t.Description, t.PartyShare)

	return err
}

func (db *TypeOfRegisterDB) GetType(id int) (models.TypeOfRegister, error) {
	var t models.TypeOfRegister
	err := db.QueryRow("SELECT * FROM tipo_de_registro WHERE id = ?", id).Scan(&t.Id, &t.Name, &t.Description, &t.PartyShare)

	return t, err
}

func (db *TypeOfRegisterDB) GetTypesByNucleo(nucleoId int) ([]models.TypeOfRegister, error) {
	rows, err := db.Query("SELECT * FROM tipo_de_registro WHERE nucleo = ?", nucleoId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []models.TypeOfRegister
	for rows.Next() {
		var t models.TypeOfRegister
		err := rows.Scan(&t.Id, &t.Name, &t.Description, &t.PartyShare)
		if err != nil {
			return nil, err
		}
		types = append(types, t)
	}

	return types, nil
}

func (db *TypeOfRegisterDB) GetTypes() ([]models.TypeOfRegister, error) {
	rows, err := db.Query("SELECT * FROM tipo_de_registro")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []models.TypeOfRegister
	for rows.Next() {
		var t models.TypeOfRegister
		err := rows.Scan(&t.Id, &t.Name, &t.Description, &t.PartyShare)
		if err != nil {
			return nil, err
		}
		types = append(types, t)
	}
	return types, nil
}

func (db *TypeOfRegisterDB) CountTypes() (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM tipo_de_registro").Scan(&count)
	return count, err
}

func (db *TypeOfRegisterDB) DeleteType(id string) error {
	_, err := db.Exec("DELETE FROM tipo_de_registro WHERE id = ?", id)
	return err
}
