package main

import (
	_ "beego-demo/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	user := beego.AppConfig.String("mysqluser")
	pass := beego.AppConfig.String("mysqlpass")
	host := beego.AppConfig.String("mysqlurls")
	db := beego.AppConfig.String("mysqldb")
	datasource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", user, pass, host, db)
	orm.RegisterDataBase("default", "mysql", datasource, 30)

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
	}
}

func main() {
	beego.Run()
}
