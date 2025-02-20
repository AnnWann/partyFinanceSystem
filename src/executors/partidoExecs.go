package executors

import (
	"errors"

	"github.com/AnnWann/pstu_finance_system/src/database"
	"github.com/AnnWann/pstu_finance_system/src/models"
)

func GetPartido() (models.Partido, error) {
	db := database.GetDB().GetPartidoDB()

	partido, err := db.GetPartido()
	if err != nil {
		return models.Partido{}, errors.New("erro ao obter o partido")
	}

	return partido, nil
}
