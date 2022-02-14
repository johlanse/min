package lib

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

var data = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0, shrink-to-fit=no, viewport-fit=cover">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="format-detection" content="telephone=no">
    <meta charset="UTF-8">
    <title>错误提示</title>
    <link rel="stylesheet" href="/static/common/css/error.css" type="text/css">
</head>
<body>
<div id="error-main" class="errormain">
    <div class="icon"><img src="/static/common/image/sososo.png" alt=""></div>
    <div class="name">课程章节信息不存在!</div>
    <div class="note">页面将在<span id="wait">5</span>秒后自动返回，<a id="back" href="">返回上一页</a></div>
</div>
</body>
</html>
<script type="text/javascript">var data ={"status":false,"msg":"\u8bfe\u7a0b\u7ae0\u8282\u4fe1\u606f\u4e0d\u5b58\u5728!","back":""};</script>
<script type="text/javascript" src="/static/common/js/error.js?v=1.0.3"></script>
`

func TestName(t *testing.T) {
	response, b := getResponse(data)
	if b {
		log.Println(response)
	} else {
		print(false)
	}
}
