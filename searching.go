package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"

	"github.com/gin-gonic/gin"
)

func searching(c *gin.Context) {

	var selected cadastre

	search := c.Query("search")

	e := selected.SelectAll(search)

	if e != nil {
		fmt.Println(e.Error())
	}

	c.HTML(200, "searching.html", gin.H{
		"Selected": selected.Gis,
	})
}

func uploadsearchingget(c *gin.Context) {
	c.HTML(200, "upload.html", nil)
}

func uploadsearchingpost(c *gin.Context) {

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

	list := excelfile.GetSheetName(0)
	getrow, e := excelfile.GetRows(list)
	if e != nil {
		fmt.Println(e.Error())
	}

	for _, rows := range getrow {

		res, e := Select(rows[1])
		if res != "Error" {
			e := Update(rows)
			if e != nil {
				fmt.Println(e.Error())
			}
		} else {
			e = Insert(rows)
			if e != nil {
				fmt.Println(e.Error())
			}
		}

	}
	c.HTML(200, "upload.html", nil)
}
