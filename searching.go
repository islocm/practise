package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"

	"github.com/gin-gonic/gin"
)

func searching(c *gin.Context) {
	c.HTML(200, "searching.html", nil)
}

func uploadsearching(c *gin.Context) {
	c.HTML(200, "upload.html", nil)
}

func uploadsearching1(c *gin.Context) {

	multiupload, e := c.FormFile("upload")
	if e != nil {
		fmt.Println(e.Error())
		panic("error in 15 row (searching.go)")
	}
	e = c.SaveUploadedFile(multiupload, "uploading")
	if e != nil {
		fmt.Println(e.Error())

		panic("error in 20 row (searching.go)")
	}
	excelfile, e := excelize.OpenFile("uploading")
	if e != nil {
		fmt.Println(e.Error())
		panic("error in 26 row (searching.go)")
	}
	list := excelfile.GetSheetName(0)
	getrow, e := excelfile.GetRows(list)
	if e != nil {
		fmt.Println(e.Error())
		panic("error in 38 row (searching.go)")

	}
	for _, rows := range getrow {
		fmt.Println(rows)

	}
	c.HTML(200, "upload.html", nil)
}
