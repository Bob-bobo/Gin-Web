package models

import (
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/jinzhu/gorm"
)

type UserTab struct {
	ID             int    `gorm:"primary_key" json:"id"`
	Phone          string `gorm:""json:"phone"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Gender         string `json:"gender"`
	Truename       string `json:"trueName"`
	Birthday       string `json:"birthday"`
	Email          string `json:"email"`
	Personalbrief  string `json:"personalBrief"`
	Avatarimgurl   string `json:"avatarImgUrl"`
	Recentlylanded string `json:"recentlyLanded"`
}

func CheckUser(username, password string) (bool, error) {
	var user UserTab
	//err := db.Select("id").Where(User{Username: username, Password: passwords}).First(&user).Error
	err := db.Where("username = ?", username).Limit(1).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn("CheckUser err:", err)
		return false, err
	}
	if !util.ComparePasswords(user.Password, []byte(password)) {
		logging.Warn("ComparePasswords err: ", err)
		return false, err
	}

	if user.ID > 0 {
		logging.Warn("UserId < 0 err: ", err)
		return true, nil
	}
	return false, nil
}

func GetUserInfo(username string) (*UserTab, error) {
	var user UserTab
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn("GetUserInfo err: ", err)
		return nil, err
	}

	return &user, nil
}

func GetAllUser(pageNum, PageSize int) ([]*UserTab, error) {
	var users []*UserTab
	db.Raw("select username,email,avatarimgurl from user_tab").Offset(pageNum).Limit(PageSize).Scan(&users)

	return users, nil
}
