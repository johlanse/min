package models

type Min struct {
	Id       int    `xorm:"INTEGER"`
	Account  string `xorm:"TEXT"`
	Password string `xorm:"TEXT"`
	Token    string `xorm:"TEXT"`
	Tgw      string `xorm:"TEXT"`
	State    int    `xorm:"integer"`
}

type Active struct {
	Id      int `xorm:"INTEGER"`
	UserId  int `xorm:"int"`
	CurseId int `xorm:"int"`
	Status  int `xorm:"int"`
}
