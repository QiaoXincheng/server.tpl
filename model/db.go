package model

import (
	"server.tpl/base"
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"time"
	"fmt"
)

var orm *xorm.Engine

func init() {
	SetEngine()
}

func SetEngine() {
	var err error
	host := base.Cfg.Section("mysql_h5").Key("HOST").MustString("")
	port := base.Cfg.Section("mysql_h5").Key("PORT").MustString("3306")
	username := base.Cfg.Section("mysql_h5").Key("USERNAME").MustString("")
	password := base.Cfg.Section("mysql_h5").Key("PASSWORD").MustString("")
	dbname := base.Cfg.Section("mysql_h5").Key("DBNAME").MustString("")

	orm, err = xorm.NewEngine("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8")
	if err != nil {
		fmt.Println("connect MySQL failed :", err)
	}
	orm.TZLocation = time.Local
	// orm.ShowSQL = true
	showSql := base.Cfg.Section("mysql_h5").Key("SHOW_SQL").MustBool(false)
	orm.ShowSQL(showSql)
}
