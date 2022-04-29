package models

type IntroTab struct {
	ID        int    `json:"id" gorm:"index"`
	IntroDesc string `json:"intro_desc"`
}
