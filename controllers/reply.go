package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xiefulaithu/hackathon-server/models"
)

// ReplyController handler all /api/hackathon/reply request
type ReplyController struct {
	beego.Controller
}

// QueryByQuestionID select all reply by question id from mysql
// @router / [get]
func (c *ReplyController) QueryByQuestionID() {
	questionID, err := c.GetInt("qid", -1)
	if questionID != -1 && err == nil {
		c.Data["json"] = models.QueryAllReplyByQuestionID(questionID)
	}
	c.ServeJSON()
}

// ReqReply defineed the form format of post requst
type ReqReply struct {
	Content    string `form:"content"`
	Responder  string `form:"responder"`
	QuestionID int    `form:"qid"`
}

func convert2DBReply(r ReqReply) models.Reply {
	return models.Reply{
		QuestionID: r.QuestionID,
		Content:    r.Content,
		Responder:  r.Responder,
	}
}

// CreateReply function create reply
// @router / [post]
func (c *ReplyController) CreateReply() {
	r := ReqReply{}
	if err := c.ParseForm(&r); err != nil {
		c.Data["json"] = BasicResp{
			StatusCode: 400,
			Message:    "parse form data Err: " + err.Error(),
		}
	} else {
		rr := convert2DBReply(r)
		if err := models.RecordReply(rr); err != nil {
			c.Data["json"] = BasicResp{
				StatusCode: 400,
				Message:    err.Error(),
			}
		} else {
			c.Data["json"] = BasicResp{
				StatusCode: 0,
			}
		}
	}
	c.ServeJSON()
}
