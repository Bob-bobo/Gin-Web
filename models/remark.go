package models

type RemarkTab struct {
	ID            int    `json:"id" gorm:"index"`
	RemarkContent string `json:"remark_content"`
	RemarkScore   int    `json:"remark_score"`
	RemarkDetail  string `json:"remark_detail"`
}
