package database

func GetPayday() (string, error) {
	PAYDAY := ""
	err := DB.QueryRow("SELECT * FROM payday").Scan(&PAYDAY)
	return PAYDAY, err
}

func SetPayday(PAYDAY string) error {
	_, err := DB.Exec("UPDATE payday SET payday = ?", PAYDAY)
	return err
}
