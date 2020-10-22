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

}
func prepareQueries() {
	Queries["write#anyone"], e = db.Prepare("INSERT INTO cadastre VALUES")
	if e !
}

type cadastre struct {
	Property  string
	Code      string
	District  string
	Ownership string
	Passport  string
	Document  string
	Regbook   string
	Regpage   string
	Govnum    string
	Govdate   string
	Amount    string
	Rooms     string
	Costa     string
	Costr     string
	Totala    string
	Livinga   string
	Usefula   string
	Pzuo      string
	Pzuz      string
	Pzuzaxvat string
	Pzupd     string
	Pzupp     string
	Npp       string
	Npk       string
	Spp       string
	Spk       string
	Gis       []cadastre
}

func (geo *cadastre) Firstquery() error {

}
