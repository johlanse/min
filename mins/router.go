package mins

import (
	"embed"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"min/controller"
	"min/docs"
)

var (
	token = "qqqq"
)

func init() {
	t, b := os.LookupEnv("token")
	if b {
		token = t
	}
}

//go:embed dist
var static embed.FS

func Rout() *gin.Engine {
	engine := gin.Default()

	engine.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "token"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))
	engine.Use(func(context *gin.Context) {
		header := context.GetHeader("token")
		if header != token && context.Request.Method == http.MethodPost {
			context.JSON(403, gin.H{"code": 1403, "err": "token 鉴权失败"})
		}
	})
	docs.SwaggerInfo.BasePath = "/"
	engine.StaticFS("/static", http.FS(static))

	engine.GET("/", func(context *gin.Context) {
		context.Redirect(301, "/static/dist/default.html")
	})

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
	admin.POST("/delete_user", controller.DeleteUser())

	// 任务相关
	active := engine.Group("/active")
	active.POST("/do_active", controller.DoActive())
	active.POST("/get_actives", controller.GetActives())
	active.POST("/stop_active", controller.StopActive())
	active.POST("/delete_active", controller.DeleteActive())

	engine.StaticFS("/log", gin.Dir("./logs/", false))

	return engine
}
