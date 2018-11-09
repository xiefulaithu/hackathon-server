package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func init() {
	fmt.Println("init models")
}

type DiseaseDetail struct {
	DiseaseName string                `json:"name" gorm:"column:diseasename"`
	Brief       string                `json:"brief" gorm:"column:brief"`
	Incident    string                `json:"incident" gorm:"column:incident"`
	MainReason  string                `json:"main_reason" gorm:"column:main_reason"`
	Medicament  string                `json:"medicament" gorm:"column:medicament"`
	Pics        []DiseaseTemplatePics `json:"pics"`
}

type DiseaseTemplatePics struct {
	URL   string `json:"url" gorm:"column:url"`
	Title string `json:"title" gorm:"title"`
}

var (
	db *gorm.DB
)

func InitDB(connect string) error {
	d, err := gorm.Open("mysql", connect)
	if err != nil {
		return err
	}
	db = d
	return nil
}

func LoadDiseaseDetailByName(name string) DiseaseDetail {
	ret := DiseaseDetail{}
	table := db.Table("diseasedetail")
	table.Where("diseasename = ?", name).First(&ret)

	pics := make([]DiseaseTemplatePics, 0)
	db.Table("diseasetemplatepics").Where("diseasename = ?", name).Find(&pics)

	ret.Pics = pics
	return ret
}
