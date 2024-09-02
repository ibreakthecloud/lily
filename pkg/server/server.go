package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ibreakthecloud/lily/pkg/auth"
)

func InitServer() *gin.Engine {
	r := gin.Default()
	r.POST("/login", auth.LoginHandler)
	r.POST("/monte-carlo", auth.Authenticate(), monteCarloIncident)
	r.POST("/annotate", auth.Authenticate(), annotateData)

	return r
}
