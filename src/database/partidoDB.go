package database

import (
	"database/sql"

	"github.com/AnnWann/pstu_finance_system/src/models"
)

type PartidoDB struct {
	*sql.DB
}

func (db *DBWrapper) GetPartidoDB() *PartidoDB {
	return &PartidoDB{db.DB}
}

func (db *PartidoDB) AlreadyExists() bool {
	var exists bool
	db.QueryRow("SELECT EXISTS(SELECT 1 FROM Partido)").Scan(&exists)
	return exists
}

func (db *PartidoDB) InsertPartido(partido models.Partido) error {
	_, err := db.Exec("INSERT INTO Partido (id, nome, credito) VALUES (?, ?, ?)", partido.Id, partido.Name, partido.Credit)
	return err
}

func (db *PartidoDB) GetPartido() (models.Partido, error) {
	var partido models.Partido
	err := db.QueryRow("SELECT * FROM Partido").Scan(&partido.Id, &partido.Name, &partido.Credit)
	return partido, err
}

func (db *PartidoDB) UpdateCredit(credit float64) error {
	_, err := db.Exec("UPDATE Partido SET credito = ?", credit)
	return err
}
