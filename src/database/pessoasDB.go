package database

import (
	"database/sql"

	"github.com/AnnWann/pstu_finance_system/src/models"
)

type PessoasDB struct {
	*sql.DB
}

func (db *DBWrapper) GetPessoasDB() *PessoasDB {
	return &PessoasDB{db.DB}
}

func (db *PessoasDB) getNextId() (int, error) {
	var id int
	err := db.QueryRow("SELECT MAX(id) FROM pessoas").Scan(&id)
	if err != nil {
		return 0, err
	}

	if id < 1 {
		return 1, nil
	}

	return id + 1, nil
}

func (db *PessoasDB) InsertPessoa(classe string) (int, error) {
	id, err := db.getNextId()
	if err != nil {
		return 0, err
	}

	_, err = db.Exec("INSERT INTO pessoas (id, classe) VALUES (?, ?)", id, classe)

	return id, err
}

func (db *PessoasDB) GetPessoa(id int) (models.Pessoa, error) {
	var p models.Pessoa
	err := db.QueryRow("SELECT * FROM pessoas WHERE id = ?", id).Scan(&p.ID, &p.Classe)

	return p, err
}

func (db *PessoasDB) GetPessoas() ([]models.Pessoa, error) {
	rows, err := db.Query("SELECT * FROM pessoas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var people []models.Pessoa
	for rows.Next() {
		var p models.Pessoa
		err := rows.Scan(&p.ID, &p.Classe)
		if err != nil {
			return nil, err
		}
		people = append(people, p)
	}

	return people, nil
}

func (db *PessoasDB) PessoaExists(id int) bool {
	var exists bool
	db.QueryRow("SELECT EXISTS(SELECT 1 FROM pessoas WHERE id = ?)", id).Scan(&exists)

	return exists
}

func (db *PessoasDB) GetExterno() int {
	return -100
}

func (db *PessoasDB) GetPartido() int {
	return -200
}

func (db *PessoasDB) GetNucleoGeral() int {
	return -300
}
