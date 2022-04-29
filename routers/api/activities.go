package api

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/EDDYCJY/go-gin-example/service/activity_service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
	"strconv"
)

func GetActivities(c *gin.Context) {
	appG := app.Gin{C: c}

	activityService := activity_service.Activity{
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	activities, err := activityService.GetAll()
	if err != nil {
		log.Printf("[info] routers/api/activites %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ACTIVITIES_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = activities

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func QueryAct(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	//This is condition of the query
	var (
		catego   int
		actstart string
		actend   string
	)
	json := make(map[string]interface{})
	c.BindJSON(&json)
	catego, err := strconv.Atoi(fmt.Sprintf("%1.0f", json["actcatego"].(float64)))

	if catego != 0 {
		valid.Min(catego, 1, "actcatego")
	} else {
		catego = 0
	}

	if actstart = json["actstart"].(string); actstart != "" {
		valid.MaxSize(actstart, 10, "actstart")
	}

	if actend = json["actend"].(string); actend != "" {
		valid.MaxSize(actend, 10, "actend")
	}

	if valid.HasErrors() {
		log.Printf("[info] routers/api/activites/QueryAct  %s", "params valid error")
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	activityService := activity_service.Activity{
		ActStart:  actstart,
		ActEnd:    actend,
		ActCatego: catego,
		PageNum:   util.GetPage(c),
		PageSize:  setting.AppSetting.PageSize,
	}

	activities, err := activityService.QueryAll()

	if err != nil {
		log.Printf("[info] routers/api/activites %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ACTIVITIES_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = activities

	appG.Response(http.StatusOK, e.SUCCESS, data)

}

func GetActInfo(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Query("id")).MustInt()
	fmt.Println(id)
	valid := validation.Validation{}
	valid.Min(id, 1, "id")
	var (
		activityuser *models.ActivityUser
		err          error
	)
	if valid.HasErrors() {
		log.Printf("[info] routers/api/activites-GetActInfo")
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	activityService := activity_service.Activity{ID: id}
	exists, err := activityService.ExistById()

	if err != nil {
		log.Printf("[info] routers/api/activites %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ACTIVITY_FAIL, nil)
		return
	}

	if !exists {
		logging.Warn(exists)
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ACTIVITY, nil)
		return
	}

	activityuser, err = activityService.GetInfo()
	if err != nil {
		log.Printf("[info] routers/api/activites %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ACTIVITIES_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, activityuser)

}
