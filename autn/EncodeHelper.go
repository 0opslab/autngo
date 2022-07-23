package autn

import (
	"bytes"
	"compress/flate"
	"encoding/hex"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type EncodeHelper struct {
}

// transform GBK bytes to UTF-8 bytes
func (gt *EncodeHelper) GbkToUtf8(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}

// transform UTF-8 bytes to GBK bytes
func (gt *EncodeHelper) Utf8ToGbk(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewEncoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}

// transform GBK string to UTF-8 string and replace it, if transformed success
// returned nil error, or died by error message
func (gt *EncodeHelper) StrToUtf8(str string) (string, error) {
	b, err := gt.GbkToUtf8([]byte(str))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// transform UTF-8 string to GBK string and replace it, if transformed success
// returned nil error, or died by error message
func (gt *EncodeHelper) StrToGBK(str string) (string, error) {
	b, err := gt.Utf8ToGbk([]byte(str))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// 字符串转16进制
func (gt *EncodeHelper) StrHex(str string) string {
	return hex.EncodeToString([]byte(str))
}

// 十六进制转字符串
func (gt *EncodeHelper) HexStr(hexStr string) string {
	b, err := hex.DecodeString(hexStr)
	if err != nil {
		return ""
	}
	return string(b)
}

//@func 字符串zip编码压缩
func (gt *EncodeHelper) ZipEncode(input string) (result []byte, err error) {
	var buf bytes.Buffer
	w, err := flate.NewWriter(&buf, -1)
	w.Write([]byte(input))
	w.Close()
	result = buf.Bytes()
	return
}

//@func 字符串zip编码解压缩还原
func (gt *EncodeHelper) ZipDecode(input []byte) (result string, err error) {
	bytess, err := ioutil.ReadAll(flate.NewReader(bytes.NewReader(input)))
	return string(bytess), err
}
