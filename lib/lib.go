package lib

import (
	"compress/gzip"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/htmlquery"
	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/dataflow"
	"github.com/imroc/req/v3"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"github.com/tidwall/gjson"

	"min/model"
	"min/utill"
)

var (
	headers = map[string]string{
		"Accept": "application/json,text/javascript,*/*;q=0.01",

		"X-Requested-With": "XMLHttpRequest",
		"Host":             "shixun.cdcas.com",
		"Origin":           "https://shixun.cdcas.com",
		"Content-Type":     "application/x-www-form-urlencoded;charset=UTF-8",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36 Edg/91.0.864.48",
	}
	codeURL = "http://139.155.247.64:7000/data"
	client  *dataflow.Gout
	C       *req.Client
)

func init() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%time%] [%lvl%]: %msg%\n",
	})
	log.SetLevel(log.DebugLevel)

	client = gout.New()
	//client.SetProxy("http://127.0.0.1:8898")
	C = req.C().DevMode()
	C.SetCommonHeaders(headers)
	//C.SetProxyURL("http://127.0.0.1:8898")
	C.OnAfterResponse(func(c *req.Client, response *req.Response) error {
		log.Debugln("开始执行中间件", response.Header.Get("Content-Encoding"))
		log.Debugln("已处理gzip结果")
		response.Body, _ = gzip.NewReader(response.Body)
		err := response.UnmarshalJson(response.Result())
		if err != nil {
			return err
		}
		return nil
	})

}

type Course struct {
	Name     string `json:"name"`
	Id       int    `json:"id"`
	Link     string `json:"link"`
	Progress string `json:"progress"`
}

type Response struct {
	Status  bool           `json:"status"`
	Data    string         `json:"data"`
	Cookies []*http.Cookie `json:"cookies"`
}

type Info struct {
	Name      string `json:"name"`
	StudentID string `json:"student_id"`
	Grade     string `json:"grade"`
}

// Identify
/**
 * @Description: 二维码识别
 * @param code
 * @return string
 * example
 */
func Identify(code []byte) string {
	code1 := base64.StdEncoding.EncodeToString(code)
	var result string
	err := gout.POST(codeURL).SetJSON(gout.H{"img": code1}).BindBody(&result).Do()
	if err != nil {
		return ""
	}
	log.Debugln(gjson.Get(result, "@this|@pretty"))
	return gjson.Get(result, "value").String()
}

func onLine(cookies []*http.Cookie, base string) {
	defer func() {
		i := recover()
		if i != nil {
			log.Errorln("服务器保活出现异常")
		}
	}()
	err := client.POST(base + "/user/online").SetCookies(cookies...).Do()
	if err != nil {
		return
	}
	log.Infoln("向服务器保活成功")
}

// GetProgress
// @Description:
// @param courseID
// @param cookies
// @return string
// example
// @Summary
// @Tags
// @Accept
// @Produce
// @Param
// @Success 200 {object}
// @Failure 403 {object}
// @Router / [POST]
func GetProgress(courseID int, cookies []*http.Cookie, base string) string {

	list, err := GetChapter(courseID, cookies, base)
	if err != nil {
		return "0%"
	}
	n := 0.00
	for _, course := range list.List {
		videoLen := stringTimeToLong(course.VideoDuration)
		if stringTimeToLong(course.ViewedDuration) >= videoLen {
			n++
		}
	}
	progress := n / float64(len(list.List))
	log.Debugln(progress)
	return fmt.Sprintf("%v%%", int(progress*100))
}

// GetLoginInfo
/**
 * @Description: 获取用户信息
 * @param cookies
 * @return Info
 * example
 */
