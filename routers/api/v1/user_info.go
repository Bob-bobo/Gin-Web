package v1

import (
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/EDDYCJY/go-gin-example/service/user_service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type GetUserForm struct {
	UserName string `form:"username" valid:"Required;MaxSize(50)"`
	PassWord string `form:"password" valid:"Required;MaxSize(50)"`
}

func GetUserInfo(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form GetUserForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		log.Printf("[info] routers/api/v1/user_info-GetUserInfo %d", errCode)
		logging.Error(errCode)
		appG.Response(httpCode, errCode, nil)
		return
	}

	userService := user_service.User{
		UserName: form.UserName,
		PassWord: form.PassWord,
	}
	user, err := userService.Get()
	if err != nil {
		log.Printf("[info] routers/api/v1/user_info-GetUserInfo %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_INFO_FALI, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, user)
}

func FindAllUser(c *gin.Context) {
	appG := app.Gin{C: c}

	userService := user_service.User{
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	users, err := userService.GetAll()
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USER_FALI, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = users

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
