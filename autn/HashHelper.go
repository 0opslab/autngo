package autn

// 封装一些字符串相关的常用方法
type HashHelper struct {
	bb ByteHelper
}

// Md5String 获取字符串md5值
func (tt *HashHelper) Md5String(s string) string {
	return tt.bb.Md5Byte([]byte(s))
}

// Sha1String 获取字符串sha1值
func (tt *HashHelper) Sha1String(s string) string {
	return tt.bb.Sha1Byte([]byte(s))
}

// Sha256String 获取字符串sha256值
func (tt *HashHelper) Sha256String(s string) string {
	return tt.bb.Sha256Byte([]byte(s))
}

// Sha512String 获取字符串sha512值
func (tt *HashHelper) Sha512String(s string) string {
	return tt.bb.Sha512Byte([]byte(s))
}
