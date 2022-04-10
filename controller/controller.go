package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"min/config"
	"min/lib"
	"min/model"
)

// Login
// 登录明华账号
// @Summary 登录明华账号
// @Description 登录账号
// @Tags min
// @Accept  json
// @Produce  json
// @Param  user body config.User true "账号密码"
// @Success 200 {object} lib.Response
// @Failure 403 {object} Mess
// @Router /min/login [POST]
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user map[string]string
		err := ctx.BindJSON(&user)
		if err != nil {
			ctx.JSON(403, Mess{
				Code: 1403,
				Msg:  "获取数据失败",
				Err:  err.Error(),
				Data: nil,
			})
			return
		}
		status := user["status"]
		account, ok := user["account"]
		password, ok := user["password"]
		sta, _ := strconv.Atoi(status)
		var base string
		if sta == 0 {
			base = "https://mooc.cdcas.com"
		} else if sta == 1 {
			base = "https://shixun.cdcas.com"
		}
		if ok {
			login := lib.Login(account, password, base)
			if login.Status {
				min := model.Min{}
				for _, cookie := range login.Cookies {
					if cookie.Name == "token" {
						min.Token = cookie.Value
					} else if cookie.Name == "tgw_l7_route" {
						min.Tgw = cookie.Value
					}
					min.Account = user["account"]
					min.Password = user["password"]
					min.State = sta
				}
				id, err := model.Add(min)
				if err != nil {
					ctx.JSON(403, Mess{
						Code: 1403,
						Msg:  "获取数据失败",
						Err:  err.Error(),
						Data: nil,
					})
					return
				}
				ctx.JSON(200, Mess{
					Code: 1200,
					Msg:  "操作成功",
					Err:  "",
					Data: gin.H{"data": login, "id": id},
				})
				return
			}
			ctx.JSON(200, Mess{
				Code: 1200,
				Msg:  "操作成功",
				Err:  "",
				Data: gin.H{"data": login},
			})
		}
	}
}

// LoginWeiXinController
// @Description: 登录微信授权，获取二维码
// @return gin.HandlerFunc
// example
// @Summary 登录微信授权，获取二维码
// @Tags min
// @Accept json
// @Produce json
// @Param id query int true "ID"
// @Success 200 {object} Mess
// @Failure 403 {object} Mess
// @Router /min/logician [POST]
func LoginWeiXinController() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, b := context.GetQuery("id")
		if !b {
			context.JSON(403, Mess{
				Code: 1403,
				Msg:  "权限不足",
				Err:  "",
				Data: nil,
			})
			return
		}
		min, err := model.Find("id=" + id)
		if err != nil {
			context.JSON(403, Mess{
				Code: 1403,
				Msg:  "权限不足",
				Err:  err.Error(),
				Data: nil,
			})
			return
		}
		link, err := lib.LoginWeiXin(min.Cookies(), min.GetBase())
		if err != nil {
			context.JSON(403, Mess{
				Code: 1403,
				Msg:  "权限不足",
				Err:  err.Error(),
				Data: nil,
			})
			return
		}

		context.JSON(200, Mess{
			Code: 1200,
			Msg:  "操作成功",
			Err:  "",
			Data: gin.H{"link": link, "url": "https://open.weixin.qq.com" + link},
		})
	}
}

// GetInfo
// @Description: 获取用户信息
// @return gin.HandlerFunc
// example
// @Summary 获取用户信息
// @Tags min
// @Accept json
// @Produce json
// @Param id query int true "用户ID"
// @Success 200 {object}  lib.Info
// @Failure 403 {object}  Mess
// @Router /min/info [POST]
func GetInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, b := context.GetQuery("id")
		if !b {
			context.JSON(403, Mess{
				Code: 1403,
				Msg:  "权限不足",
				Err:  "",
				Data: nil,
			})
			return
		}
		min, err := model.Find("id=" + id)
		if err != nil {
			context.JSON(403, Mess{
				Code: 1403,
				Msg:  "权限不足",
				Err:  err.Error(),
				Data: nil,
			})
			return
		}
		context.JSON(200, lib.GetLoginInfo(min.Cookies(), min.GetBase()))
	}
}

