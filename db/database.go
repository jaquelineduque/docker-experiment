package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

func CreateConnection() (db *sql.DB) {

	dbDriver := "postgres"
	dbHost := os.Getenv("DB_HOST")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err.Error())
	}
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	psql := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)


	db, err = sql.Open(dbDriver, psql)
	/*
	   	"host=%s port=%d user=%s "+
	       "password=%s dbname=%s sslmode=disable
	*/
	if err != nil {
		panic(err.Error())
	}

	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
