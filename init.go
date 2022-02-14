package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"min/lib"
	"min/model"
)

// 初始化时检查未运行完的任务
func startActivity() {
	actives, err := model.QueryActive("status=1")
	if err != nil {
		return
	}
	for _, active := range actives {
		min, err := model.Find(fmt.Sprintf("id=%v", active.UserId))
		if err != nil {
			log.Errorln("查找用户出现错误" + err.Error())
			return
		}
		list, err := lib.GetChapter(active.CourseId, min.Cookies())
		if err != nil {
			log.Errorln("查询课程信息错误")
			return
		}
		go func(active2 model.Active) {
			defer func() {
				i := recover()
				if i != nil {
					log.Errorln("学习出现异常")
				}
			}()
			defer func() {
				active2.Status = 0
				_, err := model.UpdateActive(active2)
				if err != nil {
					return
				}
			}()

			lib.CommitTime(min.Cookies(), list, active2.Id)
		}(active)
		log.Infoln("已开始运行任务", active.Id)
	}
}
