package executors

import (
	"errors"
	"strconv"

	"github.com/AnnWann/pstu_finance_system/src/database"
	"github.com/AnnWann/pstu_finance_system/src/models"
)

func AddNucleo(nome, cidade, estado, payday string) (int, error) {
	db := database.GetDB().GetNucleoDB()

	nucleo := models.Nucleo{
		ID:               -1,
		Nome:             nome,
		Cidade:           cidade,
		Estado:           estado,
		Reserva:          0,
		Dia_de_Pagamento: payday,
	}

	id, err := db.InsertNucleo(nucleo)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetNucleo(modifiers map[string]string) ([]models.Nucleo, error) {
	db := database.GetDB().GetNucleoDB()

	nucleos, err := db.GetNucleo()
	if err != nil {
		return nil, errors.New("erro ao obter os nucleos")
	}

	if len(modifiers) > 0 {
		nucleos = filterNucleos(nucleos, modifiers)
	}
	
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

	pDb := DB.GetMembroDB()
	persons, err := pDb.GetMembroByNucleo(idInt)
	if err != nil {
		tx.Rollback()
		return errors.New("erro ao deletar o nucleo")
	}

	for _, person := range persons {
		err = pDb.DeleteMembro(person.ID)
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
			isValid = nucleo.ID == id
		case "--nome":
			isValid = nucleo.Nome == value
		case "--cidade":
			isValid = nucleo.Cidade == value
		case "--estado":
			isValid = nucleo.Estado == value
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
	err = db.UpdateDiaDePagamento(nucleoId, payday)
	if err != nil {
		return errors.New("erro ao atualizar o dia de pagamento")
	}

	return nil
}
