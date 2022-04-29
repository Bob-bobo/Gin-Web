package comment_service

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
)

type Comment struct {
	ID             int
	UserID         int
	ActivityID     int
	CommentContent string
}

func (c *Comment) Add() error {
	comment := map[string]interface{}{
		"user_id":         c.UserID,
		"activity_id":     c.ActivityID,
		"comment_content": c.CommentContent,
	}
	if err := models.AddComment(comment); err != nil {
		logging.Warn(err.Error())
		return err
	}
	return nil
}
