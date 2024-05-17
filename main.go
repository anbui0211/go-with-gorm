package main

import (
	ilogger "gwg/internal/logger"
	"gwg/src/route"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	ilogger.Init()
	route.Route(r)

	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}
