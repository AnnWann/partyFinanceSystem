package database

import (
	"github.com/AnnWann/pstu_finance_system/lib/models"
)

func InsertPerson(p models.Person) error {
	_, err := DB.Exec("INSERT INTO persons (name, role, credit) VALUES (?, ?, ?, ?, ?)", p.Id, p.Name, p.Role, p.MonthlyContribution, p.Credit)
	return err
}

func GetPerson(id string) (models.Person, error) {
	var p models.Person
	err := DB.QueryRow("SELECT * FROM persons WHERE id = ?", id).Scan(&p.Id, &p.Name, &p.Role, p.MonthlyContribution, &p.Credit)

	return p, err
}

func GetMembers() ([]models.Person, error) {
	var members []models.Person
	rows, err := DB.Query("SELECT * FROM persons WHERE role = 'militant' OR role = 'aspirant'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Person
		err = rows.Scan(&p.Id, &p.Name, &p.Role, p.MonthlyContribution, &p.Credit)
		if err != nil {
			return nil, err
		}
		members = append(members, p)
	}

	return members, nil
}

func GetCore() (models.Person, error) {
	var p models.Person
	err := DB.QueryRow("SELECT * FROM persons WHERE role = 'core'").Scan(&p.Id, &p.Name, &p.Role, p.MonthlyContribution, &p.Credit)

	return p, err
}

func GetParty() (models.Person, error) {
	var p models.Person
	err := DB.QueryRow("SELECT * FROM persons WHERE role = 'party'").Scan(&p.Id, &p.Name, &p.Role, p.MonthlyContribution, &p.Credit)

	return p, err
}

func GetPersonByName(name string) (models.Person, error) {
	var p models.Person
	err := DB.QueryRow("SELECT * FROM persons WHERE name = ?", name).Scan(&p.Id, &p.Name, &p.Role, p.MonthlyContribution, &p.Credit)

	return p, err
}

func GetPersonsByRole(role string) (models.Person, error) {
	var p models.Person
	err := DB.QueryRow("SELECT * FROM persons WHERE role = ?", role).Scan(&p.Id, &p.Name, &p.Role, p.MonthlyContribution, &p.Credit)

	return p, err
}

func DeletePerson(id string) error {
	_, err := DB.Exec("DELETE FROM persons WHERE id = ?", id)

	return err
}

func UpdateMonthlyContribution(id string, monthlyContribution int16) error {
	_, err := DB.Exec("UPDATE persons SET monthlyContribution = ? WHERE id = ?", monthlyContribution, id)
	return err
}

func UpdateCredit(id string, credit float32) error {
	_, err := DB.Exec("UPDATE persons SET credit = ? WHERE id = ?", credit, id)
	return err
}
