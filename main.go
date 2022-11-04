package main

import (
	"database/sql"
	"fmt"
	"net/http"

	postgresDB "github.com/haagor/orderMP/adapter"
	"github.com/haagor/orderMP/controller"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		postgresDB.Host, postgresDB.Port, postgresDB.User, postgresDB.Password, postgresDB.Dbname)

	var err error
	var db *sql.DB
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	dba := postgresDB.PostgresAdapter{db}

	http.Handle("/ticket", controller.OrderHandler(dba))
	http.ListenAndServe(":8080", nil)
}
