package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func ConnectSQL(host, port, dbname string) (*DB, error) {
	dbSource := fmt.Sprintf(
		"postgresql://root@%s:%s/%s?sslmode=disable",
		host,
		port,
		dbname)

	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		panic(err)
	}

	dbConn.SQL = db
	return  dbConn, err
}