package executors

import (
	"errors"
	"strconv"

	"github.com/AnnWann/pstu_finance_system/src/database"
	"github.com/AnnWann/pstu_finance_system/src/models"
)

func AddNucleo(nome, cidade, estado, payday string) (int, error) {
	db := database.GetDB().GetNucleoDB()

	id, err := db.GetNextId()
	if err != nil {
		return 0, errors.New("erro ao obter o próximo id")
	}

	nucleo := models.Nucleo{
		Id:     id,
		Name:   nome,
		City:   cidade,
		State:  estado,
		Credit: 0,
		Payday: payday,
	}

	err = db.InsertNucleo(nucleo)
	if err != nil {
		return 0, errors.New("erro ao inserir o nucleo")
	}

	return id, nil
}

func GetNucleo(modifiers map[string]string) ([]models.Nucleo, error) {
	db := database.GetDB().GetNucleoDB()

	nucleos, err := db.GetNucleo()
	if err != nil {
		return nil, errors.New("erro ao obter os nucleos")
	}

	nucleos = filterNucleos(nucleos, modifiers)

	if len(nucleos) == 0 {
		return nil, errors.New("nenhum nucleo encontrado")
	}

	return nucleos, nil
}

func DeleteNucleo(id string) error {
	DB := database.GetDB()
	tx, err := DB.Begin()
	if err != nil {
		return errors.New("erro ao deletar o nucleo")
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		tx.Rollback()
		return errors.New("id inválido")
	}

	pDb := DB.GetPersonDB()
	persons, err := pDb.GetPersonByNucleo(idInt)
	if err != nil {
		tx.Rollback()
		return errors.New("erro ao deletar o nucleo")
	}

	for _, person := range persons {
		err = pDb.DeletePerson(person.Id)
		if err != nil {
			tx.Rollback()
			return errors.New("erro ao deletar o nucleo")
		}
	}

	nDb := database.GetDB().GetNucleoDB()

	err = nDb.DeleteNucleo(idInt)
	if err != nil {
		tx.Rollback()
		return errors.New("erro ao deletar o nucleo")
	}

	tx.Commit()
	return nil
}

func filterNucleos(nucleos []models.Nucleo, filterOptions map[string]string) []models.Nucleo {
	if filterOptions == nil {
		return nucleos
	}

	var filteredNucleos []models.Nucleo
	for _, nucleo := range nucleos {
		if filterNucleo(nucleo, filterOptions) {
			filteredNucleos = append(filteredNucleos, nucleo)
		}
	}

	return filteredNucleos
}

func filterNucleo(nucleo models.Nucleo, filterOptions map[string]string) bool {
	isValid := false
	for key, value := range filterOptions {
		switch key {
		case "--id":
			id, err := strconv.Atoi(value)
			if err != nil {
				return false
			}
			isValid = nucleo.Id == id
		case "--name":
			isValid = nucleo.Name == value
		case "--city":
			isValid = nucleo.City == value
		case "--state":
			isValid = nucleo.State == value
		}
	}
	return isValid
}

func UpdatePayday(id, payday string) error {
	nucleoId, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("id inválido")
	}

	db := database.GetDB().GetNucleoDB()
	err = db.UpdatePayday(nucleoId, payday)
	if err != nil {
		return errors.New("erro ao atualizar o dia de pagamento")
	}

	return nil
}
