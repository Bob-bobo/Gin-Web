package v1

import (
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/service/actuser_service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

type ActivityForm struct {
	UserID     int `form:"user_id" valid:"Required;Min(1)"`
	ActivityID int `form:"activity_id" valid:"Required;Min(1)"`
}

func JoinActivity(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form ActivityForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		log.Printf("[info] routers/api/v1/user_activity-JoinActivity %d", errCode)
		logging.Error(errCode)
		appG.Response(httpCode, errCode, nil)
		return
	}

	actuserService := actuser_service.ActUser{
		UserID:     form.UserID,
		ActivityID: form.ActivityID,
	}

	if err := actuserService.Add(); err != nil {
		log.Printf("[info] routers/api/v1/user_activity-JoinActivity %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_JOIN_ACTIVITY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

func ExitActivity(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	userid := com.StrTo(c.Query("userId")).MustInt()
	actid := com.StrTo(c.Query("actId")).MustInt()
	valid.Min(userid, 1, "userId").Message("必须大于0")
	valid.Min(actid, 1, "actId").Message("必须大于0")

	if valid.HasErrors() {
		log.Printf("[info] routers/api/v1/user_activity-EditActivity %s", "INVALID_PARAMS")
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	actuserService := actuser_service.ActUser{UserID: userid, ActivityID: actid}
	exists := actuserService.ExistById()

	if !exists {
		logging.Info("not exist")
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ACTUSER, nil)
		return
	}

	err := actuserService.DeleteById()
	if err != nil {
		log.Printf("[info] routers/api/v1/user_activity-EditActivity %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_ACTUSER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
