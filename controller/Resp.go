package controller

type Mess struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Err  string      `json:"err"`
	Data interface{} `json:"data"`
}
