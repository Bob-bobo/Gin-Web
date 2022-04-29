package v1

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/service/activity_service"
	"github.com/EDDYCJY/go-gin-example/service/category_service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

func GetActivity(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		log.Printf("[info] routers/api/v1/activity-GetActivity %s", "INVALID_PARAMS")
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	activityService := activity_service.Activity{ID: id}
	exists, err := activityService.ExistById()

	if err != nil {
		log.Printf("[info] routers/api/v1/activity-GetActivity %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ACTIVITY_FAIL, nil)
		return
	}

	if !exists {
		logging.Info("not exist")
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ACTIVITY, nil)
		return
	}
	activity, err := activityService.Get()
	if err != nil {
		log.Printf("[info] routers/api/v1/activity-GetActivity %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ACTIVITY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, activity)
}

type AddActivityForm struct {
	ActName   string `form:"act_name" valid:"Required;MaxSize(100)"`
	ActStart  string `form:"act_start" valid:"Required;MaxSize(20)"`
	ActEnd    string `form:"act_end" valid:"Required;MaxSize(20)"`
	ActSite   string `form:"act_site" valid:"Required;MaxSize(50)"`
	ActPers   int    `form:"act_pers" valid:"Max(1000)"`
	ActDetail string `form:"act_detail" valid:"Required;MaxSize(255)"`
	ActCatego int    `form:"act_catego" valid:"Range(1,5)"`
}

func AddActivity(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddActivityForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		log.Printf("[info] routers/api/v1/activity-AddActivity %d", errCode)
		logging.Error(errCode)
		appG.Response(httpCode, errCode, nil)
		return
	}
	cateService := category_service.Category{ID: form.ActCatego}
	exists, err := cateService.ExistByID()
	if err != nil {
		log.Printf("[info] routers/api/v1/activity-AddActivity %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_CATEGO, nil)
		return
	}

	if !exists {
		logging.Info("not exist")
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_CATEGO, nil)
		return
	}

	activityService := activity_service.Activity{
		ActName:   form.ActName,
		ActStart:  form.ActStart,
		ActEnd:    form.ActEnd,
		ActSite:   form.ActSite,
		ActPers:   form.ActPers,
		ActDetail: form.ActDetail,
		ActCatego: form.ActCatego,
	}
	if err := activityService.Add(); err != nil {
		log.Printf("[info] routers/api/v1/activity-AddActivity %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ACTIVITY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditActivityForm struct {
	ID        int    `form:"id" valid:"Required;Min(1)"`
	ActName   string `form:"act_name" valid:"Required;MaxSize(100)"`
	ActStart  string `form:"act_start" valid:"Required;MaxSize(20)"`
	ActEnd    string `form:"act_end" valid:"Required;MaxSize(20)"`
	ActSite   string `form:"act_site" valid:"Required;MaxSize(50)"`
	ActPers   int    `form:"act_pers" valid:"Max(1000)"`
	ActDetail string `form:"act_detail" valid:"Required;MaxSize(255)"`
	ActCatego int    `form:"act_catego" valid:"Range(1,5)"`
}

func EditActivity(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form EditActivityForm //{ID: com.StrTo(c.Query("id")).MustInt()}
	)
	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		log.Printf("[info] routers/api/v1/activity-EditActivity %d", errCode)
		logging.Error(errCode)
		appG.Response(httpCode, errCode, nil)
		return
	}

	activityService := activity_service.Activity{
		ID:        form.ID,
		ActName:   form.ActName,
		ActStart:  form.ActStart,
		ActEnd:    form.ActEnd,
		ActSite:   form.ActSite,
		ActPers:   form.ActPers,
		ActDetail: form.ActDetail,
		ActCatego: form.ActCatego,
	}

	exists, err := activityService.ExistById()
	if err != nil {
		log.Printf("[info] routers/api/v1/activity-EditActivity %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ACTIVITY_FAIL, nil)
		return
	}
	if !exists {
		logging.Info("not exist")
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ACTIVITY, nil)
		return
	}

	categoryService := category_service.Category{ID: form.ActCatego}
	exists, err = categoryService.ExistByID()
	if err != nil {
		log.Printf("[info] routers/api/v1/activity-EditActivity %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_CATEGO_FAIL, nil)
		return
	}

	if !exists {
		logging.Info("not exist")
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_CATEGO, nil)
		return
	}

	err = activityService.Edit()
	if err != nil {
		log.Printf("[info] routers/api/v1/activity-EditActivity %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_ACTIVITY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteActivity(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Query("id")).MustInt()
	valid.Min(id, 1, "id").Message("必须大于0")

	if valid.HasErrors() {
		log.Printf("[info] routers/api/v1/activity-DeleteActivity %s", "INVALID_PARAMS")
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	activityService := activity_service.Activity{ID: id}
	exists, err := activityService.ExistById()

	if err != nil {
		log.Printf("[info] routers/api/v1/activity-DeleteActivity %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ACTIVITY_FAIL, nil)
		return
	}

	if !exists {
		logging.Info("not exist")
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ACTIVITY, nil)
		return
	}

	err = activityService.Delete()
	if err != nil {
		log.Printf("[info] routers/api/v1/activity-DeleteActivity %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_ACTIVITY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func FindJoinAct(c *gin.Context) {

	var (
		activities []*models.ActivityTab
		err        error
	)

	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Query("id")).MustInt()
	valid.Min(id, 1, "id").Message("必须大于0")

	if valid.HasErrors() {
		log.Printf("[info] routers/api/v1/activity-FindJoinAct %s", "INVALID_PARAMS")
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	activityervice := activity_service.Activity{ID: id}

	activities, err = activityervice.FindActById()

	if err != nil {
		log.Printf("[info] routers/api/v1/activity-FindJoinAct %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ACTIVITY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, activities)
}
