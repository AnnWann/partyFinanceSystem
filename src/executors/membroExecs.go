package executors

import (
	"errors"
	"strconv"

	"github.com/AnnWann/pstu_finance_system/src/database"
	"github.com/AnnWann/pstu_finance_system/src/models"
)

func AddMembro(name string, nucleo string) (int, error) {
	nucleoInt, err := strconv.Atoi(nucleo)
	if err != nil {
		return 0, errors.New("núcleo deve ser um número")
	}

	nucleoExists := database.GetDB().GetNucleoDB().NucleoExists(nucleoInt)
	if !nucleoExists {
		return 0, errors.New("núcleo não existe")
	}

	personDB := database.GetDB().GetMembroDB()
	membersCount, err := personDB.CountNucleoMembers(nucleoInt)
	if err != nil {
		return 0, err
	}

	designacao := database.GetDB().GetCargoDB().GetAspiranteId() //aspirante
	if membersCount < 1 {
		designacao = database.DB.GetCargoDB().GetDirigenteId() //dirigente
	} else if membersCount < 2 {
		designacao = database.DB.GetCargoDB().GetDirigenteFinanceiroId() //dirigente_financeiro
	}

	person := models.Membro{
		ID:                  -1,
		Nome:                name,
		Nucleo:              nucleoInt,
		Cargo:               designacao,
		Credito:             0,
		Contribuicao_mensal: 0,
	}

	id, err := personDB.InsertMembro(person)
	return id, err
}

func GetMembro(filterOptions map[string]string) ([]models.Membro, error) {

	membros, err := database.GetDB().GetMembroDB().GetMembro()
	if err != nil {
		return nil, err
	}

	if len(filterOptions) > 0 {
		membros = filterMembros(membros, filterOptions)
	}

	if len(membros) == 0 {
		return nil, errors.New("nenhum membro encontrado")
	}

	return membros, nil
}

func Promote(id string, designacao string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("id deve ser um número")
	}

	designacaoInt, err := strconv.Atoi(designacao)
	if err != nil {
		return errors.New("designacao deve ser um número")
	}

	if !database.GetDB().GetCargoDB().CargoExists(designacaoInt) {
		return errors.New("designacao não existe")
	}

	if designacaoInt == database.GetDB().GetCargoDB().GetAspiranteId() {
		return errors.New("não é possível promover um membro para aspirante")
	}

	db := database.GetDB().GetMembroDB()
	if designacaoInt == database.GetDB().GetCargoDB().GetDirigenteId() || designacaoInt == database.GetDB().GetCargoDB().GetDirigenteFinanceiroId() {
		person, err := db.GetMembroById(idInt)
		if err != nil {
			return err
		}
		if person.Cargo == database.GetDB().GetCargoDB().GetAspiranteId() {
			return errors.New("não é possível promover um aspirante para posição de liderança")
		}
		lPerson, err := db.GetMembroByCargo(designacao)
		if err != nil {
			return err
		}
		if lPerson.ID == idInt {
			return errors.New("não é possível promover a pessoa para o mesmo cargo")
		}
		tx, err := db.Begin()
		if err != nil {
			tx.Rollback()
			return err
		}
		err = db.Promote(lPerson.ID, person.Cargo)
		if err != nil {
			tx.Rollback()
			return err
		}

		err = db.Promote(idInt, designacaoInt)
		if err != nil {
			tx.Rollback()
			return err
		}

		tx.Commit()
		return nil
	}

	return db.Promote(idInt, designacaoInt)
}

func UpdateMembro(id string, nucleo string, payment string) error {
	if nucleo == "" && payment == "" {
		return errors.New("nada para atualizar")
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("id deve ser um número")
	}

	if !database.GetDB().GetMembroDB().MembroExists(idInt) {
		return errors.New("membro não existe")
	}

	db := database.GetDB().GetMembroDB()
	if nucleo != "" {
		nucleoInt, err := strconv.Atoi(nucleo)
		if err != nil {
			return errors.New("núcleo deve ser um número")
		}
		nucleoExists := database.GetDB().GetNucleoDB().NucleoExists(nucleoInt)
		if !nucleoExists {
			return errors.New("núcleo não existe")
		}
		err = db.UpdateNucleo(idInt, nucleoInt)
		if err != nil {
			return err
		}
	}

	if payment != "" {
		paymentFloat, err := strconv.ParseFloat(payment, 64)
		if err != nil {
			return errors.New("pagamento deve ser um número")
		}
		err = db.UpdateContribuicaoMensal(idInt, paymentFloat)
		if err != nil {
			return err
		}
	}

	return nil

}

func DeleteMembro(id string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("id deve ser um número")
	}
	err = database.GetDB().GetMembroDB().DeleteMembro(idInt)
	return err
}

func filterMembros(membros []models.Membro, filterOptions map[string]string) []models.Membro {
	if filterOptions == nil {
		return membros
	}

	var filteredMembros []models.Membro
	for _, membro := range membros {
		if filterMembro(membro, filterOptions) {
			filteredMembros = append(filteredMembros, membro)
		}
	}

	return filteredMembros
}

func filterMembro(membro models.Membro, filterOptions map[string]string) bool {
	var allValidValues = []bool{} //This is dumb, but for some reason, there is a chance that filterOptions will be ordered differently each time
	for key, value := range filterOptions {
		isValid := false
		switch key {
		case "--id":
			valueInt, err := strconv.Atoi(value)
			if err != nil {
				isValid = false
			} else {
				isValid = membro.ID == valueInt
			}
		case "--nome":
			isValid = membro.Nome == value
		case "--cargo":
			role, err := strconv.Atoi(value)
			if err != nil {
				isValid = false
			} else {
				isValid = membro.Cargo == role
			}
		case "--nucleo":
			nucleo, err := strconv.Atoi(value)
			if err != nil {
				isValid = false
			} else {
				isValid = membro.Nucleo == nucleo
			}
		}
		allValidValues = append(allValidValues, isValid)
	}
	return AllTrue(allValidValues)
}
