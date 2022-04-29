package autn

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"strconv"
)

type ByteHelper struct {
}

// Md5Byte 获取字节数组md5值
func (sa *ByteHelper) Md5Byte(s []byte) string {
	h := md5.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1Byte 获取节数组sha1值
func (sa *ByteHelper) Sha1Byte(s []byte) string {
	h := sha1.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

// Sha256Byte 获取节数组sha256值
func (sa *ByteHelper) Sha256Byte(s []byte) string {
	h := sha256.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

// Sha512Byte 获取节数组sha512值
func (sa *ByteHelper) Sha512Byte(s []byte) string {
	h := sha512.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

// bytes to hex string
func (sa *ByteHelper) BytesToHexString(b []byte) string {
	var buf bytes.Buffer
	for _, v := range b {
		t := strconv.FormatInt(int64(v), 16)
		if len(t) > 1 {
			buf.WriteString(t)
		} else {
			buf.WriteString("0" + t)
		}
	}
	return buf.String()
}

// hex string to bytes
func (sa *ByteHelper) HexStringToBytes(s string) []byte {
	bs := make([]byte, 0)
	for i := 0; i < len(s); i = i + 2 {
		b, _ := strconv.ParseInt(s[i:i+2], 16, 16)
		bs = append(bs, byte(b))
	}
	return bs
}
