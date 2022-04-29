package models

import (
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/jinzhu/gorm"
)

type Role struct {
	Id   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

func QueryRole(id int) (string, error) {
	var role Role
	err := db.Where("id = ?", id).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn("QueryRole err: ", err)
		return "", err
	}
	return role.Name, nil
}

func QueryIdByName(name string) (int, error) {
	var role Role
	err := db.Where("name = ?", name).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn("QueryIdByName err :", err)
		return 0, err
	}

	return role.Id, nil
}
