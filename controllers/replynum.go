package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xiefulaithu/hackathon-server/models"
)

// ReplyNumController handler all /api/hackathon/replynum requst
type ReplyNumController struct {
	beego.Controller
}

// QueryReplyNum count reply nums by question id
// @router / [get]
func (c *ReplyNumController) QueryReplyNum() {
	qid, err := c.GetInt("qid", -1)
	if qid != -1 && err == nil {
		c.Data["json"] = BasicResp{
			StatusCode: 0,
			Data:       models.GetReplyNumByQuestionID(qid),
		}
	} else {
		c.Data["json"] = BasicResp{
			StatusCode: 400,
			Message:    "bad request",
		}
	}
	c.ServeJSON()
}
