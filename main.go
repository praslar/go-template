package main

import (
	"fmt"

	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/pkg/utils"
	"gitlab.com/jamalex-ltd/r-d-department/go-template/routers"

	"github.com/astaxie/beego"
)

var (
	name    = "Variant"
	version = "0.0.0"
)

func main() {
	routers.Initialize(name, version)
	_ = beego.LoadAppConfig("ini", "conf/app.conf")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	err := utils.AutoMigration()
	if err != nil {
		fmt.Println(err.Error())
	}
	beego.Run()
}
