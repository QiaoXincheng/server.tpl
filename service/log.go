package service

import (
	"server.tpl/base"
	
	"github.com/layidao/logs"
	"github.com/layidao/utilx"

	"strconv"
)

var err *logs.BeeLogger
var access *logs.BeeLogger
var exception *logs.BeeLogger

func init() {

	chanLength := base.Cfg.Section("log").Key("log_message_chan_length").MustInt(1000)

	errorConf := base.Cfg.Section("logger").Key("error").MustString("")
	err = logs.NewLogger(int64(chanLength))
	err.SetLogger("file", errorConf)

	accessConf := base.Cfg.Section("logger").Key("access").MustString("")
	access = logs.NewLogger(int64(chanLength))
	access.SetLogger("file", accessConf)

	exceptionConf := base.Cfg.Section("logger").Key("exception").MustString("")
	exception = logs.NewLogger(int64(chanLength))
	exception.SetLogger("file", exceptionConf)

}

func LogError(content string) {
	content = commonContent() + "|" + content
	err.Info(content)
}

func LogAccess(clientIP string, content string) {
	content = commonContent() + "|" + clientIP + "|" + content
	access.Info(content)
}

func LogException(clientIP string, content string) {
	content = commonContent() + "|" + clientIP + "|" + content
	exception.Info(content)
}

func commonContent() string {
	currTime := utilx.GetCurrentTime()
	currTimeString := strconv.Itoa(currTime)
	serverIP, lanIP := utilx.ServerIP()
	if serverIP == "" {
		serverIP = lanIP
	}
	currDateTime := utilx.GetCurrentDate("2006-01-02 15:04:05")
	str := currDateTime + "|" + currTimeString + "|" + serverIP
	return str
}
