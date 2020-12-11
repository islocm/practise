package main

import (
	"fmt"

	"github.com/gin-gonic/contrib/sessions"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	IsLogin := false
	s := sessions.Default(c)
	name, ok := s.Get("KeyForLogin").(string)
	if ok {
		IsLogin = ok
	}

	c.HTML(200, "index.html", gin.H{
		"IsLogin": IsLogin,
		"Name":    name,
	})
}

func login(c *gin.Context) {

	var id authfirst

	e := c.BindJSON(&id)
	if e != nil {
		fmt.Println(e.Error())
	}
	e = id.logindb()
	if e != nil {
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}
	session := sessions.Default(c)
	session.Set("KeyForLogin", id.Name)
	session.Set("KeyForDistrict", id.Districtu)

	var config sessions.Options
	// config.Path = "/login"
	// config.Domain = "null"
	config.MaxAge = 0
	// config.Secure = true
	config.HttpOnly = true
	session.Options(config)

	e = session.Save()
	if e != nil {
		fmt.Println(e.Error())
	}
	c.JSON(200, gin.H{
		"Error": nil,
		"Name":  id.Name,
	})
}
func loginexit(c *gin.Context) {

	session := sessions.Default(c)

	session.Delete("KeyForLogin")
	e := session.Save()
	if e != nil {
		fmt.Println(e.Error())
	}

}

func information(c *gin.Context) {
	s := sessions.Default(c)
	name, ok := s.Get("KeyForLogin").(string)
	district := s.Get("KeyForDistrict")
	if !ok {
		c.HTML(200, "error.html", gin.H{
			"Error": "You are wrong :)",
		})
		return

	}
	getqaror, _ := c.GetQuery("addqaror")
	if len(getqaror) > 6 {
		Insertq(getqaror, name)
	}
	var qaror tableqaror

	err := qaror.Selectqaror()
	if err != nil {
		fmt.Println(err)
		return
	}

	c.HTML(200, "information.html", gin.H{
		"Name":     name,
		"Qaror":    qaror.Allqaror,
		"District": district,
	})

}

type modelinformation struct {
	Input   string `json:"Input"`
	Qaror   string `json:"Qaror"`
	Procent string `json:"Mahalla"`
}

type modelmahalla struct {
	Decisio string `json:"Decision"`
}

func informationp(c *gin.Context) {
	var err error
	s := sessions.Default(c)
	name, ok := s.Get("KeyForLogin").(string)
	if !ok {
		c.HTML(200, "error.html", gin.H{
			"Error": "You are wrong :)",
		})
		return
	}
	district, ok := s.Get("KeyForDistrict").(string)
	if !ok {
		c.HTML(200, "error.html", gin.H{
			"Error": "You are wrong :)",
		})
		return

	}

	var bind modelinformation

	err = c.BindJSON(&bind)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var Information tableinformation
	err = Information.Select(bind.Qaror, district, bind.Input, bind.Procent)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"Name":     name,
		"District": district,
		"All":      Information.Allinformation,
	})

}

func informationpart(c *gin.Context) {
	var err error
	s := sessions.Default(c)
	_, ok := s.Get("KeyForLogin").(string)
	if !ok {
		c.HTML(200, "error.html", gin.H{
			"Error": "You are wrong :)",
		})
		return
	}

	var dis modelmahalla
	err = c.BindJSON(&dis)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Area := MahallaInfo(dis.Decisio)

	c.JSON(200, gin.H{

		"Area": Area,
	})
}

func selyami(c *gin.Context) {

	s := sessions.Default(c)
	name, ok := s.Get("KeyForLogin").(string)
	if !ok {
		c.HTML(200, "Error.html", gin.H{
			"Error": "You are wrong :)",
		})
	}

	c.HTML(200, "selyami.html", gin.H{
		"Name": name,
	})
}
