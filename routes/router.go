package routes

import (
	"TARG_revenue_report_backend/middleware"
	"TARG_revenue_report_backend/sso"
	"TARG_revenue_report_backend/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	//	r.Use(middleware.Cors())

	r.GET("/sso", sso.SsoLogin)

	router := r.Group("api/v1")

	router.Use(middleware.JWT())
	{

		// 测试接口
		//router.GET("hello", func(c *gin.Context) {
		//	c.JSON(http.StatusOK, gin.H{
		//		"msg": "ok",
		//	})
		//})

		router.GET("/hello", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "hello world",
			})
		})

		//router.POST("/user/register", v1.UserRegister)
		//
		//router.POST("/user/login", v1.UserLogin)
		//
		//// 用户模块的路由接口
		//router.POST("/user/add", v1.AddUser)
		//router.GET("/users", v1.GetUsers)
		//router.PUT("/user/editUser", v1.EditUser)
		//router.DELETE("/user/:id", v1.DeleteUser)

		// 分类模块的路由接口

		// 任务模块的路由接口

	}

	r.Run(utils.HttpPort)
}
