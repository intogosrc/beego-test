package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/go-sql-driver/mysql"

)

func InitSQL() {

	conf := &mysql.Config{
		User:                 beego.AppConfig.String("mysqluser"),
		Passwd:               beego.AppConfig.String("mysqlpass"),
		Net:                  "tcp",
		Addr:                 beego.AppConfig.String("mysqlurls"),
		DBName:               beego.AppConfig.String("mysqldb"),
		AllowNativePasswords: true,
	}

	orm.RegisterDataBase("default", "mysql", conf.FormatDSN(), 30)
	orm.RegisterModel(
		new(UserModel),
	)
}