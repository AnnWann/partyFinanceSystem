package database

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(file string) {
	var err error
	DB, err = sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}

	DB.Exec("CREATE TABLE IF NOT EXISTS persons (id INTEGER NOT NULL PRIMARY KEY," +
		"name TEXT NOT NULL, " +
		"role TEXT NOT NULL, " +
		"monthly_contribution REAL NOT NULL, " +
		"credit float NOT NULL)");

	DB.Exec("CREATE TABLE IF NOT EXISTS registers (id INTEGER NOT NULL PRIMARY KEY," +
		"day TEXT NOT NULL, " +
		"month TEXT NOT NULL, " +
		"year TEXT NOT NULL, " +
		"type TEXT NOT NULL, " +
		"giver TEXT NOT NULL, " +
		"receiver TEXT NOT NULL, " +
		"ammount INTEGER NOT NULL, " +
		"value float NOT NULL, " +
		"party_share float NOT NULL, " +
		"description TEXT NOT NULL)" + 
		"FOREIGN KEY (giver) REFERENCES persons(id)" +
		"FOREIGN KEY (receiver) REFERENCES persons(id)");
	
	DB.Exec("CREATE TABLE IF NOT EXISTS reports (id INTEGER NOT NULL PRIMARY KEY," +
		"month TEXT NOT NULL, " +
		"year TEXT NOT NULL, " +
		"members TEXT NOT NULL, " +
		"membersPayments TEXT NOT NULL, " +
		"expenses TEXT NOT NULL, " +
		"sales TEXT NOT NULL, " +
		"totalEarned float NOT NULL, " +
		"totalLiquid float NOT NULL, " +
		"partyDebts float NOT NULL, " +
		"coreSurplus float NOT NULL)" +
		"foreign key (members) references persons(id)" + 
		"foreign key (membersPayments) references registers(id)" +
		"foreign key (expenses) references registers(id)" +
		"foreign key (sales) references registers(id)");

}
