package models

import (
	"errors"

	"github.com/facebookgo/errgroup"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

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

func LoadDiseaseDetailByName(name string) DiseaseDetail {
	ret := DiseaseDetail{}
	table := db.Table("diseasedetail")
	table.Where("diseasename = ?", name).First(&ret)

	pics := make([]DiseaseTemplatePics, 0)
	db.Table("diseasetemplatepics").Where("diseasename = ?", name).Find(&pics)

	ret.Pics = pics
	return ret
}

// Question defines the Question type in hackathon backend server
type Question struct {
	Content    string `json:"content" gorm:"column:content"`
	Pics       string `json:"pics" gorm:"column:pics"`
	Location   string `json:"location" gorm:"column:location"`
	CreateTime string `json:"create_time" gorm:"column:create_time"`
	Questioner string `json:"questioner" gorm:"column:questioner"`
	Deleted    int    `json:"deleted" gorm:"column:deleted"`
}

// QueryLatest13Question query latest quetsion and limited to 13 records
func QueryLatest13Question() []Question {
	ret := make([]Question, 0)
	table := db.Table("question")
	table.Limit(13).Order("id desc").Where("deleted = 1").Find(&ret)
	return ret
}

// RecordQuestion insert question to mysql
func RecordQuestion(q Question) error {
	table := db.Table("question")
	if !table.NewRecord(&q) {
		return errors.New("invalid pri key")
	}

	localdb := table.Create(&q)
	return errgroup.NewMultiError(localdb.GetErrors()...)
}
