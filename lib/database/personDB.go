package database

import (
	"github.com/AnnWann/pstu_finance_system/lib/models"
)

func InsertPerson(p models.Person) error {
	_, err := DB.Exec("INSERT INTO persons (name, role, credit) VALUES (?, ?, ?, ?, ?)", p.Id, p.Name, p.Role, p.MonthlyPayment, p.Credit)
	return err
}

func GetPerson(id string) (models.Person, error) {
	var p models.Person
	err := DB.QueryRow("SELECT * FROM persons WHERE id = ?", id).Scan(&p.Id, &p.Name, &p.Role, p.MonthlyPayment, &p.Credit)

	return p, err
}

func GetAllPersons() ([]models.Person, error) {
	var persons []models.Person
	rows, err := DB.Query("SELECT * FROM persons")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Person
		err = rows.Scan(&p.Id, &p.Name, &p.Role, p.MonthlyPayment, &p.Credit)
		if err != nil {
			return nil, err
		}
		persons = append(persons, p)
	}

	return persons, nil
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
		err = rows.Scan(&p.Id, &p.Name, &p.Role, p.MonthlyPayment, &p.Credit)
		if err != nil {
			return nil, err
		}
		members = append(members, p)
	}

	return members, nil
}

func GetCore() (models.Person, error) {
	var p models.Person
	err := DB.QueryRow("SELECT * FROM persons WHERE role = 'core'").Scan(&p.Id, &p.Name, &p.Role, p.MonthlyPayment, &p.Credit)

	return p, err
}

func GetParty() (models.Person, error) {
	var p models.Person
	err := DB.QueryRow("SELECT * FROM persons WHERE role = 'party'").Scan(&p.Id, &p.Name, &p.Role, p.MonthlyPayment, &p.Credit)

	return p, err
}

func GetPersonByName(name string) (models.Person, error) {
	var p models.Person
	err := DB.QueryRow("SELECT * FROM persons WHERE name = ?", name).Scan(&p.Id, &p.Name, &p.Role, p.MonthlyPayment, &p.Credit)

	return p, err
}

func GetPersonsByRole(role string) (models.Person, error) {
	var p models.Person
	err := DB.QueryRow("SELECT * FROM persons WHERE role = ?", role).Scan(&p.Id, &p.Name, &p.Role, p.MonthlyPayment, &p.Credit)

	return p, err
}

func CountPersons() (int, error) {
	var size int
	err := DB.QueryRow("SELECT COUNT(*) FROM persons").Scan(&size)

	return size, err
}

func CountMembers() (int, error) {
	var size int
	err := DB.QueryRow("SELECT COUNT(*) FROM persons WHERE role = 'militant' OR role = 'aspirant' OR 'leader' OR 'financeLeader'").Scan(&size)

	return size, err
}

func CountMilitants() (int, error) {
	var size int
	err := DB.QueryRow("SELECT COUNT(*) FROM persons WHERE role = 'militant'").Scan(&size)

	return size, err
}

func CountNonAspirantMembers() (int, error) {
	var size int
	err := DB.QueryRow("SELECT COUNT(*) FROM persons WHERE role = 'militant' OR role = 'leader' OR role = 'financeLeader'").Scan(&size)

	return size, err
}

func DeletePerson(id string) error {
	_, err := DB.Exec("DELETE FROM persons WHERE id = ?", id)

	return err
}

func UpdateRole(id string, role string) error {
	_, err := DB.Exec("UPDATE persons SET role = ? WHERE id = ?", role, id)
	return err
}

func PromoteToMilitant(id string) error {
	_, err := DB.Exec("UPDATE persons SET role = 'militant' WHERE id = ?", id)
	return err
}
func PromoteNewLeader(promotee string, demotee models.Person) error {
	tx, err := DB.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = DB.Exec("UPDATE persons SET role = 'militant' WHERE id = ?", demotee)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = DB.Exec("UPDATE persons SET role = ? WHERE id = ?", demotee.Role, promotee)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
func UpdateMonthlyContribution(id string, monthlyContribution int16) error {
	_, err := DB.Exec("UPDATE persons SET monthlyContribution = ? WHERE id = ?", monthlyContribution, id)
	return err
}

func UpdateCredit(id string, credit float32) error {
	_, err := DB.Exec("UPDATE persons SET credit = ? WHERE id = ?", credit, id)
	return err
}
