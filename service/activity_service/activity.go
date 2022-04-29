package activity_service

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
)

type Activity struct {
	ID        int
	ActName   string
	ActStart  string
	ActEnd    string
	ActSite   string
	ActPers   int
	ActDetail string
	ActCatego int

	PageNum  int
	PageSize int
}

func (a *Activity) Add() error {
	activity := map[string]interface{}{
		"act_name":   a.ActName,
		"act_start":  a.ActStart,
		"act_end":    a.ActEnd,
		"act_pers":   a.ActPers,
		"act_site":   a.ActSite,
		"act_detail": a.ActDetail,
		"act_catego": a.ActCatego,
	}

	if err := models.AddActivity(activity); err != nil {
		logging.Warn(err.Error())
		return err
	}

	return nil
}

func (a *Activity) Join() {

}

func (a *Activity) Edit() error {
	return models.EditActivity(a.ID, map[string]interface{}{
		"act_name":   a.ActName,
		"act_start":  a.ActStart,
		"act_end":    a.ActEnd,
		"act_site":   a.ActSite,
		"act_pers":   a.ActPers,
		"act_detail": a.ActDetail,
		"act_catego": a.ActCatego,
	})
}

func (a *Activity) Get() (*models.ActivityTab, error) {
	var activity *models.ActivityTab

	activity, err := models.GetActivityOne(a.ID)
	if err != nil {
		logging.Warn(err.Error())
		return nil, err
	}

	return activity, nil
}

func (a *Activity) GetInfo() (*models.ActivityUser, error) {
	var activity *models.ActivityUser

	activity, err := models.GetActivityInfo(a.ID)
	if err != nil {
		logging.Warn(err.Error())
		return nil, err
	}
	return activity, nil
}

func (a *Activity) GetAll() ([]*models.ActivityTab, error) {
	var activities []*models.ActivityTab

	activities, err := models.GetActivitys(a.PageNum, a.PageSize, a.getMaps())

	if err != nil {
		logging.Warn(err.Error())
		return nil, err
	}

	return activities, nil
}

func (a *Activity) QueryAll() ([]*models.ActivityTab, error) {
	var activities []*models.ActivityTab
	activities, err := models.QueryAll(a.PageNum, a.PageSize, a.getMap())

	if err != nil {
		logging.Warn(err.Error())
		return nil, err
	}

	return activities, nil
}

func (a *Activity) Delete() error {
	return models.DeleteActivity(a.ID)
}

func (a *Activity) ExistById() (bool, error) {
	return models.ExistActivityById(a.ID)
}

//通过用户id查询活动
func (a *Activity) FindActById() ([]*models.ActivityTab, error) {

	var activities []*models.ActivityTab

	activities, err := models.GetActivitiesById(a.ID)

	if err != nil {
		logging.Warn(err.Error())
		return nil, err
	}

	return activities, nil
}

func (a *Activity) Count() (int, error) {
	return models.GetActivityTotal(a.getMaps())
}

func (a *Activity) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	maps["deleted_on"] = 0
	maps["created_on"] = 0
	maps["modified_on"] = 0

	return maps
}

func (a *Activity) getMap() map[string]interface{} {
	maps := make(map[string]interface{})

	maps["act_catego"] = a.ActCatego
	maps["act_start"] = a.ActStart
	maps["act_end"] = a.ActEnd

	return maps
}
