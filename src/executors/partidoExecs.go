package executors

import (
	"errors"

	"github.com/AnnWann/pstu_finance_system/src/database"
	"github.com/AnnWann/pstu_finance_system/src/models"
)

func AddPartido(nome string) error {
	db := database.GetDB().GetPartidoDB()

	if db.AlreadyExists() {
		return errors.New("Já existe um partido cadastrado")
	}

	id := 1

	partido := models.Partido{
		Id:     id,
		Name:   nome,
		Credit: 0,
	}

	err := db.InsertPartido(partido)
	if err != nil {
		return errors.New("Erro ao inserir o partido")
	}

	return nil
}

func GetPartido() (models.Partido, error) {
	db := database.GetDB().GetPartidoDB()

	partido, err := db.GetPartido()
	if err != nil {
		return models.Partido{}, errors.New("Erro ao obter o partido")
	}

	return partido, nil
}

