package mins

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"min/controller"
	"min/docs"
)

func Rout() *gin.Engine {
	engine := gin.Default()

	engine.Use(cors.Default())
	docs.SwaggerInfo.BasePath = "/"
	// 明华在线相关管理
	engine.GET("/swag/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	group := engine.Group("/min")
	group.POST("/login", controller.Login())
	group.POST("/logician", controller.LoginWeiXinController())
	group.POST("/info", controller.GetInfo())
	group.POST("/check_qr_code", controller.CheckQrCode())
	group.POST("/get_courses", controller.GetCourses())
	group.POST("/get_course_list", controller.GetCourseList())

	// 用户分组
	admin := engine.Group("/admin")
	admin.POST("/get_login_user", controller.GetLoginUser())
	admin.POST("/get_all_user", controller.GetAllUser())

	// 任务相关
	active := engine.Group("/active")
	active.POST("/do_active", controller.DoActive())
	active.POST("/get_actives", controller.GetActives())
	active.POST("/stop_active", controller.StopActive())
	active.POST("/delete_active", controller.DeleteActive())

	engine.StaticFS("/log", gin.Dir("./logs/", false))

	return engine
}
