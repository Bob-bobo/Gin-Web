package v1

import (
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/service/comment_service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AddCommentForm struct {
	UserID     int    `form:"user_id" valid:"Required;Min(1)"`
	ActID      int    `form:"activity_id" valid:"Required;Min(1)"`
	CommentCon string `form:"comment_content" valid:"Required;MaxSize(255)";`
}

func AddComment(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddCommentForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		log.Printf("[info] routers/api/v1/comment-AddComment %d", errCode)
		logging.Error(errCode)
		appG.Response(httpCode, errCode, nil)
		return
	}

	commentService := comment_service.Comment{
		UserID:         form.UserID,
		ActivityID:     form.ActID,
		CommentContent: form.CommentCon,
	}
	if err := commentService.Add(); err != nil {
		log.Printf("[info] routers/api/v1/comment-AddComment %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_COMMENT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
