package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type setting struct {
	ServerHost               string
	ServerPort               string
	PgHost                   string
	PgPort                   string
	PgUser                   string
	PgPass                   string
	PgBase                   string
	Data                     string
	Assets                   string
	HTML                     string
	Selectcadastrecode       string
	Cadastre                 string
	Updatecadastre           string
	Selectcadastreall        string
	Selectuserall            string
	Selectallfromqaror       string
	Selectallfrominformation string
	Insertalltoinformation   string
	Insertqaror              string
	GetMahallaFromQaror      string
}

var cfg setting

func init() {

	jsopen, e := os.Open("settings.cfg")
	if e != nil {
		fmt.Println(e.Error())

		panic("Json ochilmadi")
	}
	defer jsopen.Close()

	jsstat, e := jsopen.Stat()
	if e != nil {
		fmt.Println(e.Error())

		panic("jsstatda hatolik")
	}

	jssize := jsstat.Size()

	info := make([]byte, jssize)

	_, e = jsopen.Read(info)

	if e != nil {
		fmt.Println(e.Error())

		panic("o'qilmadi")
	}

	e = json.Unmarshal(info, &cfg)
	if e != nil {
		fmt.Println(e.Error())
	}
}

func random() string {
	rand.Seed(time.Now().Unix())
	in := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ12345678910112"
	inRune := []rune(in)
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
