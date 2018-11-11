package main

import (
	"fmt"

	"github.com/xiefulaithu/hackathon-server/models"
	_ "github.com/xiefulaithu/hackathon-server/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	mysql_connect := beego.AppConfig.String("mysql")
	err := models.InitDB(mysql_connect)
	if err != nil {
		fmt.Println(err)
		return
	}
	beego.Run()
}
