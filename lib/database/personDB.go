package database

import (
	"database/sql"

	"github.com/AnnWann/pstu_finance_system/lib/models"
)

type PersonDB struct {
	*sql.DB
}

func (db *DBWrapper) GetPersonDB() PersonDB {
	return PersonDB{db.DB}
}

func (db PersonDB) InsertPerson(p models.Person) error {
	_, err := db.Exec("INSERT INTO persons (name, role, credit) VALUES (?, ?, ?, ?, ?)", p.Id, p.Name, p.Role, p.MonthlyPayment, p.Credit)
	return err
}

func (db PersonDB) GetPerson(id string) (models.Person, error) {
	var p models.Person
	err := db.QueryRow("SELECT * FROM persons WHERE id = ?", id).Scan(&p.Id, &p.Name, &p.Role, p.MonthlyPayment, &p.Credit)

	return p, err
}

func (db PersonDB) GetAllPersons() ([]models.Person, error) {
	var persons []models.Person
	rows, err := db.Query("SELECT * FROM persons")
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

func (db PersonDB) GetMembers() ([]models.Person, error) {
	var members []models.Person
	rows, err := db.Query("SELECT * FROM persons WHERE role = 'militant' OR role = 'aspirant'")
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

func (db PersonDB) GetCore() (models.Person, error) {
	var p models.Person
	err := db.QueryRow("SELECT * FROM persons WHERE id = 'core'").Scan(&p.Id, &p.Name, &p.Role, p.MonthlyPayment, &p.Credit)

	return p, err
}

func (db PersonDB) GetParty() (models.Person, error) {
	var p models.Person
	err := db.QueryRow("SELECT * FROM persons WHERE id = 'party'").Scan(&p.Id, &p.Name, &p.Role, p.MonthlyPayment, &p.Credit)

	return p, err
}

func (db PersonDB) GetPersonByName(name string) (models.Person, error) {
	var p models.Person
	err := db.QueryRow("SELECT * FROM persons WHERE name = ?", name).Scan(&p.Id, &p.Name, &p.Role, p.MonthlyPayment, &p.Credit)

	return p, err
}

func (db PersonDB) GetPersonsByRole(role string) (models.Person, error) {
	var p models.Person
	err := db.QueryRow("SELECT * FROM persons WHERE role = ?", role).Scan(&p.Id, &p.Name, &p.Role, p.MonthlyPayment, &p.Credit)

	return p, err
}

func (db PersonDB) CountPersons() (int, error) {
	var size int
	err := db.QueryRow("SELECT COUNT(*) FROM persons").Scan(&size)

	return size, err
}

func (db PersonDB) CountMembers() (int, error) {
	var size int
	err := db.QueryRow("SELECT COUNT(*) FROM persons WHERE role = 'militant' OR role = 'aspirant' OR 'leader' OR 'financeLeader'").Scan(&size)

	return size, err
}

func (db PersonDB) CountMilitants() (int, error) {
	var size int
	err := db.QueryRow("SELECT COUNT(*) FROM persons WHERE role = 'militant'").Scan(&size)

	return size, err
}

func (db PersonDB) CountNonAspirantMembers() (int, error) {
	var size int
	err := db.QueryRow("SELECT COUNT(*) FROM persons WHERE role = 'militant' OR role = 'leader' OR role = 'financeLeader'").Scan(&size)

	return size, err
}

func (db PersonDB) DeletePerson(id string) error {
	_, err := db.Exec("DELETE FROM persons WHERE id = ?", id)

	return err
}

func (db PersonDB) UpdateRole(id string, role string) error {
	_, err := db.Exec("UPDATE persons SET role = ? WHERE id = ?", role, id)
	return err
}

func (db PersonDB) PromoteToMilitant(id string) error {
	_, err := db.Exec("UPDATE persons SET role = 'militant' WHERE id = ?", id)
	return err
}
func (db PersonDB) PromoteNewLeader(promotee string, demotee models.Person) error {
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = db.Exec("UPDATE persons SET role = 'militant' WHERE id = ?", demotee)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = db.Exec("UPDATE persons SET role = ? WHERE id = ?", demotee.Role, promotee)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
func (db PersonDB) UpdateMonthlyContribution(id string, monthlyContribution int16) error {
	_, err := db.Exec("UPDATE persons SET monthlyContribution = ? WHERE id = ?", monthlyContribution, id)
	return err
}

func (db PersonDB) UpdateCredit(id string, credit float32) error {
	_, err := db.Exec("UPDATE persons SET credit = ? WHERE id = ?", credit, id)
	return err
}
