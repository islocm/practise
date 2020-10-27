package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	e := connect()

	if e != nil {
		fmt.Println(e.Error())
	}

	router := gin.Default()

	router.Static("/assets", cfg.Assets)

	router.LoadHTMLFiles(cfg.HTML+"index.html", cfg.HTML+"searching.html", cfg.HTML+"upload.html")

	router.GET("/", index)

	router.GET("/searching", searching)

	router.GET("/upload", uploadsearchingget)
	router.POST("/upload", uploadsearchingpost)

	router.Run(cfg.ServerHost + ":" + cfg.ServerPort)
}

func index(c *gin.Context) {

	c.HTML(200, "index.html", nil)
}
