package controller

import (
	"github.com/gin-gonic/gin"

	"min/model"
)

// GetAllUser
// @Description: 获取所有已创建的用户
// @return gin.HandlerFunc
// example
// @Summary 获取所有已创建的用户
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} Mess{data=[]model.Min}
// @Failure 403 {object} Mess
// @Router /admin/get_all_user [POST]
func GetAllUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		mins, err := model.Query("1=1")
		if err != nil {
			context.JSON(401, Mess{
				Code: 1401,
				Msg:  "",
				Err:  err.Error(),
				Data: nil,
			})
			return
		}
		context.JSON(200, Mess{
			Code: 1200,
			Msg:  "操作成功",
			Err:  "",
			Data: mins,
		})
	}
}

// GetLoginUser
// @Description: 获取所有已登录的用户
// @return gin.HandlerFunc
// example
// @Summary 获取所有已登录的用户
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} Mess{data=[]model.Min}
// @Failure 403 {object} Mess
// @Router /admin/get_login_user [POST]
func GetLoginUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		mins, err := model.Query("state=1")
		if err != nil {
			context.JSON(401, Mess{
				Code: 1401,
				Msg:  "",
				Err:  err.Error(),
				Data: nil,
			})
			return
		}
		context.JSON(200, Mess{
			Code: 1200,
			Msg:  "操作成功",
			Err:  "",
			Data: mins,
		})
	}
}
