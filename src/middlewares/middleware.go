package middlewares

import (
	iauth "gwg/internal/auth"
	icontext "gwg/internal/context"
	db "gwg/src/database"

	"github.com/gin-gonic/gin"
)

func DBConnect(c *gin.Context) {
	conn, err := db.OpenDB()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error opening database",
			"err":     err.Error(),
		})

	}
	c.Set(icontext.DataBaseContext, conn)
	c.Next()
}

func SetUser(c *gin.Context) {
	user := iauth.User{
		ID:   "US01",
		Name: "Ân",
	}
	c.Set(icontext.UserKeyContext, user)
	c.Next()
}
