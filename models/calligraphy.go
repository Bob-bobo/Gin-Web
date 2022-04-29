package models

import (
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/jinzhu/gorm"
)

type CalligraphyTab struct {
	ID         int `json:"id" gorm:"index"`
	PreImgId   int `json:"pre_img_id"`
	AfterImgId int `json:"after_img_id"`
	RemarkId   int `json:"remark_id"`
	IntroId    int `json:"intro_id"`
}

type CalligraphiesTab struct {
	ID            int    `json:"id" gorm:"index"`
	PreImgId      int    `json:"pre_img_id"`
	AfterImgId    int    `json:"after_img_id"`
	RemarkId      int    `json:"remark_id"`
	IntroId       int    `json:"intro_id"`
	PreUrl        string `json:"pre_url"`
	PreTime       int    `json:"pre_time"`
	AfterUrl      string `json:"after_url"`
	AfterTime     int    `json:"after_time"`
	RemarkContent int    `json:"remark_content"`
	RemarkScore   string `json:"remark_score"`
	RemarkDetail  string `json:"remark_detail"`
	IntroDesc     string `json:"intro_desc"`
}

func GetCalligraphys(pageNum, pageSize int, maps interface{}) ([]*CalligraphyTab, error) {

	var (
		calligraphys []*CalligraphyTab
		//preimages    []*PreImageTab
		//afterImages  []*AfterImageTab
		//remark		 []*RemarkTab
		//intro		 []*IntroTab
	)

	err := db.Offset(pageNum).Limit(pageSize).Find(&calligraphys).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err.Error())
		return nil, err
	}

	return calligraphys, nil
}

//获取其它列表的值
//func GetCallDetail(tab interface{})(maps interface{})  {
//	maps := make(map[string]interface{})
//
//}