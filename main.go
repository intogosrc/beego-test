package main

import (
	"github.com/astaxie/beego"
	"github.com/intogosrc/beego-test/models"
	_ "github.com/intogosrc/beego-test/routers"
)

func main() {

	models.InitSQL()
	beego.Run()
}

