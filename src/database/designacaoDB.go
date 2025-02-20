package database

import (
	"database/sql"

	"github.com/AnnWann/pstu_finance_system/src/models"
)

type DesignacaoDB struct {
	*sql.DB
}

func (db *DBWrapper) GetDesignacaoDB() *DesignacaoDB {
	return &DesignacaoDB{db.DB}
}

func (db *DesignacaoDB) GetNextId() (int, error) {
	var id int
	err := db.QueryRow("SELECT MAX(id) FROM designacao").Scan(&id)
	if err != nil {
		return 0, err
	}
	return id + 1, nil
}

func (db *DesignacaoDB) AddDesignacao(designacao models.Designacao) (int, error) {
	id, err := db.GetNextId()
	if err != nil {
		return -1, err
	}
	_, err = db.Exec("INSERT INTO designacao (id, titulo, descricao, nucleo) VALUES (?, ?, ?, ?)", id, designacao.Titulo, designacao.Descricao, designacao.Nucleo)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (db *DesignacaoDB) GetDesignacao(id int) (models.Designacao, error) {
	var designacao models.Designacao
	err := db.QueryRow("SELECT * FROM designacao WHERE id = ?", id).Scan(&designacao.ID, &designacao.Titulo, &designacao.Descricao, &designacao.Nucleo)
	if err != nil {
		return models.Designacao{}, err
	}
	return designacao, nil
}

func (db *DesignacaoDB) GetDesignacoesByTituloAndNucleo(titulo string, nucleo string) ([]models.Designacao, error) {
	rows, err := db.Query("SELECT * FROM designacao WHERE titulo = ? AND nucleo = ?", titulo, nucleo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	designacoes := []models.Designacao{}
	for rows.Next() {
		var designacao models.Designacao
		err := rows.Scan(&designacao.ID, &designacao.Titulo, &designacao.Descricao, &designacao.Nucleo)
		if err != nil {
			return nil, err
		}
		designacoes = append(designacoes, designacao)
	}

	return designacoes, nil
}

func (db *DesignacaoDB) GetDesignacoesGerais() ([]models.Designacao, error) {
	rows, err := db.Query("SELECT * FROM designacao WHERE nucleo = 2")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	designacoes := []models.Designacao{}
	for rows.Next() {
		var designacao models.Designacao
		err := rows.Scan(&designacao.ID, &designacao.Titulo, &designacao.Descricao, &designacao.Nucleo)
		if err != nil {
			return nil, err
		}
		designacoes = append(designacoes, designacao)
	}

	return designacoes, nil
}

func (db *DesignacaoDB) GetDesignacoes() ([]models.Designacao, error) {
	rows, err := db.Query("SELECT * FROM designacao")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	designacoes := []models.Designacao{}
	for rows.Next() {
		var designacao models.Designacao
		err := rows.Scan(&designacao.ID, &designacao.Titulo, &designacao.Descricao, &designacao.Nucleo)
		if err != nil {
			return nil, err
		}
		designacoes = append(designacoes, designacao)
	}

	return designacoes, nil
}

func (db *DesignacaoDB) DesignacaoExistsByTituloAndNucleo(titulo string, nucleo int) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM designacao WHERE titulo = ? AND nucleo = ?", titulo, nucleo).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (db *DesignacaoDB) DesignacaoExists(id int) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM designacao WHERE id = ?", id).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (DesignacaoDB) GetAspiranteId() int {
	return -100
}

func (DesignacaoDB) GetMilitanteId() int {
	return -200
}

func (DesignacaoDB) GetDirigenteId() int {
	return -300
}

func (DesignacaoDB) GetDirigenteFinanceiroId() int {
	return -400
}
