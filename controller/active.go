package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"min/config"
	"min/lib"
	"min/model"
)

func GetLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// DoActive
// @Description: 执行一个任务
// @return gin.HandlerFunc
// example
// @Summary 执行一个任务
// @Tags active
// @Accept json
// @Produce json
// @Param active body config.DoActive true "用户id以及课程id"
// @Success 200 {object} Mess
// @Failure 403 {object} Mess
// @Router /active/do_active [POST]
func DoActive() gin.HandlerFunc {
	return func(context *gin.Context) {
		var ac config.DoActive
		err := context.BindJSON(&ac)
		if err != nil {
			context.JSON(403, Mess{
				Code: 1403,
				Msg:  "",
				Err:  err.Error(),
				Data: nil,
			})
			return
		}
		min, err := model.Find(fmt.Sprintf("id=%d", ac.UserID))
		if err != nil {
			context.JSON(403, Mess{
				Code: 1403,
				Msg:  "",
				Err:  err.Error(),
				Data: nil,
			})
			return
		}
		list, err := lib.GetChapter(ac.CourseID, min.Cookies())
		if err != nil {
			context.JSON(403, Mess{
				Code: 1403,
				Msg:  "获取课程信息错误",
				Err:  err.Error(),
				Data: nil,
			})
			return
		}

		go func() {
			defer func() {
				err := recover()
				if err != nil {
					log.Errorln("提交学时出现不可预料的错误")
					log.Errorln(err)
				}
			}()
			id := 0
			log.Infoln(ac)

			find, err := model.FindActive(fmt.Sprintf("user_id=%d and course_id=%d", ac.UserID, ac.CourseID))
			if err != nil || find.Id == 0 {
				id, err = model.AddActive(model.Active{
					UserId:   ac.UserID,
					CourseId: ac.CourseID,
					Status:   1,
				})
				if err != nil {
					return
				}
			} else {
				id = find.Id
				find.Status = 1
				_, err := model.UpdateActive(find)
				if err != nil {
					log.Errorln("更新active错误" + err.Error())
					return
				}
			}
			lib.CommitTime(min.Cookies(), list, id)
		}()
		context.JSON(200, Mess{
			Code: 1200,
			Msg:  "操作成功",
			Err:  "",
			Data: nil,
		})
	}
}

// GetActives
// @Description: 获取所有活动任务列表
// @return gin.HandlerFunc
// example
// @Summary 获取所有活动任务列表
// @Tags active
// @Accept json
// @Produce json
// @Success 200 {object} Mess{data=[]model.Active}
// @Failure 403 {object} Mess
// @Router /active/get_actives [POST]
func GetActives() gin.HandlerFunc {
	return func(context *gin.Context) {
		actives, err := model.QueryActive("1=1")
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
			Data: actives,
		})
	}
}

// StopActive
// @Description: 停止一个活动任务
// @return gin.HandlerFunc
// example
// @Summary 停止一个活动任务
// @Tags active
// @Accept json
// @Produce json
// @Param active_id query int true "任务id"
// @Success 200 {object} Mess
// @Failure 403 {object} Mess
// @Router /active/stop_active [POST]
func StopActive() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, b := context.GetQuery("active_id")
		if !b {
			context.JSON(403, Mess{
				Code: 1403,
				Msg:  "参数不足",
				Err:  "",
				Data: nil,
			})
			return
		}
		active, err := model.FindActive(fmt.Sprintf("id=%v", id))
		if err != nil {
			context.JSON(403, Mess{
				Code: 1403,
				Msg:  "权限不足",
				Err:  err.Error(),
				Data: nil,
			})
			return
		}
		active.Status = -1
		_, err = model.UpdateActive(active)
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
			Data: nil,
		})
	}
}

// DeleteActive
// @Description: 删除一个活动任务
// @return gin.HandlerFunc
// example
// @Summary 删除一个活动任务
// @Tags active
// @Accept json
// @Produce json
// @Param active_id query int true "任务id"
// @Success 200 {object} Mess
// @Failure 403 {object} Mess
// @Router /active/delete_active [POST]
func DeleteActive() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, b := context.GetQuery("active_id")
		if !b {
			context.JSON(403, Mess{
				Code: 1403,
				Msg:  "权限不足",
				Err:  "",
				Data: nil,
			})
			return
		}
		id1, err := strconv.Atoi(id)
		if err != nil {
			context.JSON(403, Mess{
				Code: 1403,
				Msg:  "权限不足",
				Err:  err.Error(),
				Data: nil,
			})
			return
		}
		err = model.DeleteActive(id1)
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
			Data: nil,
		})
	}
}
