package models

type Category struct {
	ID       int    `gorm:"primary_key" json:"id"`
	CateName string `json:"catename"`
}

//通过name
