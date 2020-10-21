package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/assets", cfg.Assets)

	router.LoadHTMLFiles(cfg.HTML+"index.html", cfg.HTML+"searching.html")

	router.GET("/", index)

	router.Any("/searching", searching)

	router.Run(cfg.ServerHost + ":" + cfg.ServerPort)
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
