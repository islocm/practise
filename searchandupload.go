package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func template(c *gin.Context) {

	s := sessions.Default(c)
	name, ok := s.Get("KeyForLogin").(string)
	if !ok {
		c.HTML(200, "Error.html", gin.H{
			"Error": "You are wrong :)",
		})
	}
	c.HTML(200, "template.html", gin.H{
		"Name": name,
	})
}

//searching system with get query
func searching(c *gin.Context) {

	var selected cadastre

	search := c.Query("search")

	if search != "" && len(search) != 1 {

		e := selected.SelectAll(search)

		if e != nil {
			fmt.Println(e.Error())
		}
	}

	c.HTML(200, "searching.html", gin.H{
		"Selected": selected.Gis,
	})
}

//get query for uploading file
func uploadsearchingget(c *gin.Context) {
	c.HTML(200, "upload.html", nil)
}

//post query for uploading file
func uploadsearchingpost(c *gin.Context) {
	var e error
	i := sessions.Default(c)
	name, ok := i.Get("KeyForLogin").(string)
	if !ok {
		c.HTML(200, "error.html", gin.H{
			"Error": "You are wrong :)",
		})
		return
	}

	atoi := c.Request.FormValue("sheetnumber")
	multiupload, e := c.FormFile("upload")
	if e != nil {
		fmt.Println(e.Error())
	}
	e = c.SaveUploadedFile(multiupload, "uploading")
	if e != nil {
		fmt.Println(e.Error())
	}

	excelfile, e := excelize.OpenFile("uploading")
	if e != nil {
		fmt.Println(e.Error())
	}
	sheetnumber, e := strconv.Atoi(atoi)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	list := excelfile.GetSheetName(sheetnumber)
	getrow, e := excelfile.GetRows(list)
	if e != nil {
		fmt.Println(e.Error())
	}
	x := 0
	y := 0
	//first check code from database if it has return update else insert it
	fmt.Println(list)
	if len(getrow[0]) == 27 {

		for _, rows := range getrow {
			res, e := Select(rows[1])
			if e != nil {
				x = x + 1
			}
			if res != "Error" {
				e := Update(rows)
				y = y + 1
				if e != nil {
					fmt.Println(e)
				}
			} else {
				e = Insert(rows)
				if e != nil {
					fmt.Println(e.Error())
				}

			}

		}

	} else if list == "Information" {
		for _, rows := range getrow {

			err := InsertI(rows)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	} else {
		fmt.Println(name)
		e = errors.New("You are wrong with length of rows \n(Sizning xatoyingiz qatorning uzunligi bilan bog'liq)")
	}

	c.HTML(200, "upload.html", gin.H{
		"Updated":  x,
		"Inserted": y,
		"Error":    e,
		"Golang":   "on",
	})
}
