package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:DiseaseDetailController"] = append(beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:DiseaseDetailController"],
		beego.ControllerComments{
			Method: "Query",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:QuestionController"] = append(beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:QuestionController"],
		beego.ControllerComments{
			Method: "LimitedQuery",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:QuestionController"] = append(beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:QuestionController"],
		beego.ControllerComments{
			Method: "CreateQuestion",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:ReplyController"] = append(beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "QueryByQuestionID",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:ReplyController"] = append(beego.GlobalControllerRouter["github.com/xiefulaithu/hackathon-server/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "CreateReply",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
