package HTTP

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/errors_"
)

type HttpInfo struct {
	HttpClient *http.Client
	Url        string
	Para       []string
	TimeOut    int
}

func (this *HttpInfo) createHttpConns(sLogin *config.HttpData) {
	if 0 == sLogin.Time_out {
		errors_.CheckFatalErr(errors.New("can't read http post timeout"))
	}
	this.HttpClient = &http.Client{Timeout: time.Duration(sLogin.Time_out) * time.Millisecond}
	this.Url = sLogin.Url
	for _, param := range sLogin.Para {
		this.Para = append(this.Para, param)
	}
	this.TimeOut = sLogin.Time_out
}

func (this *HttpInfo) GetHttpConnFromConf(c *config.ConfigEngine, name string) {
	this.createHttpConns(c.GetHttpFromConf(name))
}

// 线上设置url参数
func (this *HttpInfo) SetUrlPara(values ...interface{}) string {
	var url_tmp string = this.Url
	u, err := url.Parse(url_tmp)
	errors_.CheckCommonErr(err)
	for i, val := range values {
		sVal, err := val.(string)
		if false == err {
			sVal = strconv.Itoa(val.(int))
		}
		q := u.Query()
		if len(this.Para) <= i {
			errors_.CheckCommonErr(errors.New("Set Url Para error"))
		}
		q.Set(this.Para[i], sVal)
		u.RawQuery = q.Encode()
	}
	return u.String()
}

//发送GET请求
//url:请求地址
//response:请求返回的内容
func (this *HttpInfo) HttpGet(url string) (response string, ok bool) {
	resp, err := this.HttpClient.Get(url)
	if err != nil {
		errors_.CheckCommonErr(err)
		return "", false
	}
	defer resp.Body.Close()
	if 200 == resp.StatusCode {
		body, err := ioutil.ReadAll(resp.Body)
		errors_.CheckCommonErr(err)
		return string(body), true
	} else {
		return "", false
	}
}

//发送POST请求
//url:请求地址，data:POST请求提交的数据,contentType:请求体格式，如：application/json_struct
//content:请求放回的内容
func (this *HttpInfo) HttpPost(url string, data interface{}, contentType string) (content string) {
	jsonStr, err := json.Marshal(data)
	errors_.CheckCommonErr(err)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	errors_.CheckCommonErr(err)
	req.Header.Set("Content-Type", contentType)
	defer req.Body.Close()
	resp, err := this.HttpClient.Do(req)
	errors_.CheckCommonErr(err)
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	errors_.CheckCommonErr(err)
	return string(result)
}
