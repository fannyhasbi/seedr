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

const (
	commentDB  = "kumpalite_comment"
	categoryDB = "kumpalite_article"
)

// InitCommentDB init database connection
func InitCommentDB() *sql.DB {
	connStr := fmt.Sprintf("postgresql://root@%s?sslmode=disable", databasePort)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE DATABASE IF NOT EXISTS ` + commentDB)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()

	connStr = fmt.Sprintf("postgresql://root@%s/%s?sslmode=disable", databasePort, commentDB)
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

// InitCategoryDB init category db
func InitCategoryDB() *sql.DB {
	connStr := fmt.Sprintf("postgresql://root@%s?sslmode=disable", databasePort)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE DATABASE IF NOT EXISTS ` + categoryDB)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()

	connStr = fmt.Sprintf("postgresql://root@%s/%s?sslmode=disable", databasePort, categoryDB)
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
