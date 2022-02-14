package model

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"xorm.io/xorm"

	_ "github.com/mattn/go-sqlite3"
)

type Min struct {
	Id       int    `xorm:"INTEGER autoincr pk"`
	Account  string `xorm:"TEXT"`
	Password string `xorm:"TEXT"`
	Token    string `xorm:"TEXT"`
	Tgw      string `xorm:"TEXT"`
	State    int    `xorm:"INTEGER"`
}

func init() {
	err := connect()
	if err != nil {
		log.Errorln("数据库链接失败")
		log.Errorln(err.Error())
		return
	}
	err = engine.CreateTables(&Course{})
	if err != nil {
		return
	}
}

var engine *xorm.Engine

func connect() error {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./data/data.db")
	if err != nil {
		return err
	}
	return err
}

// Add
/**
 * @Description:
 * @param min
 * @return error
 * example
 */
func Add(min Min) (int, error) {
	err := engine.Ping()
	if err != nil {
		err1 := connect()
		if err1 != nil {
			return -1, err1
		}
	}
	one, err := engine.InsertOne(&min)
	if err != nil {
		return -1, err
	}
	log.Infoln(min.Id)
	if one < 1 {
		return -1, errors.New("影响行数为0")
	}
	return min.Id, err
}

// Find
/**
 * @Description:
 * @param condition
 * @return Min
 * @return error
 * example
 */
func Find(condition string) (Min, error) {
	min := Min{}
	err := engine.Ping()
	if err != nil {
		err1 := connect()
		if err1 != nil {
			return min, err1
		}
	}
	get, err := engine.Where(condition).Get(&min)
	if err != nil {
		return Min{}, err
	}
	if !get {
		return min, errors.New("don't have data")
	}
	return min, err
}

// Query
/**
 * @Description:
 * @param condition
 * @return []Min
 * @return error
 * example
 */
func Query(condition string) ([]Min, error) {
	var mins = make([]Min, 0)
	err := engine.Where(condition).Find(&mins)
	if err != nil {
		return nil, err
	}
	return mins, err
}

func Update(min Min) (int, error) {
	update, err := engine.ID(min.Id).Update(&min)
	if err != nil {
		return -1, err
	}
	return int(update), nil
}
