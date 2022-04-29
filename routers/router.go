package routers

import (
	"github.com/EDDYCJY/go-gin-example/middleware/jwt"
	"github.com/gin-gonic/gin"

	_ "github.com/EDDYCJY/go-gin-example/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/EDDYCJY/go-gin-example/routers/api"
	"github.com/EDDYCJY/go-gin-example/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//针对非登陆用户
	//查看所有活动
	r.GET("/activities", api.GetActivities)
	//搜索活动
	r.POST("/query", api.QueryAct)
	//活动详情页
	r.GET("/actInfo", api.GetActInfo)
	//查看所有界面
	r.GET("/findAll", api.GetAllCalligraphy)

	//登陆操作
	r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{

		//针对管理员
		//查看活动
		apiv1.GET("/activity", v1.GetActivity)
		//添加活动
		apiv1.POST("/activity", v1.AddActivity)
		//编辑活动
		apiv1.PUT("/activity", v1.EditActivity)
		//删除活动
		apiv1.DELETE("/activity", v1.DeleteActivity)
		//查看分类
		apiv1.GET("/category", v1.GetCategories)
		//添加分类
		apiv1.POST("/category", v1.AddCategory)
		//编辑分类
		apiv1.PUT("/category", v1.EditCategory)
		//删除分类
		apiv1.DELETE("/category", v1.DeleteCategory)
		//查询所有用户信息
		apiv1.GET("/alluser", v1.FindAllUser)

		//针对登陆用户
		//在活动详情页发表评论
		apiv1.POST("/comment", v1.AddComment)
		//参加活动
		apiv1.POST("/joinAct", v1.JoinActivity)
		//退出活动
		apiv1.POST("/exitAct", v1.ExitActivity)
		//查看参加的活动
		apiv1.GET("/findAct", v1.FindJoinAct)
		//查看个人信息
		apiv1.GET("/selfInfo", v1.GetUserInfo)
	}

	return r
}
