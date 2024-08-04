package db

import (
	"database/sql"
	"fmt"
)

var Postgres *sql.DB
var err error

func ConnectDB() {
	Postgres, err = sql.Open("postgres", "host=pg-cool-dev.postgres.database.azure.com  port=5432 user=cool-developer password=ip_e28Da dbname=cool")
	if err != nil {
		fmt.Println("DB Connection error:", err)
	}

	// verify cridentials
	err = Postgres.Ping()
	if err != nil {
		fmt.Println("cridentials not correct::", err)
	} else {
		fmt.Println("Postgres databse connected!!")
	}
}
