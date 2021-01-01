package main

import (
	"fmt"

	"go-template/internal/pkg/utils"
	"go-template/routers"

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
