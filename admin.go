package main

import "github.com/gin-gonic/gin"

func adminGet(c *gin.Context) {
	c.HTML(200, "upload.html", nil)
}
