// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/xiefulaithu/hackathon-server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/hackathon",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/diseasedetail",
			beego.NSInclude(
				&controllers.DiseaseDetailController{},
			),
		),
		beego.NSNamespace("/question",
			beego.NSInclude(
				&controllers.QuestionController{},
			),
		),
		beego.NSNamespace("/reply",
			beego.NSInclude(
				&controllers.ReplyController{},
			),
		),
		beego.NSNamespace("/hotmap",
			beego.NSInclude(
				&controllers.HotMapController{},
			),
		),
		beego.NSNamespace("/replynum",
			beego.NSInclude(
				&controllers.ReplyNumController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
