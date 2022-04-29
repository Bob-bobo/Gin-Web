package v1

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/service/category_service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

//type Category struct {
//	"id":
//}

func GetCategories(c *gin.Context) {
	appG := app.Gin{C: c}

	category := models.CategoryTab{}

	categoryService := category_service.Category{
		ID:           category.ID,
		CategoryName: category.CategoryName,
	}

	total, err := categoryService.Count()
	if err != nil {
		log.Printf("[info] routers/api/v1/category-GetCategories %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_CATEGO_FAIL, nil)
		return
	}
	categories, err := categoryService.GetAll()
	if err != nil {
		log.Printf("[info] routers/api/v1/category-GetCategories %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_CATEGOS_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = categories
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddCategoryForm struct {
	CategoryName string `form:"category_name" valid:"Required;MaxSize(100)"`
}

func AddCategory(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddCategoryForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		log.Printf("[info] routers/api/v1/category-AddCategory %d", errCode)
		logging.Error(errCode)
		appG.Response(httpCode, errCode, nil)
		return
	}

	categoryService := category_service.Category{CategoryName: form.CategoryName}

	if err := categoryService.Add(); err != nil {
		log.Printf("[info] routers/api/v1/category-AddCategory %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_CATEGO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditCategoryForm struct {
	ID           int    `form:"id" valid:"Required;Min(1)"`
	CategoryName string `form:"category_name" valid:"Required;MaxSize(100)"`
}

func EditCategory(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form EditCategoryForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		log.Printf("[info] routers/api/v1/category-EditCategory %d", errCode)
		logging.Error(errCode)
		appG.Response(httpCode, errCode, nil)
		return
	}

	categoryService := category_service.Category{
		ID:           form.ID,
		CategoryName: form.CategoryName,
	}

	exists, err := categoryService.ExistByID()
	if err != nil {
		log.Printf("[info] routers/api/v1/category-EditCategory %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_CATEGORY_FAIL, nil)
		return
	}
	if !exists {
		logging.Info("not exist")
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_CATEGO, nil)
		return
	}

	err = categoryService.Edit()
	if err != nil {
		log.Printf("[info] routers/api/v1/category-EditCategory %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_CATEGO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteCategory(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Query("id")).MustInt()
	valid.Min(id, 1, "id").Message("id必须大于0")

	if valid.HasErrors() {
		log.Printf("[info] routers/api/v1/category-DeleteCategory %s", "INVALID_PARAMS")
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	categoryService := category_service.Category{ID: id}
	exists, err := categoryService.ExistByID()
	if err != nil {
		log.Printf("[info] routers/api/v1/category-DeleteCategory %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_CATEGORY_FAIL, nil)
		return
	}
	if !exists {
		logging.Info("not exist")
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_CATEGO, nil)
		return
	}

	err = categoryService.Delete()
	if err != nil {
		log.Printf("[info] routers/api/v1/category-DeleteCategory %s", err.Error())
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_CATEGO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
