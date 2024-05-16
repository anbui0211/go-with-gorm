package main

import (
	icontext "gwg/internal/context"
	"gwg/src/middlewares"
	"gwg/src/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.DBConnect)

	r.GET("/users", list)

	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}

func list(c *gin.Context) {
	var (
		conn = icontext.Connection(c)
		s    = &service.User{}
	)

	res := s.List(conn)
	c.JSON(200, gin.H{
		"message": "pong",
		"data":    res,
	})
}
