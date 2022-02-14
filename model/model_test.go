package model

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestAdd(t *testing.T) {
	err := Add(Min{
		Id:       0,
		Account:  "19104978",
		Password: "1743224847gou",
	})
	if err != nil {
		log.Errorln(err.Error())
		return
	}
}

func TestFind(t *testing.T) {
	min, err := Find("id=0")
	if err != nil {
		return
	}
	log.Infoln(min)
}

func TestQuery(t *testing.T) {
	mins, err := Query("id=1")
	if err != nil {
		return
	}
	log.Infoln(mins)
}
