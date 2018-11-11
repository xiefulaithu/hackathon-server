package controllers

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/xiefulaithu/hackathon-server/models"
)

// QuestionController deals all /api/hackathon/question request
type QuestionController struct {
	beego.Controller
}

// LimitedQuery function query all questions
// @APIVersion 1.0.0
// @Title question query api
// @Description LimitedQuery() function query "question" table in mysql and return all active questions
// @Contact xiefulai@sogou-inc.com xiefulaithu@163.com
// @License Apache2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
// @Success 200
// @Failure 400
// @router / [get]
func (c *QuestionController) LimitedQuery() {
	questions := models.QueryLatest13Question()
	c.Data["json"] = questions
	c.ServeJSON()
}

// ReqQuestion defined the form format of post request
type ReqQuestion struct {
	Content    string `form:"content"`
	Pics       string `form:"pics"`
	Location   string `form:"location"`
	Questioner string `form:"questioner"`
}

func convert2DBQuestion(q ReqQuestion) models.Question {
	return models.Question{
		Content:    q.Content,
		Pics:       q.Pics,
		Location:   q.Location,
		CreateTime: time.Now(),
		Questioner: q.Questioner,
	}
}

// CreateQuestion function create record in "question" table
// @router / [post]
func (c *QuestionController) CreateQuestion() {
	q := ReqQuestion{}
	if err := c.ParseForm(&q); err != nil {
		c.Data["json"] = BasicResp{
			StatusCode: 400,
			Message:    "parse form data Err: " + err.Error(),
		}
	} else {
		rq := convert2DBQuestion(q)
		if err := models.RecordQuestion(rq); err != nil {
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
