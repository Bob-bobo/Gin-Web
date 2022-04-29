package models

type PreImageTab struct {
	ID        int    `json:"id" gorm:"index"`
	PreUserId int    `json:"pre_user_id"`
	PreUrl    string `json:"pre_url"`
	PreTime   int    `json:"pre_time"`
	PreType   string `json:"pre_type"`
	PreHeight int    `json:"pre_height"`
	PreWidth  int    `json:"pre_width"`
}
