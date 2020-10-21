package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type setting struct {
	ServerHost string
	ServerPort string
	PgHost     string
	PgPort     string
	PgUser     string
	PgPass     string
	PgBase     string
	Data       string
	Assets     string
	HTML       string
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
