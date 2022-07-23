package autn

// @Title  web相关的常用方法封装
// @Description  封装一些http相关的一些常用操作
// @Author  0opslab

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type HttpHelper struct {
}

// JsQueryEscape escapes the string in javascript standard so it can be safely placed
// inside a URL query.
func (ss *HttpHelper) JsQueryEscape(s string) string {
	return strings.Replace(url.QueryEscape(s), "+", "%20", -1)
}

// JsQueryUnescape does the inverse transformation of JsQueryEscape, converting
// %AB into the byte 0xAB and '+' into ' ' (space). It returns an error if
// any % is not followed by two hexadecimal digits.
func (ss *HttpHelper) JsQueryUnescape(s string) (string, error) {
	return url.QueryUnescape(strings.Replace(s, "%20", "+", -1))
}

//Get HTTP RESPONSE BODY DATA
func (ss *HttpHelper) HttpGet(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(data[:])
}

//DOWNLOAD FILE FORM URL
func (ss *HttpHelper) HttpDownload(url string, file string) bool {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存
	out, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
	return true
}

// @title    getCurrentIP
// @description   获取http请求端ip地址
func (ss *HttpHelper) GetRemoteIP(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-Ip"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}

//统一响应
func (ss *HttpHelper) HttpResponse(w http.ResponseWriter, response Response) {
	w.Header().Add("Content-Type", "application/json;charset:utf-8;")
	json.NewEncoder(w).Encode(response)
}

func (ss *HttpHelper) HttpResponseCode(w http.ResponseWriter, code int, message string) {
	response := Response{
		Code: code,
		Msg:  message,
		Data: nil,
	}
	w.Header().Add("Content-Type", "application/json;charset:utf-8;")
	json.NewEncoder(w).Encode(response)
}

func (ss *HttpHelper) HttpResponseCodeData(w http.ResponseWriter, code int, message string, data interface{}) {
	response := Response{
		Code: code,
		Msg:  message,
		Data: data,
	}
	w.Header().Add("Content-Type", "application/json;charset:utf-8;")
	json.NewEncoder(w).Encode(response)
}

//判断某个ip是否属于某个网段
func (ss *HttpHelper) IpBeing(ip, cidr string) bool {
	ipAddr := strings.Split(ip, `.`)
	if len(ipAddr) < 4 {
		return false
	}
	cidrArr := strings.Split(cidr, `/`)
	if len(cidrArr) < 2 {
		return false
	}
	var tmp = make([]string, 0)
	for key, value := range strings.Split(`255.255.255.0`, `.`) {
		iint, _ := strconv.Atoi(value)

		iint2, _ := strconv.Atoi(ipAddr[key])

		tmp = append(tmp, strconv.Itoa(iint&iint2))
	}
	return strings.Join(tmp, `.`) == cidrArr[0]
}
