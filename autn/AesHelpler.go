package autn

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type AesHelper struct {
}

// GetCipher
func (a *AesHelper) GetCipher(key string, iv string) (cipher.Block, error) {
	return aes.NewCipher([]byte(key))
}

// 加密
func (a *AesHelper) AesBase64Encrypt(cip cipher.Block, iv string, in string) (string, error) {
	origData := []byte(in)
	origData = PKCS5Padding(origData, cip.BlockSize())
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	bm := cipher.NewCBCEncrypter(cip, []byte(iv))
	bm.CryptBlocks(crypted, origData)
	var b = base64.StdEncoding.EncodeToString(crypted)
	return b, nil
}

// 解密
func (a *AesHelper) AesBase64Decrypt(cip cipher.Block, iv string, b string) (string, error) {
	crypted, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		return "", err
	}
	origData := make([]byte, len(crypted))
	bm := cipher.NewCBCDecrypter(cip, []byte(iv))
	bm.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	var out = string(origData)
	return out, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
