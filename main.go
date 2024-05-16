package main

import (
	iauth "gwg/internal/auth"
	icontext "gwg/internal/context"
	ilogger "gwg/internal/logger"
	"gwg/src/middlewares"
	"gwg/src/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	ilogger.Init()

	r.Use(middlewares.DBConnect)
	r.Use(middlewares.SetUser)

	r.GET("/users", list)
	r.GET("/users/info", getInfo)

	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}

func list(c *gin.Context) {
	var (
		conn = icontext.Connection(c)
		s    = service.NewUser(nil)
	)

	res := s.List(conn)
	response(c, 200, res, nil)
}

func getInfo(c *gin.Context) {
	var (
		conn = icontext.Connection(c)
		user = icontext.GetCurrenUser(c).(iauth.User)
		s    = service.NewUser(&user)
	)

	res, err := s.GetInfo(conn)
	if err != nil {
		response(c, 400, nil, err)
		return
	}

	response(c, 200, res, nil)
}

func response(c *gin.Context, code int, data interface{}, err error) {
	res := gin.H{
		"data": data,
	}

	if err != nil {
		res["err"] = err.Error()
	}

	c.JSON(code, res)
}
