package actuser_service

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
)

type ActUser struct {
	UserID     int
	ActivityID int
}

func (a *ActUser) Add() error {
	actuser := map[string]interface{}{
		"user_id":     a.UserID,
		"activity_id": a.ActivityID,
	}

	if err := models.AddActUser(actuser); err != nil {
		logging.Warn(err.Error())
		return err
	}
	return nil
}

func (a *ActUser) ExistById() bool {
	return models.QueryActUser(a.UserID, a.ActivityID)
}

func (a *ActUser) DeleteById() error {
	return models.ExitActUser(a.UserID, a.ActivityID)
}
