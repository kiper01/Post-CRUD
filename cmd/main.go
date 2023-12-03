package main

import (
	"database/sql"
	"fmt"

	"github.com/kiper01/Post-CRUD/migrations"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:postgres@localhost/postgres?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection successful")

	migrations.RunMigrations()
}