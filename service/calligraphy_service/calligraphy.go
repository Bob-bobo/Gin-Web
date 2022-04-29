package calligraphy_service

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
)

type Calligraphy struct {
	ID         int
	PreImgId   int
	AfterImgId int
	RemarkId   int
	IntroId    int

	PageNum  int
	PageSize int
}

func (c *Calligraphy) GetAll() ([]*models.CalligraphyTab, error) {

	var calligraphys []*models.CalligraphyTab

	calligraphys, err := models.GetCalligraphys(c.PageNum, c.PageSize, c.getMaps())

	if err != nil {
		logging.Warn(err.Error())
		return nil, err
	}

	return calligraphys, nil
}

func (c *Calligraphy) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	maps["intro_id"] = c.ID

	//maps["pre_url"] = c.PreUrl
	//maps["pre_time"] = c.PreTime
	//maps["after_url"] = c.AfterUrl
	//maps["after_time"] = c.AfterTime
	//maps["remark_content"] = c.RemarkContent
	//maps["remark_score"] = c.RemarkScore
	//maps["remark_detail"] = c.RemarkDetail
	//maps["intro_desc"] = c.IntroDesc

	return maps
}
