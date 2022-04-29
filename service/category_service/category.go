package category_service

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
)

type Category struct {
	ID           int
	CategoryName string
}

func (c *Category) ExistByID() (bool, error) {
	return models.ExistCateByID(c.ID)
}

func (c *Category) Count() (int, error) {
	return models.GetCategoryTotal(c.getMaps())
}

func (c *Category) GetAll() ([]*models.CategoryTab, error) {
	var categories []*models.CategoryTab

	categories, err := models.GetCategories(c.getMaps())
	if err != nil {
		logging.Warn(err.Error())
		return nil, err
	}
	return categories, nil
}

func (c *Category) Add() error {
	category := map[string]interface{}{
		"category_name": c.CategoryName,
	}

	if err := models.AddCategory(category); err != nil {
		logging.Warn(err.Error())
		return err
	}
	return nil
}

func (c *Category) Edit() error {
	return models.EditCategory(c.ID, map[string]interface{}{
		"category_name": c.CategoryName,
	})
}

func (c *Category) Delete() error {
	return models.DeleteCategory(c.ID)
}

func (c *Category) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	return maps
}
