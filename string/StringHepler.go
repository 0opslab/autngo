package string

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"regexp"
)

// 封装一些字符串相关的常用方法
type StringHelper struct{

}
// 字符串Base64编码
func (this *StringHelper) Base64Encode(data string) string {
	return base64.URLEncoding.EncodeToString([]byte(data))
}

// 字符串Base64解码
func (this *StringHelper) Base64Decode(data string) string {
	b, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

// // Md5String 获取字符串md5值
// func (this *StringHelper) Md5String(s string) string {
// 	return Md5Byte([]byte(s))
// }

// // Sha1String 获取字符串sha1值
// func (this *StringHelper) Sha1String(s string) string {
// 	return Sha1Byte([]byte(s))
// }

// // Sha256String 获取字符串sha256值
// func (this *StringHelper) Sha256String(s string) string {
// 	return Sha256Byte([]byte(s))
// }

// // Sha512String 获取字符串sha512值
// func (this *StringHelper) Sha512String(s string) string {
// 	return Sha512Byte([]byte(s))
// }

// 字符串反转
func (this *StringHelper) Reverse(s string) string{
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
//展开并格式化对象以json方式
func (this *StringHelper) Export(v interface{}) string {
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
func (this *StringHelper) Json(v interface{}) string {
	return this.Export(v)
}

// 检测字符串是否为空
// check string is empty
func (this *StringHelper) IsEmpty(s string) bool {
	if len(s) == 0 {
		return true
	}
	return false
}

// 判断字符串是否不为空
// check string is not empty
func (this *StringHelper) IsNotEmpty(s string) bool{
	return !this.IsEmpty(s)
}


// 判断字符串是否是空白字符串
// check string is whitespace,empty
func (this *StringHelper) IsBlank(s string) bool{
	if len(s) == 0{
		return true
	}
	reg := regexp.MustCompile(`^\s+$`)
	actual := reg.MatchString(s)
	if actual {
		return true
	}
	return false
}

// 判断字符串是否为不为空白字符串
func (this *StringHelper) IsNotBlank(s string) bool{
	return !this.IsBlank(s)
}

