package models

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/jinzhu/gorm"
)

type ActivityTab struct {
	ID        int `json:"id" gorm:"index"`
	ActCatego int `json:"act_catego"`

	ActName   string `json:"act_name"`
	ActStart  string `json:"act_start"`
	ActEnd    string `json:"act_end"`
	ActSite   string `json:"act_site"`
	ActPers   int    `json:"act_pers"`
	ActDetail string `json:"act_detail"`

	Model
}

//ExistActivityById find the activity
func ExistActivityById(id int) (bool, error) {
	var activity ActivityTab
	err := db.Where("id = ?", id).First(&activity).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return false, err
	}

	if activity.ID > 0 {
		logging.Warn(err)
		return true, nil
	}

	return false, nil
}

//GetActivityTotal gets the total number of activity based on the function
func GetActivityTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Select("id").Model(&ActivityTab{}).Where(maps).Count(&count).Error; err != nil {
		logging.Warn(err)
		return 0, err
	}
	return count, nil
}

//GetActivitys gets the all of the activity based on the way
func GetActivitys(pageNum int, pageSize int, maps interface{}) ([]*ActivityTab, error) {
	var activitys []*ActivityTab
	err := db.Model(maps).Offset(pageNum).Limit(pageSize).Find(&activitys).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return nil, err
	}
	return activitys, nil
}

func QueryAll(pageNum, pageSize int, maps map[string]interface{}) ([]*ActivityTab, error) {
	var activities []*ActivityTab
	sql := ""
	if actCatego := maps["act_catego"]; actCatego != 0 {
		sql = fmt.Sprintf("act_catego = %d", actCatego)
	}
	if actStart := maps["act_start"]; actStart != "" {
		if sql == "" {
			sql = fmt.Sprintf("str_to_date(act_start,%s) > str_to_date( '%s',%s)", "'%Y-%m-%d'", actStart, "'%Y-%m-%d'")
		} else {
			sql = fmt.Sprintf("%s and str_to_date(act_start,%s) > str_to_date( '%s',%s)", sql, "'%Y-%m-%d'", actStart, "'%Y-%m-%d'")
		}
	}
	if actEnd := maps["act_end"]; actEnd != "" {
		if sql == "" {
			sql = fmt.Sprintf("str_to_date(act_end,%s) > str_to_date( '%s',%s)", "'%Y-%m-%d'", actEnd, "'%Y-%m-%d'")
		} else {
			sql = fmt.Sprintf("%s and str_to_date(act_end,%s) > str_to_date( '%s',%s)", sql, "'%Y-%m-%d'", actEnd, "'%Y-%m-%d'")
		}
	}
	sql = fmt.Sprintf("select act_name,act_start,act_end,act_catego from activity_tab where %s", sql)

	db.Select("act_name", "act_start", "act_end").Raw(sql).Offset(pageNum).Limit(pageSize).Scan(&activities)

	return activities, nil
}

//Get one of the activitys what is you want

func GetActivityOne(id int) (*ActivityTab, error) {
	var activity ActivityTab
	err := db.Where("id = ?", id).First(&activity).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return nil, err
	}
	/**
	关联查询先不加
	err = db.Model(&activity).Related(&activity.ActCatego).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil,err
	}
	*/

	return &activity, nil
}

type ActivityUser struct {
	ActName      string
	ActStart     string
	ActEnd       string
	ActDetail    string
	ActSite      string
	UserName     string
	AvatarImgUrl string
}

func GetActivityInfo(id int) (*ActivityUser, error) {
	var activityUser ActivityUser
	db.Raw("select act_name,act_start,act_end,act_site,username,avatarImgUrl,comment_content from"+
		"((activity_tab a1 left join activity_user_tab a2 "+
		"on a1.id = a2.activity_id) left join user_tab a3 "+
		"on a3.id = a2.user_id) left join comment_tab a4 "+
		"on a1.id = a4.activity_id where a1.id = ?", id).Limit(20).Scan(&activityUser)

	return &activityUser, nil
}
func EditActivity(id int, data interface{}) error {
	if err := db.Model(&ActivityTab{}).Where("id = ?", id).Updates(data).Error; err != nil {
		logging.Warn(err)
		return err
	}

	return nil
}

func AddActivity(data map[string]interface{}) error {
	activity := ActivityTab{
		ActName:   data["act_name"].(string),
		ActStart:  data["act_start"].(string),
		ActEnd:    data["act_end"].(string),
		ActSite:   data["act_site"].(string),
		ActPers:   data["act_pers"].(int),
		ActDetail: data["act_detail"].(string),
		ActCatego: data["act_catego"].(int),
	}
	if err := db.Create(&activity).Error; err != nil {
		logging.Warn(err)
		return err
	}
	return nil
}

//delete a single activity

func DeleteActivity(id int) error {
	fmt.Println("-------id = ", id)
	if err := db.Unscoped().Where("id = ?", id).Delete(ActivityTab{}).Error; err != nil {
		logging.Warn(err)
		return err
	}
	return nil
}

func GetActivitiesById(id int) ([]*ActivityTab, error) {
	var activity []*ActivityTab
	db.Raw("select act_name,act_start,act_end,act_site from"+
		" activity_user_tab a1 left join activity_tab a2 on"+
		" a1.activity_id = a2.id where a1.user_id = ?", id).Limit(20).Scan(&activity)
	return activity, nil
}
