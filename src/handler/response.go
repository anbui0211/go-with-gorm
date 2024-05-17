package handler

import "github.com/gin-gonic/gin"

func response(c *gin.Context, code int, data interface{}, err error) {
	res := gin.H{
		"data": data,
	}

	if err != nil {
		res["err"] = err.Error()
	}

	c.JSON(code, res)
}
