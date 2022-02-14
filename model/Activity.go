package model

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

type Active struct {
	Id       int `xorm:"INTEGER autoincr pk" json:"id,omitempty"`
	UserId   int `xorm:"int" json:"user_id,omitempty"`
	CourseId int `xorm:"int" json:"course_id,omitempty"`
	Status   int `xorm:"int" json:"status,omitempty"`
}

func AddActive(actives Active) (int, error) {
	err := engine.Ping()
	if err != nil {
		err1 := connect()
		if err1 != nil {
			return -1, err1
		}
	}
	one, err := engine.InsertOne(&actives)
	if err != nil {
		return -1, err
	}
	log.Infoln(actives.Id)
	if one < 1 {
		return -1, errors.New("影响行数为0")
	}
	return actives.Id, err
}

// FindActive
/**
 * @Description:
 * @param condition
 * @return Active
 * @return error
 * example
 */
func FindActive(condition string) (Active, error) {
	actives := Active{}
	err := engine.Ping()
	if err != nil {
		err1 := connect()
		if err1 != nil {
			return actives, err1
		}
	}
	get, err := engine.Where(condition).Get(&actives)
	if err != nil || !get {
		return Active{}, err
	}
	return actives, err
}

// QueryActive
/**
 * @Description:
 * @param condition
 * @return []Active
 * @return error
 * example
 */
func QueryActive(condition string) ([]Active, error) {
	var activess = make([]Active, 0)
	err := engine.Where(condition).Find(&activess)
	if err != nil {
		return nil, err
	}
	return activess, err
}

func UpdateActive(actives Active) (int, error) {
	update, err := engine.ID(actives.Id).Update(&actives)
	if err != nil {
		return -1, err
	}
	return int(update), nil
}

func DeleteActive(id int) error {
	_, err := engine.ID(id).Delete()
	if err != nil {
		return err
	}
	return nil
}
