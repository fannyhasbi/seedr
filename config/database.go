package config

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
)

var (
	databasePort = ":26257"
)

const dbName = "kumpalite_comment"

// InitCockroachDB init database connection
func InitCockroachDB() *sql.DB {
	connStr := fmt.Sprintf("postgresql://root@%s?sslmode=disable", databasePort)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE DATABASE IF NOT EXISTS ` + dbName)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()

	connStr = fmt.Sprintf("postgresql://root@%s/%s?sslmode=disable", databasePort, dbName)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	ddl, err := ioutil.ReadFile("database/cockroach/ddl.sql")
	if err != nil {
		log.Fatal(err)
	}

	sql := string(ddl)
	_, err = db.Exec(sql)
	if err != nil {
		log.Println(err)
	}

	return db
}
