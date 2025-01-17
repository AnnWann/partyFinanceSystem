package executors

import (
	"crypto/sha256"
	"errors"

	"github.com/AnnWann/pstu_finance_system/lib/database"
	"github.com/AnnWann/pstu_finance_system/lib/models"
)

func AddPerson(name string, role string) error {
	if name == "" || role == "" {
		err := errors.New("arguments cannot be empty. The correct format is 'add person <name> <role>'")
		return err
	}
	if role != "militant" && role != "aspirant" && role != "core" &&
		role != "party" && role != "outsider" &&
		role != "leader" && role != "financeLeader" {
		err := errors.New("can't add person with role " + role +
			". Role must be 'militant', 'aspirant', 'core', 'party', 'outsider', 'leader' or 'financeLeader'")
		return err
	}

	hash := sha256.New()
	hash.Write([]byte(name))

	person := models.Person{
		Id:             string(hash.Sum(nil)),
		Name:           name,
		Role:           role,
		Credit:         0,
		MonthlyPayment: 0,
	}

	if role == "core" || role == "party" || role == "outsider" || role == "leader" || role == "financeLeader" {
		err := errors.New("can't add person with role " + role)
		return err
	}

	db := database.GetDB().GetPersonDB()

	if role == "aspirant" {
		count, err := db.CountNonAspirantMembers()
		if err != nil {
			return err
		}
		if count < 2 {
			err := errors.New("can't add aspirant, there must be at least 2 members")
			return err
		}
	}

	membersCount, err := db.CountMembers()
	if err != nil {
		return err
	}

	if membersCount < 1 {
		person.Role = "leader"
	} else if membersCount < 2 {
		person.Role = "financeLeader"
	}

	err = db.InsertPerson(person)
	return err
}

func GetAllPersons() ([]models.Person, error) {
	persons, err := database.GetDB().GetPersonDB().GetAllPersons()
	return persons, err
}

func GetMembers() ([]models.Person, error) {
	members, err := database.GetDB().GetPersonDB().GetMembers()
	return members, err
}

func GetPerson(id string) (models.Person, error) {
	if id == "" {
		err := errors.New("arguments cannot be empty. the correct format is 'get person <id>'")
		return models.Person{}, err
	}

	person, err := database.GetDB().GetPersonDB().GetPerson(id)
	return person, err
}

func PromoteToMilitant(id string) error {
	if id == "" {
		return errors.New("arguments cannot be empty. the correct format is 'promote <id>'")
	}

	db := database.GetDB().GetPersonDB()

	p, err := db.GetPerson(id)
	if err != nil {
		return err
	}

	if p.Role != "aspirant" {
		err := errors.New("can't promote person with role " + p.Role)
		return err
	}

	err = db.PromoteToMilitant(id)
	return err
}

func PromoteNewLeader(promotee string, demotee string) error {
	if promotee == "" || demotee == "" {
		err := errors.New("arguments cannot be empty. the correct format is 'promote <promoteeId> <demoteeId>'")
		return err
	}

	db := database.GetDB().GetPersonDB()

	demoteeP, err := db.GetPerson(demotee)
	if err != nil {
		return err
	}

	if demoteeP.Role != "leader" && demoteeP.Role != "financeLeader" {
		err := errors.New("can't demote person with role " + demoteeP.Role)
		return err
	}

	err = db.PromoteNewLeader(promotee, demoteeP)
	return err
}

func DeletePerson(id string) error {
	if id == "" {
		return errors.New("arguments cannot be empty. the correct format is 'delete <id>'")
	}
	err := database.GetDB().GetPersonDB().DeletePerson(id)
	return err
}
