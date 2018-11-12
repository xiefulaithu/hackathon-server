package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xiefulaithu/hackathon-server/models"
)

// HotMapController defines all handler used in /hackathon/hotmap
type HotMapController struct {
	beego.Controller
}

// QueryByName function query all hotpoints
// @router / [get]
func (c *HotMapController) QueryByName() {
	recResult := c.GetString("name", "潜叶蛾")
	hotpoints := models.QueryAllHotPointByRecResult(recResult)
	c.Data["json"] = hotpoints
	c.ServeJSON()
}

// ReqHotPoint defines the form format of post request
type ReqHotPoint struct {
	Name   string `form:"name"`
	Lat    string `form:"lat"`
	Lng    string `form:"lng"`
	PicURL string `form:"pic_url"`
}

func convert2DBHotPoint(hp ReqHotPoint) models.HotPoint {
	return models.HotPoint{
		Name:   hp.Name,
		Lat:    hp.Lat,
		Lng:    hp.Lng,
		PicURL: hp.PicURL,
	}
}

// SaveHotPointByName function create record in "hotmap" table
// @router / [post]
func (c *HotMapController) SaveHotPointByName() {
	hp := ReqHotPoint{}
	if err := c.ParseForm(&hp); err != nil {
		c.Data["json"] = BasicResp{
			StatusCode: 400,
			Message:    "parse form data Err: " + err.Error(),
		}
	} else {
		rhp := convert2DBHotPoint(hp)
		if err := models.RecordHotPoint(rhp); err != nil {
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
