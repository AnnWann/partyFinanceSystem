package database

import (
	"database/sql"

	"github.com/AnnWann/pstu_finance_system/src/models"
)

type TiposDeRegistroDB struct {
	*sql.DB
}

func (db *DBWrapper) GetTiposDeRegistroDB() *TiposDeRegistroDB {
	return &TiposDeRegistroDB{db.DB}
}

func (db *TiposDeRegistroDB) GetNextId() (int, error) {
	var id int
	err := db.QueryRow("SELECT MAX(id) FROM tipos_de_registro").Scan(&id)
	if err != nil {
		return 0, err
	}

	if id < 1 {
		return 1, nil
	}

	return id + 1, nil
}

func (db *TiposDeRegistroDB) InsertTipo(t models.Tipo_de_registro) error {
	_, err := db.Exec("INSERT INTO tipos_de_registro (id, nome, nucleo, descricao, partilha_partidaria) VALUES (?, ?, ?, ?, ?)", t.ID, t.Nome, t.Nucleo, t.Descricao, t.Parcela_partidaria)

	return err
}

func (db *TiposDeRegistroDB) GetTipo(id int) (models.Tipo_de_registro, error) {
	var t models.Tipo_de_registro
	err := db.QueryRow("SELECT * FROM tipos_de_registro WHERE id = ?", id).Scan(&t.ID, &t.Nome, &t.Nucleo, &t.Descricao, &t.Parcela_partidaria)

	return t, err
}

func (db *TiposDeRegistroDB) GetTipoPorNucleo(nucleoId int) ([]models.Tipo_de_registro, error) {
	rows, err := db.Query("SELECT * FROM tipos_de_registro WHERE nucleo = ?", nucleoId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []models.Tipo_de_registro
	for rows.Next() {
		var t models.Tipo_de_registro
		err := rows.Scan(&t.ID, &t.Nome, &t.Nucleo, &t.Descricao, &t.Parcela_partidaria)
		if err != nil {
			return nil, err
		}
		types = append(types, t)
	}

	generalTypes, err := db.GetTiposGeral()
	if err != nil {
		return nil, err
	}

	pagamentoPartidoId := db.GetPagamentoAdministrador()
	for _, t := range generalTypes {
		if t.ID == pagamentoPartidoId {
			continue
		}
		types = append(types, t)
	}

	return types, nil
}

func (db *TiposDeRegistroDB) GetTipos() ([]models.Tipo_de_registro, error) {
	rows, err := db.Query("SELECT * FROM tipos_de_registro")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []models.Tipo_de_registro
	for rows.Next() {
		var t models.Tipo_de_registro
		err := rows.Scan(&t.ID, &t.Nome, &t.Nucleo, &t.Descricao, &t.Parcela_partidaria)
		if err != nil {
			return nil, err
		}
		types = append(types, t)
	}
	return types, nil
}
func (db *TiposDeRegistroDB) GetTiposGeral() ([]models.Tipo_de_registro, error) {
	rows, err := db.Query("SELECT * FROM tipos_de_registro WHERE nucleo = 2")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []models.Tipo_de_registro
	for rows.Next() {
		var t models.Tipo_de_registro
		err := rows.Scan(&t.ID, &t.Nome, &t.Nucleo, &t.Descricao, &t.Parcela_partidaria)
		if err != nil {
			return nil, err
		}
		types = append(types, t)
	}
	return types, nil
}

func (db *TiposDeRegistroDB) GetContribuicaoId() int {
	return -100
}

func (db *TiposDeRegistroDB) GetJornalId() int {
	return -200
}

func (db *TiposDeRegistroDB) GetPagamentoAdministrador() int {
	return -300
}

func (db *TiposDeRegistroDB) GetGastosId() int {
	return -400
}

func (db *TiposDeRegistroDB) GetAbatimentosId() int {
	return -500
}

func (db *TiposDeRegistroDB) CountTipos() (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM tipos_de_registro").Scan(&count)
	return count, err
}

func (db *TiposDeRegistroDB) DeleteTipo(id int) error {
	_, err := db.Exec("DELETE FROM tipos_de_registro WHERE id = ?", id)
	return err
}

func (db *TiposDeRegistroDB) UpdatePartilhaPartidaria(id int, partilha float64) error {
	_, err := db.Exec("UPDATE tipos_de_registro SET partilha_partidaria = ? WHERE id = ?", partilha, id)
	return err
}