func GetLoginInfo(cookies []*http.Cookie, base string) Info {
	if !CheckLogin(cookies, base) {
		return Info{}
	}
	var info Info
	response, err := client.GET(base + `/user/member`).SetHeader(gout.H{
		"Host":       "shixun.cdcas.com",
		"Origin":     base + "",
		"Referer":    base + "/user/",
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36 Edg/91.0.864.48",
		"Accept":     "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	}).SetCookies(cookies...).Response()
	if err != nil {
		return Info{}
	}
	log.Infoln(response.Status)
	doc, err := htmlquery.Parse(response.Body)
	if err != nil {
		return Info{}
	}
	one := htmlquery.FindOne(doc, `//*[@id="userForm"]/div[2]/div[1]/div[2]/text()`)
	if err != nil {
		log.Errorln(err.Error())
		return Info{}
	}
	log.Infoln(one)
	info.Name = one.Data
	node1 := htmlquery.FindOne(doc, "//*[@id=\"userForm\"]/div[2]/div[2]/div[2]/text()")
	if err != nil {
		log.Errorln(err.Error())
		return Info{}
	}
	info.StudentID = node1.Data
	node2 := htmlquery.FindOne(doc, "//*[@id=\"userForm\"]/div[2]/div[4]/div[2]/text()")
	if err != nil {
		return Info{}
	}
	info.Grade = node2.Data
	return info
}

func CommitTime(cookies []*http.Cookie, list CourseList, ActiveID int, base string) {
	if !CheckLogin(cookies, base) {
		log.Errorln("cookie已过期")
		return
	}
	active, err := model.FindActive(fmt.Sprintf("id=%d", ActiveID))
	if err != nil {
		log.Errorln("未能查找到active")
	}
	if active.Status == -1 {
		log.Errorln("检测到已手动终止任务 :", active)
	}

	logger := log.New()
	logger.SetLevel(log.DebugLevel)
	logger.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       fmt.Sprintf("[%%time%%] [%d] [%%lvl%%]: %%msg%%\n", active.CourseId),
	})
	utill.CheckDir("./logs/")
	file, err := os.OpenFile(fmt.Sprintf("./logs/%v_%v.log", ActiveID, active.CourseId), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		logger.Errorln("打开日志文件失败")
		return
	}
	logger.SetOutput(io.MultiWriter(os.Stdout, file))
	info := GetLoginInfo(cookies, base)
	logger.Infoln("开始学习任务，姓名==》", info.Name, "课程号==》", active.CourseId)
	logger.Infoln("共获取到", len(list.List), "个视频")
	defer logger.Infoln("学习任务已退出，姓名==》", info.Name, "课程号==》", active.CourseId)
	// 退出时将任务state设置为-1
	defer func() {
		active.Status = -1
		_, err := model.UpdateActive(active)
		if err != nil {
			logger.Errorln("更新active的state出现错误")
			return
		}
	}()

	for _, course := range list.List {
		studyTime := 1
		studyID := 0
		videoLen := stringTimeToLong(course.VideoDuration)
		status := false
		if stringTimeToLong(course.ViewedDuration) >= videoLen {
			logger.Infoln("已跳过···" + course.Name)
			continue
		}
		logger.Infoln("正在观看："+course.Name+"，视频时长:", videoLen)

		if videoLen >= 600 {
			go onLine(cookies, base)
			for !status {
				ac, err := model.FindActive(fmt.Sprintf("id=%d", ActiveID))
				if err != nil {
					log.Errorln("未能查找到active")
					return
				}
				if ac.Status == -1 {
					logger.Infoln("已手动停止该任务")
				}
				var code []byte
				err = client.GET(base + "/service/code").SetHeader(headers).SetCookies(cookies...).BindBody(&code).Do()
				if err != nil {
					logger.Errorln("获取验证码出现错误")
					logger.Errorln(err.Error())
					logger.Warnln("正在尝试重新获取验证码")
					err = client.GET(base + "/service/code").SetHeader(headers).SetCookies(cookies...).BindBody(&code).Do()
					if err != nil {
						logger.Errorln("第二次获取验证码失败")
						break
					}
				}
				var res []byte

				values := url.Values{}
				values.Add("nodeId", course.Id)
				values.Add("studyId", strconv.Itoa(studyID))
				values.Add("studyTime", strconv.Itoa(studyTime))
				values.Add("code", Identify(code))
				//logger.Infoln("已提交请求",values.Encode())
				err = client.POST(base + "/user/node/study").SetHeader(gout.H{
					"Accept":           "application/json,text/javascript,*/*;q=0.01",
					"X-Requested-With": "XMLHttpRequest",
					"Host":             "shixun.cdcas.com",
					"Origin":           base + "",
					"Content-Type":     "application/x-www-form-urlencoded;charset=UTF-8",
					"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36 Edg/91.0.864.48",
				}).SetCookies(cookies...).SetBody(values.Encode()).BindBody(&res).Do()
				if err != nil {
					logger.Errorln("请求提交验证码学时出现错误" + err.Error())
					continue
				}
				//res, _ = getResponse(res)
				//logger.Debugln("已获取服务端响应", " ", res)
				//status = getStatus(res)
				//data, err := utill.GbkToUtf8(res)
				//if err != nil {
				//	return
				//}

				status = gjson.GetBytes(res, "status").Bool()
				logger.Infoln(gjson.GetBytes(res, "msg"))
				if status {
					studyID = int(gjson.GetBytes(res, "studyId").Int())
					if studyID == 0 {
						status = false
						continue
					}
					logger.Infoln("通过验证码识别")
				}
				time.Sleep(3 * time.Second)
			}
			studyTime += 30
			time.Sleep(time.Second * 30)
		}
		for studyTime-30 < videoLen {
			ac, err := model.FindActive(fmt.Sprintf("id=%d", ActiveID))
			if err != nil {
				log.Errorln("未能查找到active")
				return
			}
			if ac.Status == -1 {
				logger.Infoln("已手动停止该任务")
			}
			go onLine(cookies, base)
			var res []byte
			values := url.Values{}
			values.Add("nodeId", course.Id)
			values.Add("studyId", strconv.Itoa(studyID))
			values.Add("studyTime", strconv.Itoa(studyTime))
			logger.Infoln("已提交请求", values.Encode())
			response, err := client.POST(base + "/user/node/study").SetHeader(gout.H{
				"Accept":           "application/json,text/javascript,*/*;q=0.01",
				"X-Requested-With": "XMLHttpRequest",
				"Host":             "shixun.cdcas.com",
				"Origin":           base + "",
				"Content-Type":     "application/x-www-form-urlencoded;charset=UTF-8",
				"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36 Edg/91.0.864.48",
			}).SetCookies(cookies...).BindBody(&res).SetBody(values.Encode()).Response()
			if err != nil {
				logger.Errorln("请求提交学时出现错误")
				logger.Errorln(err.Error())
				continue
			}
			res, err = io.ReadAll(response.Body)
			if err != nil {
				return
			}
			err = response.Body.Close()
			if err != nil {
				return
			}
			// res, _ = getResponse(res)
			// logger.Infoln("获取服务器请求"+res)
			logger.Debugln(gjson.GetBytes(res, "@this").String())
			// if !gjson.GetBytes(res, "status").Bool() {
			//	logger.Errorln("提交学时出现了异常情况")
			//	logger.Errorln(gjson.GetBytes(res, "msg"))
			//	break
			//}

			if studyID == 0 {
				studyID = int(gjson.GetBytes(res, "studyId").Int())
				// 检查是否获取到studyID，未获取到提交时长会被记录非法时长
				if studyID != 0 {
					logger.Infoln("已获取study_id", studyID)
				} else {
					continue
				}
			}
			logger.Infoln("已提交学时"+"  ", studyTime)
			time.Sleep(time.Second * 30)
			studyTime += 30
		}
	}
}

