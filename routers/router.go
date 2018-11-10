package routers

import (
	"github.com/astaxie/beego"
	"github.com/xiefulaithu/hackathon-server/controllers"
)

func init() {
	ns := beego.NewNamespace("/api/hackathon",
		beego.NSNamespace("/diseasedetail",
			beego.NSInclude(
				&controllers.DiseaseDetailController{},
			)),
	)

	beego.AddNamespace(ns)
}
