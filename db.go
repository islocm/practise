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
	Queries["select#user#all"], e = db.Prepare(cfg.Selectuserall)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	Queries["selectallfromuser"], e = db.Prepare(cfg.Selectallfromqaror)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	Queries["Selectallfrominformation"], e = db.Prepare(cfg.Selectallfrominformation)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	Queries["Insertalltoinformation"], e = db.Prepare(cfg.Insertalltoinformation)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	Queries["Insertqaror"], e = db.Prepare(cfg.Insertqaror)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	Queries["GetMahallaFromQaror"], e = db.Prepare(cfg.GetMahallaFromQaror)
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
type tableinformation struct {
	Idi             string
	Qaror           string
	Tuman           string
	Mahalla         string
	Kod             string
	Nedvijimost     string
	Pravoobladatel  string
	Soprovoditelniy string
	Pzuo            string
	Po              string
	Pj              string
	Xona            string
	Liver           string
	Degree          string
	Datei           string
	Useri           string
	Complete        string
	Allinformation  []tableinformation
}

type tableselyami struct {
	Ids        string
	Fios       string
	Kods       string
	Births     string
	Relations  string
	Jons       string
	Manzils    string
	Vaqts      string
	Yermaydons string
	Umumiys    string
	Yashashs   string
	Hujjats    string
	Izohs      string
	Times      string
	Users      string
	Anyideas   string
	Allselyami []tableselyami
}

type tabletarkib struct {
	Idt        string
	Fiot       string
	Kodt       string
	Birtht     string
	Relationt  string
	Jont       string
	Manzilt    string
	Vaqtt      string
	Yermaydont string
	Umumiyt    string
	Yashasht   string
	Hujjatt    string
	Izoht      string
	Timet      string
	Usert      string
	Idselyamit string
	Anyideat   string
	Alltarkib  []tabletarkib
}

type tablecompensation struct {
	Idc        string
	Kodimport  string
	Visionc    string
	Kodc       string
	Manzilc    string
	Maydonc    string
	Xonac      string
	Yermaydonc string
	Bozorc     string
	Ijarac     string
	Ijaramc    string
	Protokolc  string
	Orderc     string
	Datac      string
	Userc      string
	Idselyamic string
	Alltarkib  []tabletarkib
}

type tableqaror struct {
	Idq      int
	Qaror    string
	Dateq    string
	Userq    string
	Allqaror []tableqaror
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

type authfirst struct {
	Tel       string `json:"Tel"`
	Name      string `json:"Name"`
	Password  string `json:"Password"`
	Dateu     string `json:"Dateu"`
	Role      string `json:"Role"`
	Districtu string `json:"Districtu"`
	Rolegov   string `json:"Rolegov"`
}

func (i *authfirst) logindb() error {
	err := Queries["select#user#all"].QueryRow(i.Tel, i.Password).Scan(&i.Name, &i.Dateu, &i.Role, &i.Districtu, &i.Rolegov)
	if err != nil {
		fmt.Println(err)
		return errors.New("User not found")
	}

	return nil
}

func (i *tableqaror) Selectqaror() error {
	Rows, err := Queries["selectallfromuser"].Query()
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Problem with selectqaror")
	}

	for Rows.Next() {
		err = Rows.Scan(&i.Idq, &i.Qaror, &i.Dateq, &i.Userq)
		if err != nil {
			fmt.Println(err)
			return err
		}
		i.Allqaror = append(i.Allqaror, tableqaror{Idq: i.Idq, Qaror: i.Qaror, Dateq: i.Dateq, Userq: i.Userq})
	}
	return nil
}

func (i *tableinformation) Select(s ...string) error {

	s[2] = "%" + s[2] + "%"

	Rows, err := Queries["Selectallfrominformation"].Query(s[0], s[1], s[2], s[3])
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Problem with Select information")
	}

	for Rows.Next() {
		err = Rows.Scan(&i.Idi, &i.Qaror, &i.Tuman, &i.Mahalla, &i.Kod, &i.Nedvijimost, &i.Pravoobladatel, &i.Soprovoditelniy, &i.Pzuo, &i.Po, &i.Pj, &i.Xona, &i.Liver, &i.Degree, &i.Datei, &i.Useri, &i.Complete)
		if err != nil {
			fmt.Println(err)
			return err
		}
		i.Allinformation = append(i.Allinformation, tableinformation{Idi: i.Idi, Qaror: i.Qaror, Tuman: i.Tuman, Mahalla: i.Mahalla, Kod: i.Kod, Nedvijimost: i.Nedvijimost, Pravoobladatel: i.Pravoobladatel, Soprovoditelniy: i.Soprovoditelniy, Pzuo: i.Pzuo, Po: i.Po, Pj: i.Pj, Xona: i.Xona, Liver: i.Liver, Degree: i.Degree, Datei: i.Datei, Useri: i.Useri, Complete: i.Complete})
	}
	return nil
}

//InsertI this from excel to table information
func InsertI(i []string) error {

	_, err := Queries["Insertalltoinformation"].Exec(i[0], i[1], i[2], i[3], i[4], i[5], i[6], i[7], i[8], i[9], i[10], i[11], i[12], i[13], i[14])
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Insertq this insert get query to database (url : infoemation)
func Insertq(i ...string) error {
	_, err := Queries["Insertqaror"].Exec(i[0], i[1])
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// MahallaInfo get mahalla group by query qaror
func MahallaInfo(i string) []string {
	Rows, err := Queries["GetMahallaFromQaror"].Query(i)
	if err != nil {
		fmt.Println(err.Error())

		return nil
	}
	var info []string
	for Rows.Next() {

		var temp string
		Rows.Scan(&temp)

		info = append(info, temp)
	}
	return info
}
