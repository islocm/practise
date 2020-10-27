package main

import (
	"database/sql"
	"errors"
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

	Queries = make(map[string]*sql.Stmt)
	prepareQueries()
	return nil

}
func prepareQueries() {
	var e error
	Queries["create#database"], e = db.Prepare("CREATE DATABASE web WITH OWNER = postgres ENCODING = 'UTF8' CONNECTION LIMIT = -1;")
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	Queries["write#cadastre"], e = db.Prepare(cfg.Cadastre)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	Queries["select#code"], e = db.Prepare(cfg.Selectcadastrecode)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	Queries["update#cadastre"], e = db.Prepare(cfg.Updatecadastre)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	Queries["select#cadastreall"], e = db.Prepare(cfg.Selectcadastreall)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
}

//Cadastre struct here...
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

//Insert insert to database
func Insert(i []string) error {

	stmt, ok := Queries["write#cadastre"]

	if !ok {
		return errors.New("Cadastre query for insert")
	}

	_, e := stmt.Exec(i[0], i[1], i[2], i[3], i[4], i[5], i[6], i[7], i[8], i[9], i[10], i[11], i[12], i[13], i[14], i[15], i[16], i[17], i[18], i[19], i[20], i[21], i[22], i[23], i[24], i[25])

	if e != nil {
		fmt.Println(e.Error())

		return e
	}

	return nil

}

//Select function for check code
func Select(s string) (string, error) {
	resulterror := "Error"
	stmt, ok := Queries["select#code"]
	if !ok {
		return resulterror, errors.New("select#code is false")

	}
	var result string
	err := stmt.QueryRow(s).Scan(&result)
	if err != nil {
		return resulterror, err
	}
	//Hello ku
	return result, nil
}

//Update update to database
func Update(i []string) error {

	stmt, ok := Queries["update#cadastre"]

	if !ok {
		return errors.New("Cadastre query for update")
	}

	_, e := stmt.Exec(i[0], i[2], i[3], i[4], i[5], i[6], i[7], i[8], i[9], i[10], i[11], i[12], i[13], i[14], i[15], i[16], i[17], i[18], i[19], i[20], i[21], i[22], i[23], i[24], i[25], i[1])

	if e != nil {
		fmt.Println(e.Error())

		return e
	}

	return nil

}

func (c *cadastre) SelectAll(qwe string) error {
	qwe = "%" + qwe + "%"
	Row, err := Queries["select#cadastreall"].Query(qwe)

	if err != nil {
		fmt.Println(err)
		return err
	}

	for Row.Next() {

		e := Row.Scan(&c.Property, &c.Code, &c.District, &c.Ownership, &c.Passport, &c.Document, &c.Regbook, &c.Regpage, &c.Govnum, &c.Govdate, &c.Amount, &c.Rooms, &c.Costa, &c.Costr, &c.Totala, &c.Livinga, &c.Usefula, &c.Pzuo, &c.Pzuz, &c.Pzuzaxvat, &c.Pzupd, &c.Pzupp, &c.Npp, &c.Npk, &c.Spp, &c.Spk)

		if e != nil {
			fmt.Println(e.Error())
			return e
		}

		c.Gis = append(c.Gis, cadastre{Property: c.Property, Code: c.Code, District: c.District, Ownership: c.Ownership, Passport: c.Passport, Document: c.Document, Regbook: c.Regbook, Regpage: c.Regpage, Govnum: c.Govnum, Govdate: c.Govdate, Amount: c.Amount, Rooms: c.Rooms, Costa: c.Costa, Costr: c.Costr, Totala: c.Totala, Livinga: c.Livinga, Usefula: c.Usefula, Pzuo: c.Pzuo, Pzuz: c.Pzuz, Pzuzaxvat: c.Pzuzaxvat, Pzupd: c.Pzupd, Pzupp: c.Pzupp, Npp: c.Npp, Npk: c.Npk, Spp: c.Spp, Spk: c.Spk})

	}

	return nil
}
