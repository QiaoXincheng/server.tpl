package middleware

import (
	"server.tpl/service"
	
	"github.com/gin-gonic/gin"
)

func Monitor(MonitorCount *int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		service.LogAccess(c.ClientIP(), c.Request.Method+"|"+c.Request.RequestURI)
		*MonitorCount++
	}
}
