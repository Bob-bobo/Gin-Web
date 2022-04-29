package models

type AfterImageTab struct {
	ID        int    `json:"id" gorm:"index"`
	AftUserId int    `json:"aft_user_id"`
	AftUrl    string `json:"aft_url"`
	AftTime   int    `json:"aft_time"`
	AftType   string `json:"aft_type"`
	AftHeight string `json:"aft_height"`
	AftWidth  int    `json:"aft_width"`
}
