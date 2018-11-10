package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xiefulaithu/hackathon-server/models"
)

// DiseaseDetailController handler all /api/hackathon/diseasedetail request
type DiseaseDetailController struct {
	beego.Controller
}

// Get deals http get request
// @router / [get]
func (c DiseaseDetailController) Get() {
	diseaseName := c.GetString("name")
	if diseaseName != "" {
		c.Data["json"] = models.LoadDiseaseDetailByName(diseaseName)
	}
	c.ServeJSON()
}
