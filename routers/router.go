package routers

import (
	"github.com/astaxie/beego"
	"github.com/xiefulaithu/hackathon-server/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("api/hackathon/diseasedetail", &controllers.DiseaseDetailController{})
}
