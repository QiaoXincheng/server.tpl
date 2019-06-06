package base

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"

	"fmt"
	"os"
)

var Cfg *ini.File

func init() {

	if len(os.Args) > 1 {
		mode := os.Args[1]
		if mode == "prod" {
			loadConf("conf/conf.prod.ini")
			gin.SetMode(gin.ReleaseMode)
		} else if mode == "test" {
			loadConf("conf/conf.test.ini")
			gin.SetMode(gin.TestMode)
		} else if mode == "dev" {
			loadConf("conf/conf.dev.ini")
			gin.SetMode(gin.DebugMode)
		} else {
			os.Exit(1)
		}
	} else {
		os.Exit(1)
	}
}

func loadConf(conf_file string) {
	var err error
	Cfg, err = ini.Load(conf_file)
	if err != nil {
		fmt.Println("[error] Error in loading config", err)
		os.Exit(1)
	} else {
		fmt.Println("[succ] Loaded config from api ", conf_file)
	}
}
