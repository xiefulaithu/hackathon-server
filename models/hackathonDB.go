package models

import (
	"errors"
	"time"

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

// Question defines the Question Message struct in hackathon backend server
type Question struct {
	ID         int       `json:"id" gorm:"column:id"`
	Content    string    `json:"content" gorm:"column:content"`
	Pics       string    `json:"pics" gorm:"column:pics"`
	Location   string    `json:"location" gorm:"column:location"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	Questioner string    `json:"questioner" gorm:"column:questioner"`
	Deleted    int       `json:"deleted" gorm:"column:deleted"`
	ReplyNum   int       `json:"reply_num" gorm:"-"`
}

// QueryLatest13Question query latest quetsion and limited to 13 records
func QueryLatest13Question() []Question {
	ret := make([]Question, 0)
	table := db.Table("question")
	table.Limit(13).Order("id desc").Where("deleted = 0").Find(&ret)

	for k := 0; k < len(ret); k++ {
		ret[k].ReplyNum = GetReplyNumByQuestionID(ret[k].ID)
	}
	return ret
}

// RecordQuestion insert question to mysql
func RecordQuestion(q Question) error {
	table := db.Table("question")
	if !table.NewRecord(&q) {
		return errors.New("invalid pri key in question table")
	}

	localdb := table.Create(&q)
	return errgroup.NewMultiError(localdb.GetErrors()...)
}

// Reply defines the Reply Message struct in hackathon backend server
type Reply struct {
	QuestionID int       `json:"qid" gorm:"column:question_id"`
	Content    string    `json:"content" gorm:"column:content"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	Responder  string    `json:"responder" gorm:"column:responder"`
}

// QueryAllReplyByQuestionID query all replys by question_id
func QueryAllReplyByQuestionID(qid int) []Reply {
	ret := make([]Reply, 0)
	table := db.Table("reply")
	table.Order("create_time desc").Where("question_id = ?", qid).Find(&ret)
	return ret
}

// GetReplyNumByQuestionID get how many replys for a question
func GetReplyNumByQuestionID(qid int) int {
	replys := make([]Reply, 0)
	count := 0
	table := db.Table("reply")
	table.Where("question_id = ?", qid).Find(&replys).Count(&count)
	return count
}

// RecordReply insert reply to mysql
func RecordReply(q Reply) error {
	table := db.Table("reply")
	if !table.NewRecord(&q) {
		return errors.New("invalid pri key in reply table")
	}
	localdb := table.Create(&q)
	return errgroup.NewMultiError(localdb.GetErrors()...)
}

// HotPoint defines the HotPoint struct in hackathon backend server
type HotPoint struct {
	Name   string `json:"name" gorm:"column:rec_result"`
	Lat    string `json:"lat" gorm:"column:lat"`
	Lng    string `json:"lng" gorm:"column:lng"`
	PicURL string `json:"pic_url" gorm:"column:pic_url"`
}

// QueryAllHotPointByRecResult query all hotpoint by disease name
func QueryAllHotPointByRecResult(recResult string) []HotPoint {
	ret := make([]HotPoint, 0)
	table := db.Table("hotmap")
	table.Where("rec_result = ?", recResult).Find(&ret)
	return ret
}

// RecordHotPoint insert hotpoint to table "hotmap"
func RecordHotPoint(hp HotPoint) error {
	table := db.Table("hotmap")
	if !table.NewRecord(&hp) {
		return errors.New("invalid pri key in reply table")
	}
	localdb := table.Create(&hp)
	return errgroup.NewMultiError(localdb.GetErrors()...)
}
