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
	db.QueryRow("SELECT EXISTS(SELECT * FROM partido)").Scan(&exists)
	return exists
}

func (db *PartidoDB) GetPartido() (models.Partido, error) {
	var partido models.Partido
	err := db.QueryRow("SELECT * FROM partido").Scan(&partido.ID, &partido.Nome, &partido.Reserva)
	return partido, err
}

func (db *PartidoDB) GetPartidoId() int {
	return 1
}

func (db *PartidoDB) UpdateCredit(credit float64) error {
	_, err := db.Exec("UPDATE Partido SET credito = ?", credit)
	return err
}
