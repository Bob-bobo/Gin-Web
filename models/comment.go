package models

import "github.com/EDDYCJY/go-gin-example/pkg/logging"

type CommentTab struct {
	Model

	UserId         int    `json:"user_id"`
	ActivityId     int    `json:"activity_id"`
	CommentContent string `json:"comment_content"`
}

func AddComment(data map[string]interface{}) error {
	comment := CommentTab{
		UserId:         data["user_id"].(int),
		ActivityId:     data["activity_id"].(int),
		CommentContent: data["comment_content"].(string),
	}
	if err := db.Create(&comment).Error; err != nil {
		logging.Warn(err)
		return err
	}
	return nil
}
