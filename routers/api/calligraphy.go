package api

import (
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/EDDYCJY/go-gin-example/service/calligraphy_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllCalligraphy(c *gin.Context) {

	appG := app.Gin{C: c}

	calligraphyService := calligraphy_service.Calligraphy{
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	calligraphys, err := calligraphyService.GetAll()
	if err != nil {
		logging.Warn(err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_CALLIGRAPHY_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = calligraphys

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
