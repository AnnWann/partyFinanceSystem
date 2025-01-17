package database

import "database/sql"

type PaydayDB struct {
	*sql.DB
}

func (db *DBWrapper) GetPaydayDB() PaydayDB {
	return PaydayDB{db.DB}
}

func (db PaydayDB) GetPayday() (string, error) {
	PAYDAY := ""
	err := db.QueryRow("SELECT * FROM payday").Scan(&PAYDAY)
	return PAYDAY, err
}

func (db PaydayDB) SetPayday(PAYDAY string) error {
	_, err := db.Exec("UPDATE payday SET payday = ?", PAYDAY)
	return err
}
