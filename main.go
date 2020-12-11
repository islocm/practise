package main

import (
	"fmt"

	"github.com/gin-gonic/contrib/static"

	"github.com/gin-gonic/contrib/sessions"

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

	router.LoadHTMLFiles(
		cfg.HTML+"index.html",
		cfg.HTML+"searching.html",
		cfg.HTML+"upload.html",
		cfg.HTML+"information.html",
		cfg.HTML+"error.html",
		cfg.HTML+"template.html",
		cfg.HTML+"selyami.html")

	router.Use(static.Serve("/", static.LocalFile(cfg.Data, false)))

	word := sessions.NewCookieStore([]byte("my-secret-key%%islocm&&random$$examplingget"))

	router.Use(sessions.Sessions("session", word))

	router.GET("/", index)

	router.GET("/searching", searching)

	router.GET("/upload", uploadsearchingget)
	router.POST("/upload", uploadsearchingpost)

	router.POST("/login", login)
	router.POST("/loginexit", loginexit)

	router.GET("/admin", adminGet)

	router.GET("information", information)
	router.POST("information", informationp)
	router.POST("informationpart", informationpart)

	router.GET("/template", template)

	router.GET("s/*kod", selyami)

	router.Run(cfg.ServerHost + ":" + cfg.ServerPort)
}
