package database

import (
	"database/sql"

	"github.com/AnnWann/pstu_finance_system/src/models"
)

type MembroDB struct {
	*sql.DB
}

func (db *DBWrapper) GetMembroDB() *MembroDB {
	return &MembroDB{db.DB}
}

func (db *MembroDB) InsertMembro(p models.Membro) (int, error) {
	if db.MembroExistsByNome(p.Nome) {
		return -1, nil
	}

	id, err := GetDB().GetPessoasDB().InsertPessoa("membro")
	if err != nil {
		return -1, err
	}
	_, err = db.Exec("INSERT INTO membros (id, nome, nucleo, designacao, contribuicao_mensal, credito) VALUES (?, ?, ?, ?, ?, ?)", id, p.Nome, p.Nucleo, p.Designacao, p.Contribuicao_mensal, p.Credito)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (db *MembroDB) GetMembro() ([]models.Membro, error) {
	row, err := db.Query("SELECT * FROM membros")
	if err != nil {
		return []models.Membro{}, err
	}
	defer row.Close()

	var persons []models.Membro
	for row.Next() {
		var p models.Membro
		err = row.Scan(&p.ID, &p.Nome, &p.Nucleo, &p.Designacao, &p.Contribuicao_mensal, &p.Credito)
		if err != nil {
			return []models.Membro{}, err
		}
		persons = append(persons, p)
	}

	return persons, err
}

func (db *MembroDB) GetMembroById(id int) (models.Membro, error) {
	var p models.Membro
	err := db.QueryRow("SELECT * FROM membros WHERE id = ?", id).Scan(&p.ID, &p.Nome, &p.Nucleo, &p.Designacao, &p.Contribuicao_mensal, &p.Credito)
	if err != nil {
		return models.Membro{}, nil
	}
	return p, nil
}

func (db *MembroDB) GetMembroByNome(nome string) (models.Membro, error) {
	var p models.Membro
	err := db.QueryRow("SELECT * FROM membros WHERE nome = ?", nome).Scan(&p.ID, &p.Nome, &p.Nucleo, &p.Designacao, &p.Contribuicao_mensal, &p.Credito)

	return p, err
}

func (db *MembroDB) GetMembroByDesignacao(designacao string) (models.Membro, error) {
	var p models.Membro
	err := db.QueryRow("SELECT * FROM membros WHERE designacao = ?", designacao).Scan(&p.ID, &p.Nome, &p.Nucleo, &p.Designacao, &p.Contribuicao_mensal, &p.Credito)

	return p, err
}

func (db *MembroDB) GetMembroByNucleo(nucleo int) ([]models.Membro, error) {
	row, err := db.Query("SELECT * FROM membros WHERE nucleo = ?", nucleo)
	if err != nil {
		return []models.Membro{}, err
	}
	defer row.Close()

	var persons []models.Membro
	for row.Next() {
		var p models.Membro
		err = row.Scan(&p.ID, &p.Nome, &p.Nucleo, &p.Designacao, &p.Contribuicao_mensal, &p.Credito)
		if err != nil {
			return []models.Membro{}, err
		}
		persons = append(persons, p)
	}

	return persons, err
}

func (db *MembroDB) CountMembros() (int, error) {
	var size int
	err := db.QueryRow("SELECT COUNT(*) FROM membros").Scan(&size)

	return size, err
}

func (db *MembroDB) CountNucleoMembers(nucleiId int) (int, error) {
	var size int
	err := db.QueryRow("SELECT COUNT(*) FROM membros WHERE nucleo = ?", nucleiId).Scan(&size)

	return size, err
}

func (db *MembroDB) DeleteMembro(id int) error {
	_, err := db.Exec("DELETE FROM membros WHERE id = ?", id)

	return err
}

func (db *MembroDB) Promote(id int, designacao int) error {
	_, err := db.Exec("UPDATE membros SET designacao = ? WHERE id = ?", designacao, id)
	return err
}

func (db *MembroDB) UpdateContribuicaoMensal(id int, contribuicao_mensal float64) error {
	_, err := db.Exec("UPDATE membros SET contribuicao_mensal = ? WHERE id = ?", contribuicao_mensal, id)
	return err
}

func (db *MembroDB) UpdateCredito(id int, credito float64) error {
	_, err := db.Exec("UPDATE membros SET credito = ? WHERE id = ?", credito, id)
	return err
}

func (db *MembroDB) UpdateNucleo(id int, nucleo int) error {
	_, err := db.Exec("UPDATE membros SET nucleo = ? WHERE id = ?", nucleo, id)
	return err
}

func (db *MembroDB) MembroExistsByNome(nome string) bool {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM membros WHERE nome = ?)", nome).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (db *MembroDB) MembroExists(id int) bool {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM membros WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}
