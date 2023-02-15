package Database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() (db *sql.DB) {
	databaseDriver := "mysql"
	databaseUsername := "root"
	databasePassword := ""
	databaseName := "go-mysql"
	db, err := sql.Open(databaseDriver, databaseUsername+":"+databasePassword+"@/"+databaseName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
