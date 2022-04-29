package autn

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"regexp"
)

// 封装一些字符串相关的常用方法
type StringHelper struct {
}

// 字符串Base64编码
func (tt *StringHelper) Base64Encode(data string) string {
	return base64.URLEncoding.EncodeToString([]byte(data))
}

// 字符串Base64解码
func (tt *StringHelper) Base64Decode(data string) string {
	b, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

// 字符串反转
func (tt *StringHelper) Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//展开并格式化对象以json方式
func (tt *StringHelper) Export(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, b, "", "\t")
	if err != nil {
		return ""
	}
	return buf.String()
}

// json化
func (tt *StringHelper) Json(v interface{}) string {
	return tt.Export(v)
}

//json转对象
func (tt *StringHelper) Json2Object(str string, v interface{}) error {
	return json.Unmarshal([]byte(str), &v)
}

// 检测字符串是否为空
// check string is empty
func (tt *StringHelper) IsEmpty(s string) bool {
	return len(s) == 0
}

// 判断字符串是否不为空
// check string is not empty
func (tt *StringHelper) IsNotEmpty(s string) bool {
	return !tt.IsEmpty(s)
}

// 判断字符串是否是空白字符串
// check string is whitespace,empty
func (tt *StringHelper) IsBlank(s string) bool {
	if len(s) == 0 {
		return true
	}
	reg := regexp.MustCompile(`^\s+$`)
	return reg.MatchString(s)
}

// 判断字符串是否为不为空白字符串
func (tt *StringHelper) IsNotBlank(s string) bool {
	return !tt.IsBlank(s)
}