// CheckQrCode
// @Description: 检查二维码扫码状态
// @return gin.HandlerFunc
// example
// @Summary 检查二维码扫码状态
// @Tags min
// @Accept json
// @Produce json
// @Param qr body config.CheckQrCode true "账号密码"
// @Success 200 {object} Mess
// @Failure 403 {object} Mess
// @Router /min/check_qr_code [POST]
func CheckQrCode() gin.HandlerFunc {
	return func(context *gin.Context) {
		var qr config.CheckQrCode
		err := context.BindJSON(&qr)
		if err != nil {
			context.JSON(403, Mess{
				Code: 1403,
				Msg:  "获取数据失败",
				Err:  err.Error(),
				Data: nil,
			})
			return
		}
		min, err := model.Find("id=" + strconv.Itoa(qr.ID))
		if err != nil {
			context.JSON(403, Mess{
				Code: 1403,
				Msg:  "权限不足",
				Err:  err.Error(),
				Data: nil,
			})
			return
		}
		code := lib.CheckQrCode(qr.Link, min.Cookies(), min.GetBase())
		if code {
			min.Id = qr.ID
			min.State = 1
			_, err := model.Update(*min)
			if err != nil {
				context.JSON(403, Mess{
					Code: 1403,
					Msg:  "权限不足",
					Err:  err.Error(),
					Data: nil,
				})
				return
			}
			context.JSON(200, Mess{
				Code: 1200,
				Msg:  "登录成功",
				Err:  "",
				Data: nil,
			})
			return
		} else {
			context.JSON(200, Mess{
				Code: 1401,
				Msg:  "正在等待扫码",
				Err:  "",
				Data: nil,
			})
			return
		}
	}
}

// GetCourses
// @Description: 获取课程列表
// @return gin.HandlerFunc
// example
// @Summary 获取课程列表
// @Tags min
// @Accept json
// @Produce json
// @Param id query int true "ID"
// @Success 200 {object} Mess{data=[]lib.Course}
// @Failure 403 {object} Mess
// @Router /min/get_courses [POST]
func GetCourses() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, b := context.GetQuery("id")
		if !b {
			context.JSON(403, Mess{
				Code: 1403,
				Msg:  "权限不足",
				Err:  "",
				Data: nil,
			})
			return
		}
		min, err := model.Find("id=" + id)
		if err != nil {
			context.JSON(403, Mess{
				Code: 1403,
				Msg:  "",
				Err:  err.Error(),
				Data: nil,
			})
		}
		courses, err := lib.GetCourse(min.Cookies(), min.GetBase())
		if err != nil {
			context.JSON(403, Mess{
				Code: 1403,
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
			Data: courses,
		})
	}
}

// GetCourseList
// @Description: 获取课程信息
// @return gin.HandlerFunc
// example
// @Summary 获取课程列表
// @Tags min
// @Accept json
// @Produce json
// @Param get body config.GetCourseList true "ID:用户id\nCourseID:课程id"
// @Success 200 {object} Mess{data=[]lib.Course}
// @Failure 403 {object} Mess
// @Router /min/get_course_list [POST]
func GetCourseList() gin.HandlerFunc {
	return func(context *gin.Context) {
		var getList config.GetCourseList
		err := context.BindJSON(&getList)
		if err != nil {
			context.JSON(403, Mess{
				Code: 1403,
				Msg:  "",
				Err:  err.Error(),
				Data: nil,
			})
			return
		}
		min, err := model.Find(fmt.Sprintf("id=%d", getList.ID))
		if err != nil {
			context.JSON(403, Mess{
				Code: 1403,
				Msg:  "",
				Err:  err.Error(),
				Data: nil,
			})
			return
		}
		list, err := lib.GetChapter(getList.CourseID, min.Cookies(), min.GetBase())
		if err != nil {
			context.JSON(403, Mess{
				Code: 1403,
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
			Data: list,
		})
	}
}
