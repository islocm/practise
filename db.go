package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

//Queries kk
var Queries map[string]*sql.Stmt

func connect() error {
	var e error

	db, e = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPass, cfg.PgBase))
	if e != nil {
		fmt.Println(e.Error())
		return e
	}
	return nil

	Queries["hello"] := db.Prepare("select * from users")
	Queries["hello"].Query()

}
