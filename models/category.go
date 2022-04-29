package models

import (
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/jinzhu/gorm"
)

type CategoryTab struct {
	ID           int    `gorm:"primary_key" json:"id"`
	CategoryName string `json:"category_name"`
}

//通过name
func AddCategory(data map[string]interface{}) error {
	category := CategoryTab{
		CategoryName: data["category_name"].(string),
	}
	if err := db.Create(&category).Error; err != nil {
		logging.Warn(err)
		return err
	}

	return nil
}

func QueryCategory(cateId int) (string, error) {
	var category CategoryTab
	err := db.Where("id = ?", cateId).First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return "", err
	}
	return category.CategoryName, nil
}

func ExistCateByID(id int) (bool, error) {
	var cate CategoryTab
	err := db.Select("id").Where("id = ?", id).First(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return false, err
	}

	if cate.ID > 0 {
		logging.Warn(err)
		return true, nil
	}
	return false, nil
}

func GetCategoryTotal(maps interface{}) (int, error) {

	var count int
	if err := db.Model(&CategoryTab{}).Where(maps).Count(&count).Error; err != nil {
		logging.Warn(err)
		return 0, err
	}
	return count, nil
}

func GetCategories(maps interface{}) ([]*CategoryTab, error) {
	var categories []*CategoryTab
	err := db.Where(maps).Find(&categories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Warn(err)
		return nil, err
	}
	return categories, nil
}

func EditCategory(id int, data interface{}) error {
	if err := db.Model(&CategoryTab{}).Where("id = ?", id).Updates(data).Error; err != nil {
		logging.Warn(err)
		return err
	}

	return nil
}

func DeleteCategory(id int) error {
	if err := db.Where("id = ?", id).Delete(CategoryTab{}).Error; err != nil {
		logging.Warn(err)
		return err
	}
	return nil
}
