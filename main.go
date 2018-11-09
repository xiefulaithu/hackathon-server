package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/xiefulaithu/hackathon-server/models"
	_ "github.com/xiefulaithu/hackathon-server/routers"
)

func main() {

	// load conf from conf/app.conf
	mysql_connect := beego.AppConfig.String("mysql")
	err := models.InitDB(mysql_connect)
	if err != nil {
		fmt.Println(err)
		return
	}
	beego.Run()
}
