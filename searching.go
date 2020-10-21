package main

import (
	"github.com/gin-gonic/gin"
)

func searching(c *gin.Context) {
	
	c.HTML(200, "searching.html", nil)
}
