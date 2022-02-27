package lib

type School interface {
	// GetCodeURL 返回验证码url
	GetCodeURL() string
	GETLoginURL() string
}

type ShiXun struct {
}

func (s *ShiXun) GETLoginURL() string {
	return "https://shixun.cdcas.com/user/login"
}

func (s *ShiXun) GetCodeURL() string {
	return "https://shixun.cdcas.com/service/code?r={time()}"
}
