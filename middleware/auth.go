package middleware

import (
	"server.tpl/base"
	

	"github.com/gin-gonic/gin"
	"github.com/layidao/utilx"


	"strings"
	
)

func AuthIP() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		allowIP := base.Cfg.Section("monitor").Key("ALLOW_IP").MustString("")
		ipList := strings.Split(allowIP, ",")

		ip := ctx.ClientIP()
		if strings.HasPrefix(ip, "10.") || strings.HasPrefix(ip, "192.") {
			return
		}
		if !utilx.SliceContainsString(ipList, ip) {
			ctx.JSON(403, gin.H{"code": 403, "message": "Forbidden Access"})
			ctx.Abort()
			return
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}


