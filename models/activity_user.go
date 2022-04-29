package models

import "github.com/EDDYCJY/go-gin-example/pkg/logging"

type ActivityUserTab struct {
	UserId     int `json:"user_id"`
	ActivityId int `json:"activity_id"`
}

// AddActUser insert data to this table
func AddActUser(data map[string]interface{}) error {
	actUser := ActivityUserTab{
		UserId:     data["user_id"].(int),
		ActivityId: data["activity_id"].(int),
	}
	if err := db.Create(&actUser).Error; err != nil {
		logging.Warn(err)
		return err
	}

	return nil
}

func QueryActUser(userid, actid int) bool {

	var act ActivityUserTab
	err := db.Where("user_id = ? and activity_id = ?", userid, actid).First(&act).Error
	if err != nil {
		logging.Warn(err)
		return false
	}
	return true
}

func ExitActUser(userid, actid int) error {
	if err := db.Where("user_id = ? and activity_id = ?", userid, actid).Delete(ActivityUserTab{}).Error; err != nil {
		return err
	}

	return nil
}
