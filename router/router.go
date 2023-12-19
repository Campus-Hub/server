package router

import (
	"github.com/Campus-Hub/server/api"

	"github.com/gin-gonic/gin"
)

// Setup 路由配置
func Setup() *gin.Engine {
	r := gin.Default()
	//store := cookie.NewStore([]byte("something-very-secret"))
	//
	//// middleware
	//r.Use(sessions.Sessions("mySession", store))
	////r.Use(middleware.NewLogger(), middleware.Cors())
	//
	//r.StaticFS("/static", http.Dir("./static"))

	/**********  Resource Management  **********/
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, "v1 success")
		})

		//// Resource management router version 1
		//resourcev1 := v1.Group("/resource")
		//{
		//	resourcev1.POST("/upload")
		//}

		// Courses Management router version 1
		coursev1 := v1.Group("/course")
		{
			// 查询课程清单
			coursev1.GET("/", api.ListCourse)
			// 创建课程
			coursev1.POST("/")
			// 通过ID？Name？获取课程信息
			coursev1.GET("/:id", api.ShowCourse)
			//coursev1.GET("/:name", api.ShowCourse)

			authcoursev1 := coursev1.Group("/auth")
			{
				// 更新课程
				authcoursev1.PUT("/:name")
				// 删除课程
				authcoursev1.DELETE("/:name")
			}
		}
	}

	return r
}
