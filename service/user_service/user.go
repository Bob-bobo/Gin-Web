package user_service

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
)

type User struct {
	ID             int
	UserName       string
	PassWord       string
	Phone          string
	Gender         string
	TrueName       string
	Birthday       string
	Email          string
	PersonalBrief  string
	AvatarImgUrl   string
	RecentlyLanded string

	PageNum  int
	PageSize int
}

func (u *User) Get() (*models.UserTab, error) {
	var user *models.UserTab

	user, err := models.GetUserInfo(u.UserName)
	if err != nil {
		logging.Warn(err.Error())
		return nil, err
	}
	return user, nil
}

func (u *User) GetAll() ([]*models.UserTab, error) {
	var users []*models.UserTab

	users, err := models.GetAllUser(u.PageNum, u.PageSize)
	if err != nil {
		logging.Warn(err.Error())
		return nil, err
	}
	return users, nil
}

//func (u *User) getMap() map[string]interface{} {
//	maps := make(map[string]interface{})
//
//	maps["username"] = u.UserName
//	maps["email"] = u.Email
//	maps["avatarImgUrl"] = u.AvatarImgUrl
//
//	return maps
//}