func GetChapter(courseId int, cookies []*http.Cookie, base string) (CourseList, error) {
	if !CheckLogin(cookies, base) {
		return CourseList{}, errors.New("cookie已过期")
	}
	var chapter CourseList
	//_, err := C.R().SetCookies(cookies...).SetHeaders(map[string]string{
	//	"Accept":           "application/json,text/javascript,*/*;q=0.01",
	//	"X-Requested-With": "XMLHttpRequest",
	//	"Origin":           base + "",
	//	"Content-Type":     "application/x-www-form-urlencoded;charset=UTF-8",
	//	"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36 Edg/91.0.864.48",
	//}).SetResult(&chapter).Get("https://mooc.cdcas.com" + "/user/study_record.json?courseId=" + strconv.Itoa(courseId))
	//if err != nil {
	//	log.Errorln(err.Error())
	//	return CourseList{}, err
	//}

	err := client.GET(base + "/user/study_record.json?courseId=" + strconv.Itoa(courseId)).SetCookies(cookies...).SetHeader(gout.H{}).BindJSON(&chapter).Do()
	if err != nil {
		return chapter, err
	}
	for i := 2; i <= chapter.PageInfo.PageCount; i++ {
		var chapter1 CourseList
		err = client.GET(base + "/user/study_record.json?courseId=" + strconv.Itoa(courseId) + "&page=" + strconv.Itoa(i)).SetCookies(cookies...).SetHeader(gout.H{}).BindJSON(&chapter1).Do()
		if err != nil {
			return chapter, err
		}
		//_, err := C.R().SetCookies(cookies...).SetHeaders(map[string]string{
		//	"Accept":           "application/json,text/javascript,*/*;q=0.01",
		//	"X-Requested-With": "XMLHttpRequest",
		//	"Origin":           base + "",
		//	"Content-Type":     "application/x-www-form-urlencoded;charset=UTF-8",
		//	"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36 Edg/91.0.864.48",
		//}).SetResult(&chapter1).Get("https://mooc.cdcas.com" + "/user/study_record.json?courseId=" + strconv.Itoa(courseId) + "&page=" + strconv.Itoa(i))
		//if err != nil {
		//	log.Errorln(err.Error())
		//	return CourseList{}, err
		//}
		chapter.List = append(chapter.List, chapter1.List...)
	}
	return chapter, err
}

