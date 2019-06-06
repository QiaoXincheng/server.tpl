package main

import (
	"server.tpl/base"
	"server.tpl/controller"
	"server.tpl/middleware"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/layidao/utilx"
	"gopkg.in/gin-contrib/cors.v1"

	"fmt"
	"runtime"
	"time"
)

var MonitorCount int64
var perTotalCount int64
var currentQps float64
var qpsInterval int = 5
var wanIP, lanIP string
var serverStartTime time.Time

func init() {
	if wanIP == "" && lanIP == "" {
		wanIP, lanIP = utilx.ServerIP()
		serverStartTime = time.Now()
	}
}

func main() {

	// 创建
	r := gin.New()
	r.Use(gin.Recovery())

	// Use Cros For Web Api
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"OPTIONS", "GET", "HEAD", "POST"}
	config.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"}
	config.AllowAllOrigins = true
	config.AllowCredentials = true

	api := r.Group("/v1").Use(middleware.Monitor(&MonitorCount), cors.New(config))
	{
		api.GET("hello-world", controller.HelloWorld)
	}

	// Monitor for local
	monitor := r.Group("/monitor").Use(middleware.AuthIP())
	{
		monitor.GET("/qps", func(ctx *gin.Context) {
			goroutineNum := runtime.NumGoroutine()
			var mem runtime.MemStats
			runtime.ReadMemStats(&mem)
			uptime := time.Now().Sub(serverStartTime)
			uptime = uptime / 1000000000
			ctx.JSON(200, gin.H{"wan_ip": wanIP, "lan_ip": lanIP, "uptime": uptime, "goroutine": goroutineNum, "mem_alloc": (mem.Alloc / 1048576), "mem_heap_sys": (mem.HeapSys / 1048576), "total": MonitorCount, "qps": currentQps})
		})
	}

	r.NoRoute(func(c *gin.Context) {
		c.String(404, "Not Found")
	})

	// 统计qps
	go func() {
		for {
			time.Sleep(time.Duration(qpsInterval) * time.Second)
			currentQps = float64(MonitorCount-perTotalCount) / float64(qpsInterval)
			perTotalCount = MonitorCount
		}
	}()

	// 启动程序
	host := base.Cfg.Section("server").Key("host").MustString("")
	port := base.Cfg.Section("server").Key("port").MustString("8080")
	host = host + ":" + port
	fmt.Println("[succ] Server starting at : ", host)
	endless.DefaultReadTimeOut = 10 * time.Second
	endless.DefaultWriteTimeOut = 10 * time.Second
	endless.ListenAndServe(host, r)
}
