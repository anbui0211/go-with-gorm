package handler

import (
	iauth "gwg/internal/auth"
	icontext "gwg/internal/context"
	"gwg/src/service"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (User) List(c *gin.Context) {
	var (
		conn = icontext.Connection(c)
		s    = service.NewUser(nil)
	)

	res := s.List(conn)
	response(c, 200, res, nil)
}

func (User) GetInfo(c *gin.Context) {
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
