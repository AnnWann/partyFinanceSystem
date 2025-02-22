package database

import (
	"database/sql"

	"github.com/AnnWann/pstu_finance_system/src/models"
)

type CargoDB struct {
	*sql.DB
}

func (db *DBWrapper) GetCargoDB() *CargoDB {
	return &CargoDB{db.DB}
}

func (db *CargoDB) GetNextId() (int, error) {
	var id int
	err := db.QueryRow("SELECT MAX(id) FROM cargos").Scan(&id)
	if err != nil {
		return 0, err
	}
	return id + 1, nil
}

func (db *CargoDB) AddCargo(cargo models.Cargo) (int, error) {
	id, err := db.GetNextId()
	if err != nil {
		return -1, err
	}
	_, err = db.Exec("INSERT INTO cargos (id, titulo, descricao, nucleo) VALUES (?, ?, ?, ?)", id, cargo.Titulo, cargo.Descricao, cargo.Nucleo)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (db *CargoDB) GetCargo(id int) (models.Cargo, error) {
	var designacao models.Cargo
	err := db.QueryRow("SELECT * FROM cargos WHERE id = ?", id).Scan(&designacao.ID, &designacao.Titulo, &designacao.Descricao, &designacao.Nucleo)
	if err != nil {
		return models.Cargo{}, err
	}
	return designacao, nil
}

func (db *CargoDB) GetCargosByTituloAndNucleo(titulo string, nucleo string) ([]models.Cargo, error) {
	rows, err := db.Query("SELECT * FROM cargos WHERE titulo = ? AND nucleo = ?", titulo, nucleo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cargos := []models.Cargo{}
	for rows.Next() {
		var cargo models.Cargo
		err := rows.Scan(&cargo.ID, &cargo.Titulo, &cargo.Descricao, &cargo.Nucleo)
		if err != nil {
			return nil, err
		}
		cargos = append(cargos, cargo)
	}

	return cargos, nil
}

func (db *CargoDB) GetCargosGerais() ([]models.Cargo, error) {
	rows, err := db.Query("SELECT * FROM cargos WHERE nucleo = 2")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cargos := []models.Cargo{}
	for rows.Next() {
		var cargo models.Cargo
		err := rows.Scan(&cargo.ID, &cargo.Titulo, &cargo.Descricao, &cargo.Nucleo)
		if err != nil {
			return nil, err
		}
		cargos = append(cargos, cargo)
	}

	return cargos, nil
}

func (db *CargoDB) GetCargos() ([]models.Cargo, error) {
	rows, err := db.Query("SELECT * FROM cargos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cargos := []models.Cargo{}
	for rows.Next() {
		var cargo models.Cargo
		err := rows.Scan(&cargo.ID, &cargo.Titulo, &cargo.Descricao, &cargo.Nucleo)
		if err != nil {
			return nil, err
		}
		cargos = append(cargos, cargo)
	}

	return cargos, nil
}

func (db *CargoDB) CargoExistsByTituloAndNucleo(titulo string, nucleo int) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM cargos WHERE titulo = ? AND nucleo = ?", titulo, nucleo).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (db *CargoDB) CargoExists(id int) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM cargos WHERE id = ?", id).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (CargoDB) GetAspiranteId() int {
	return -100
}

func (CargoDB) GetMilitanteId() int {
	return -200
}

func (CargoDB) GetDirigenteId() int {
	return -300
}

func (CargoDB) GetDirigenteFinanceiroId() int {
	return -400
}
