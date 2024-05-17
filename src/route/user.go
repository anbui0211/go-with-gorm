package route

import (
	"gwg/src/handler"

	"github.com/gin-gonic/gin"
)

func user(r *gin.RouterGroup) {
	var (
		g = r.Group("/users")
		h = handler.User{}
	)

	g.GET("", h.List)
	g.GET("/info", h.GetInfo)

}
