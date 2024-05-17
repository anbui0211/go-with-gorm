package route

import (
	"gwg/src/middlewares"

	"github.com/gin-gonic/gin"
)

func Route(e *gin.Engine) {
	e.Use(middlewares.DBConnect)
	e.Use(middlewares.SetUser)

	r := e.Group("")

	user(r)
}
