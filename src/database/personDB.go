package database

import (
	"database/sql"

	"github.com/AnnWann/pstu_finance_system/src/models"
)

type PersonDB struct {
	*sql.DB
}

func (db *DBWrapper) GetPersonDB() *PersonDB {
	return &PersonDB{db.DB}
}

func (db *PersonDB) InsertPerson(p models.Person) error {
	_, err := db.Exec("INSERT INTO persons (name, role, credit) VALUES (?, ?, ?, ?, ?)", p.Id, p.Name, p.Role, p.MonthlyPayment, p.Credit)
	return err
}

func (db *PersonDB) GetPerson() ([]models.Person, error) {
	row, err := db.Query("SELECT * FROM persons")
	if err != nil {
		return []models.Person{}, err
	}
	defer row.Close()

	var persons []models.Person
	for row.Next() {
		var p models.Person
		err = row.Scan(&p.Id, &p.Name, &p.Role, p.MonthlyPayment, &p.Credit)
		if err != nil {
			return []models.Person{}, err
		}
		persons = append(persons, p)
	}

	return persons, err
}

func (db *PersonDB) GetPersonById(id string) (models.Person, error) {
	var p models.Person
	err := db.QueryRow("SELECT * FROM persons WHERE id = ?", id).Scan(&p.Id, &p.Name, &p.Role, p.MonthlyPayment, &p.Credit)

	return p, err
}

func (db *PersonDB) GetPersonByName(name string) (models.Person, error) {
	var p models.Person
	err := db.QueryRow("SELECT * FROM persons WHERE name = ?", name).Scan(&p.Id, &p.Name, &p.Role, p.MonthlyPayment, &p.Credit)

	return p, err
}

func (db *PersonDB) GetPersonByRole(role string) (models.Person, error) {
	var p models.Person
	err := db.QueryRow("SELECT * FROM persons WHERE role = ?", role).Scan(&p.Id, &p.Name, &p.Role, p.MonthlyPayment, &p.Credit)

	return p, err
}

func (db *PersonDB) GetPersonByNucleo(nucleo int) ([]models.Person, error) {
	row, err := db.Query("SELECT * FROM persons WHERE nucleo = ?", nucleo)
	if err != nil {
		return []models.Person{}, err
	}
	defer row.Close()

	var persons []models.Person
	for row.Next() {
		var p models.Person
		err = row.Scan(&p.Id, &p.Name, &p.Role, p.MonthlyPayment, &p.Credit)
		if err != nil {
			return []models.Person{}, err
		}
		persons = append(persons, p)
	}

	return persons, err
}

func (db *PersonDB) CountPersons() (int, error) {
	var size int
	err := db.QueryRow("SELECT COUNT(*) FROM persons").Scan(&size)

	return size, err
}

func (db *PersonDB) CountMembers() (int, error) {
	var size int
	err := db.QueryRow("SELECT COUNT(*) FROM persons WHERE role = 'militant' OR role = 'aspirant' OR 'leader' OR 'financeLeader'").Scan(&size)

	return size, err
}

func (db *PersonDB) CountMilitants() (int, error) {
	var size int
	err := db.QueryRow("SELECT COUNT(*) FROM persons WHERE role = 'militant'").Scan(&size)

	return size, err
}

func (db *PersonDB) CountNonAspirantMembers() (int, error) {
	var size int
	err := db.QueryRow("SELECT COUNT(*) FROM persons WHERE role = 'militant' OR role = 'leader' OR role = 'financeLeader'").Scan(&size)

	return size, err
}

func (db *PersonDB) DeletePerson(id string) error {
	_, err := db.Exec("DELETE FROM persons WHERE id = ?", id)

	return err
}

func (db *PersonDB) Promote(id string, role string) error {
	_, err := db.Exec("UPDATE persons SET role = ? WHERE id = ?", role, id)
	return err
}

func (db *PersonDB) UpdateMonthlyContribution(id string, monthlyContribution float64) error {
	_, err := db.Exec("UPDATE persons SET monthlyContribution = ? WHERE id = ?", monthlyContribution, id)
	return err
}

func (db *PersonDB) UpdateCredit(id string, credit float64) error {
	_, err := db.Exec("UPDATE persons SET credit = ? WHERE id = ?", credit, id)
	return err
}

func (db *PersonDB) UpdateNucleo(id string, nucleo int) error {
	_, err := db.Exec("UPDATE persons SET nucleo = ? WHERE id = ?", nucleo, id)
	return err
}
