package autn

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

type DesHelper struct {
}

//pkcs5补码算法
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//pkcs5减码算法
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//zero补码算法
func (a *DesHelper) ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

//zero减码算法
func (a *DesHelper) ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

//---------------DES ECB加密--------------------
// data: 明文数据
// key: 密钥字符串
// 返回密文数据
func (a *DesHelper) DesECBEncrypt(data, key []byte) []byte {
	//NewCipher创建一个新的加密块
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	bs := block.BlockSize()
	// pkcs5填充
	data = pkcs5Padding(data, bs)
	if len(data)%bs != 0 {
		return nil
	}

	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//Encrypt加密第一个块，将其结果保存到dst
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return out
}

//---------------DES ECB解密--------------------
// data: 密文数据
// key: 密钥字符串
// 返回明文数据
func (a *DesHelper) DesECBDecrypter(data, key []byte) []byte {
	//NewCipher创建一个新的加密块
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return nil
	}

	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//Encrypt加密第一个块，将其结果保存到dst
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}

	// pkcs5填充
	out = pkcs5UnPadding(out)

	return out
}

//---------------DES CBC加密--------------------
// data: 明文数据
// key: 密钥字符串
// iv:iv向量
// 返回密文数据
func (a *DesHelper) DesCBCEncrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	// pkcs5填充
	data = pkcs5Padding(data, block.BlockSize())
	cryptText := make([]byte, len(data))

	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(cryptText, data)
	return cryptText
}

//---------------DES CBC解密--------------------
// data: 密文数据
// key: 密钥字符串
// iv:iv向量
// 返回明文数据
func (a *DesHelper) DesCBCDecrypter(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	cryptText := make([]byte, len(data))
	blockMode.CryptBlocks(cryptText, data)
	// pkcs5填充
	cryptText = pkcs5UnPadding(cryptText)

	return cryptText
}

//---------------DES CTR加密--------------------
// data: 明文数据
// key: 密钥字符串
// iv:iv向量
// 返回密文数据
func (a *DesHelper) DesCTREncrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	// pkcs5填充
	data = pkcs5Padding(data, block.BlockSize())
	cryptText := make([]byte, len(data))

	blockMode := cipher.NewCTR(block, iv)
	blockMode.XORKeyStream(cryptText, data)
	return cryptText
}

//---------------DES CTR解密--------------------
// data: 密文数据
// key: 密钥字符串
// iv:iv向量
// 返回明文数据
func (a *DesHelper) DesCTRDecrypter(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	blockMode := cipher.NewCTR(block, iv)
	cryptText := make([]byte, len(data))
	blockMode.XORKeyStream(cryptText, data)

	// pkcs5填充
	cryptText = pkcs5UnPadding(cryptText)

	return cryptText
}

//---------------DES OFB加密--------------------
// data: 明文数据
// key: 密钥字符串
// iv:iv向量
// 返回密文数据
func (a *DesHelper) DesOFBEncrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	// pkcs5填充
	data = pkcs5Padding(data, block.BlockSize())
	cryptText := make([]byte, len(data))

	blockMode := cipher.NewOFB(block, iv)
	blockMode.XORKeyStream(cryptText, data)
	return cryptText
}

//---------------DES OFB解密--------------------
// data: 密文数据
// key: 密钥字符串
// iv:iv向量
// 返回明文数据
func (a *DesHelper) DesOFBDecrypter(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	blockMode := cipher.NewOFB(block, iv)
	cryptText := make([]byte, len(data))
	blockMode.XORKeyStream(cryptText, data)

	// pkcs5填充
	cryptText = pkcs5UnPadding(cryptText)

	return cryptText
}

//---------------DES CFB加密--------------------
// data: 明文数据
// key: 密钥字符串
// iv:iv向量
// 返回密文数据
func (a *DesHelper) DesCFBEncrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	// pkcs5填充
	data = pkcs5Padding(data, block.BlockSize())
	cryptText := make([]byte, len(data))

	blockMode := cipher.NewCFBDecrypter(block, iv)
	blockMode.XORKeyStream(cryptText, data)
	return cryptText
}

//---------------DES CFB解密--------------------
// data: 密文数据
// key: 密钥字符串
// iv:iv向量
// 返回明文数据
func (a *DesHelper) DesCFBDecrypter(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	blockMode := cipher.NewCFBEncrypter(block, iv)
	cryptText := make([]byte, len(data))
	blockMode.XORKeyStream(cryptText, data)

	// pkcs5填充
	cryptText = pkcs5UnPadding(cryptText)

	return cryptText
}
