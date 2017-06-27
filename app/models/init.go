package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
)

var (
	DSN    string
	engine *xorm.Engine
)

func InitDB() {
	var (
		found bool
		err   error
	)
	if DSN, found = revel.Config.String("DSN"); !found {
		time.Sleep(time.Duration(10) * time.Second)
		revel.ERROR.Fatalln("can not found DSN")
		return
	}
	if engine, err = xorm.NewEngine("mysql", DSN); err != nil {
		revel.ERROR.Fatalln("link to DB error,", err)
		return
	}
	if err = engine.Sync2(&User{}); err != nil {
		revel.ERROR.Fatalln("sync user failed,", err)
		return
	}
	if err = engine.Sync2(&Admin{}); err != nil {
		revel.ERROR.Fatalln("sync admin failed,", err)
		return
	}

	revel.INFO.Println("make new engine done--------------------------------------")
}
