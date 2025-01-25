package executors

import (
	"crypto/sha256"
	"errors"
	"strconv"

	"github.com/AnnWann/pstu_finance_system/src/database"
	"github.com/AnnWann/pstu_finance_system/src/models"
)

func AddPerson(name string, nucleo string) (string, error) {
	nucleoInt, err := strconv.Atoi(nucleo)
	if err != nil {
		return "", errors.New("Núcleo deve ser um número")
	}

	nucleoExists := database.GetDB().GetNucleoDB().NucleoExists(nucleoInt)
	if !nucleoExists {
		return "", errors.New("Núcleo não existe")
	}

	personDB := database.GetDB().GetPersonDB()
	membersCount, err := personDB.CountMembers()
	if err != nil {
		return "", err
	}

	role := "aspirant"
	if membersCount < 1 {
		role = "leader"
	} else if membersCount < 2 {
		role = "financeLeader"
	}

	hash := sha256.New()
	hash.Write([]byte(name))
	person := models.Person{
		Id:             string(hash.Sum(nil)),
		Name:           name,
		Nucleo:         nucleoInt,
		Role:           role,
		Credit:         0,
		MonthlyPayment: 0,
	}

	err = personDB.InsertPerson(person)
	return person.Id, err
}

func GetPerson(filterOptions map[string]string) ([]models.Person, error) {

	persons, err := database.GetDB().GetPersonDB().GetPerson()
	if err != nil {
		return nil, err
	}

	return filterPersons(persons, filterOptions), nil
}

func Promote(id string, role string) error {
	db := database.GetDB().GetPersonDB()
	if role == "gerente" || role == "gerente_financeiro" {
		lPerson, err := db.GetPersonByRole(role)
		if err != nil {
			return err
		}
		tx, err := db.Begin()
		if err != nil {
			tx.Rollback()
			return err
		}
		err = db.Promote(lPerson.Id, "militante")
		if err != nil {
			tx.Rollback()
			return err
		}

		err = db.Promote(id, role)
		if err != nil {
			tx.Rollback()
			return err
		}

		tx.Commit()
		return nil
	}

	err := db.Promote(id, role)
	return err
}

func DeletePerson(id string) error {
	if id == "" {
		return errors.New("arguments cannot be empty. the correct format is 'delete <id>'")
	}
	err := database.GetDB().GetPersonDB().DeletePerson(id)
	return err
}

func filterPersons(persons []models.Person, filterOptions map[string]string) []models.Person {
	if filterOptions == nil {
		return persons
	}

	var filteredPersons []models.Person
	for _, person := range persons {
		if filterPerson(person, filterOptions) {
			filteredPersons = append(filteredPersons, person)
		}
	}

	return filteredPersons
}

func filterPerson(person models.Person, filterOptions map[string]string) bool {
	isValid := false
	for key, value := range filterOptions {
		switch key {
		case "--id":
			isValid = person.Id == value
		case "--name":
			isValid = person.Name == value
		case "--role":
			isValid = person.Role == value
		case "--nucleo":
			nucleo, err := strconv.Atoi(value)
			if err != nil {
				return false
			}
			isValid = person.Nucleo == nucleo
		}
	}
	return isValid
}