func GetCourse(cookies []*http.Cookie, base string) ([]Course, error) {
	var courses []Course
	if !CheckLogin(cookies, base) {
		return courses, errors.New("cookie已过期")
	}

	response, err := client.GET(base + "/user").SetCookies(cookies...).SetHeader(gout.H{
		"Host":       "shixun.cdcas.com",
		"Origin":     base + "",
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36 Edg/91.0.864.48",
		"Accept":     "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	}).Response()
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}
	doc.Find("div.item div.con").Each(func(i int, selection *goquery.Selection) {
		link, ok := selection.Find("div.name a").Attr("href")
		if ok {
			var course Course
			course.Name = selection.Find("div.name a").Text()
			course.Link = link
			course.Id, _ = strconv.Atoi(strings.ReplaceAll(link, "/user/course?courseId=", ""))
			course.Progress = selection.Find("div.progress div.txt").Text()
			log.Debugln("进度==》" + selection.Find("div.progress div.txt").Text())
			courses = append(courses, course)
		} else {
			log.Errorln("未找到课程")
		}
	})

	text := doc.Find("div.total").Text()
	s2 := regexp.MustCompile(`共(.*?)条`).FindString(text)
	s2 = strings.ReplaceAll(s2, "共", "")
	s2 = strings.ReplaceAll(s2, "条", "")
	response.Body.Close()
	log.Debugln(s2)
	count, err := strconv.Atoi(s2)
	if err != nil {
		log.Errorln("正则提取条数错误" + err.Error())
		return courses, err
	}
	log.Debugln(count)
	if count <= 5 {
		log.Debugln("条数小于5")
		return courses, nil
	}
	var i int
	if count%5 == 0 {
		i = count / 5
	} else {
		i = (count / 5) + 1
	}
	for j := 2; j <= i; j++ {
		response, err := client.GET(base + "/user&page=" + strconv.Itoa(j)).SetCookies(cookies...).SetHeader(gout.H{
			"Host":       "shixun.cdcas.com",
			"Origin":     base + "",
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36 Edg/91.0.864.48",
			"Accept":     "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		}).Response()
		if err != nil {
			return nil, err
		}

		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			return nil, err
		}

		doc.Find("div.item div.con").Each(func(i int, selection *goquery.Selection) {
			link, ok := selection.Find("div.name a").Attr("href")
			if ok {
				var course Course
				course.Name = selection.Find("div.name a").Text()
				course.Link = link
				course.Id, _ = strconv.Atoi(strings.ReplaceAll(link, "/user/course?courseId=", ""))
				course.Progress = selection.Find("div.progress dic.txt").Text()
				courses = append(courses, course)
			} else {
				log.Errorln("未找到课程")
			}
		})
		response.Body.Close()
	}
	//for _, course := range courses {
	//	course.Progress = GetProgress(course.Id, cookies)
	//}
	return courses, err
}

// CheckLogin
/**
 * @Description: 检查登录状态
 * @param cookies
 * @return bool
 * example
 */
func CheckLogin(cookies []*http.Cookie, base string) bool {
	response, err := client.GET(base + "/user").SetCookies(cookies...).Response() //nolint:bodyclose
	if err != nil {
		return false
	}
	if response.Request.URL.String() == base+"/user/login" {
		return false
	}

	return true
}

// CheckQrCode
/**
 * @Description: 检查验证码状态
 * @param src
 * @param cookies
 * @return bool
 * example
 */
func CheckQrCode(src string, cookies []*http.Cookie, base string) bool {
	tempURL := "https://lp.open.weixin.qq.com/connect/l/qrconnect?uuid=" + src[16:]
	hea := map[string]string{

		"Accept-Encoding":  "gzip,deflate,br",
		"Accept-Language":  "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"Connection":       "keep-alive",
		"Host":             "lp.open.weixin.qq.com",
		"Referer":          "https: //open.weixin.qq.com/",
		"sec-ch-ua":        `" Not;A Brand";v="99","Microsoft Edge";v="91", "Chromium"; v="91"`,
		"sec-ch-ua-mobile": "?0",
		"Sec-Fetch-Dest":   "script",
		"Sec-Fetch-Mode":   "no-cors",
		"Sec-Fetch-Site":   "same-site",
		"User-Agent":       " Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36 Edg/91.0.864.48",
	}
	var res string
	err := client.GET(tempURL).SetHeader(hea).SetCookies(cookies...).BindBody(&res).Do()
	if err != nil {
		return false
	}
	log.Debugln(res)
	if strings.Contains(res, "window.wx_errcode=405") || strings.Contains(res, "window.wx_errcode=405") {
		compile := regexp.MustCompile(`wx_code='(.*?)'`)
		wxCode := compile.FindStringSubmatch(res)[1]
		err := client.GET(base + "/user").SetCookies(cookies...).SetQuery(gout.H{"code": wxCode, "state": "STATE"}).Do()

		return err == nil
	}
	return false
}

// LoginWeiXin
/**
 * @Description: 登录到微信，获取二维码图片
 * @param cookies
 * @return string 二维码链接
 * @return error
 * example
 */
func LoginWeiXin(cookies []*http.Cookie, base string) (string, error) {
	h := gout.H{
		"authority":       "open.weixin.qq.com",
		"method":          "GET",
		"scheme":          "https",
		"accept":          "text / html, application / xhtml + xml, application / xml;q = 0.9, image / webp, image / apng, * / *;q = 0.8, application / signed - exchange;v = b3;q = 0.9",
		"accept-encoding": "gzip, deflate, br",
		"accept-language": "zh - CN, zh;q = 0.9",
		"cache-control":   "max - age = 0",
	}
	response, err := client.GET(base + "/user").SetHeader(h).SetCookies(cookies...).Response()
	if err != nil {
		log.Errorln("请求/user错误" + err.Error())
		return "", err
	}
	log.Debugln(response.Request.URL.String())
	response, err = client.GET(response.Request.URL.String()).SetCookies(cookies...).Response()
	if err != nil {
		log.Errorln("请求二维码页面错误" + err.Error())
		return "", err
	}
	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Errorln("解析二维码页面错误" + err.Error())
		return "", err
	}
	src, ok := doc.Find("img.qrcode").Attr("src")
	if !ok {
		log.Errorln("查找图片节点错误")
	}
	log.Debugln("图片链接：https://open.weixin.qq.com" + src)
	return src, err
}

// Login
/**
 * @Description: 账号密码登录
 * @param account
 * @param password
 * @return Response
 * example
 */
func Login(account, password string, base string) Response {
	count := 0
	for count < 5 {
		var code []byte
		rsp, err := client.GET(base + "/service/code").SetHeader(headers).Response()
		if err != nil {
			log.Errorln("获取验证码出现错误")
			log.Errorln(err.Error())
			break
		}
		code, _ = io.ReadAll(rsp.Body)
		rsp.Body.Close()

		values := url.Values{}
		//values.Add("schoolId", "47")
		values.Add("username", account)
		values.Add("password", password)
		values.Add("code", Identify(code))
		values.Add("backUrl", "")
		response, err := client.POST(base + "/user/login").SetHeader(headers).SetBody(values.Encode()).
			SetCookies(rsp.Cookies()...).Response()
		if err != nil {
			return Response{}
		}
		body, _ := io.ReadAll(response.Body)
		err = response.Body.Close()
		if err != nil {
			return Response{}
		}

		log.Debugln(gjson.GetBytes(body, "@this|@pretty").String())
		if gjson.GetBytes(body, "status").Bool() {
			return Response{Status: true, Data: gjson.GetBytes(body, "@this|@pretty").String(), Cookies: append(response.Cookies(), rsp.Cookies()...)}
		}
		time.Sleep(3 * time.Second)
		count++
	}

	return Response{
		Status:  false,
		Data:    "多次登录失败",
		Cookies: nil,
	}
}

func checkStr(string2 string) string {
	s := strings.ReplaceAll(string2, " ", "")
	s = strings.ReplaceAll(s, "\n", "")
	return s
}

type CourseList struct {
	List []struct {
		Id             string      `json:"id"`
		Name           string      `json:"name"`
		Type           interface{} `json:"type"`
		ChapterId      string      `json:"chapterId"`
		CourseId       string      `json:"courseId"`
		VideoFile      string      `json:"videoFile"`
		VideoDuration  string      `json:"videoDuration"`
		VotingPath     interface{} `json:"votingPath"`
		TabVideo       string      `json:"tabVideo"`
		TabFile        string      `json:"tabFile"`
		TabVote        string      `json:"tabVote"`
		TabWork        string      `json:"tabWork"`
		TabExam        string      `json:"tabExam"`
		Sort           string      `json:"sort"`
		VideoMode      string      `json:"videoMode"`
		LocalFile      string      `json:"localFile"`
		SchoolId       string      `json:"schoolId"`
		Lock           string      `json:"lock"`
		UnlockTime     string      `json:"unlockTime"`
		Bid            string      `json:"bid"`
		Duration       string      `json:"duration"`
		Progress       string      `json:"progress"`
		State          string      `json:"state"`
		ViewCount      string      `json:"viewCount"`
		FinalTime      *string     `json:"finalTime"`
		Error          int         `json:"error"`
		ErrorMessage   string      `json:"errorMessage"`
		BeginTime      string      `json:"beginTime,omitempty"`
		ViewedDuration string      `json:"viewedDuration"`
		Url            string      `json:"url"`
	} `json:"list"`
	PageInfo struct {
		KeyName      string `json:"keyName"`
		Page         int    `json:"page"`
		PageCount    int    `json:"pageCount"`
		RecordsCount int    `json:"recordsCount"`
		OnlyCount    int    `json:"onlyCount"`
		PageSize     int    `json:"pageSize"`
	} `json:"pageInfo"`
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

func stringTimeToLong(time string) int {
	times := strings.Split(time, ":")
	if len(times) < 2 {
		return 0
	}
	h, err := strconv.Atoi(times[0])
	if err != nil {
		return 0
	}

	m, err := strconv.Atoi(times[1])
	if err != nil {
		return 0
	}
	s, err := strconv.Atoi(times[2])
	if err != nil {
		return 0
	}
	return h*3600 + m*60 + s
}

func getResponse(data string) (string, bool) {
	data = strings.ReplaceAll(data, " ", "")
	data = strings.ReplaceAll(data, "\n", "")
	log.Println(data)
	compile := regexp.MustCompile(`vardata={(.*?)};`)
	match := compile.MatchString(data)
	if match {
		matchs := compile.FindStringSubmatch(data)
		return matchs[1], true
	}
	return "", false
}

func getStudyID(data string) int {
	if data == "" {
		return 0
	}
	datas := strings.Split(data, ",")
	id, err := strconv.Atoi(strings.Split(datas[0], ":")[1])
	if err != nil {
		return 0
	}
	return id
}

func getStatus(data string) bool {
	if data == "" {
		return false
	}
	datas := strings.Split(data, ",")
	parseBool, err := strconv.ParseBool(strings.Split(datas[1], ":")[1])
	if err != nil {
		return false
	}
	return parseBool
}
