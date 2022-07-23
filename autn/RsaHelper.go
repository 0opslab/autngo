package autn

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io"
	"os"
)

type RsaHelper struct {
	fs FileHelper
}

// 私钥生成
//openssl genrsa -out rsa_private_key.pem 1024
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQDcGsUIIAINHfRTdMmgGwLrjzfMNSrtgIf4EGsNaYwmC1GjF/bM
h0Mcm10oLhNrKNYCTTQVGGIxuc5heKd1gOzb7bdTnCDPPZ7oV7p1B9Pud+6zPaco
qDz2M24vHFWYY2FbIIJh8fHhKcfXNXOLovdVBE7Zy682X1+R1lRK8D+vmQIDAQAB
AoGAeWAZvz1HZExca5k/hpbeqV+0+VtobMgwMs96+U53BpO/VRzl8Cu3CpNyb7HY
64L9YQ+J5QgpPhqkgIO0dMu/0RIXsmhvr2gcxmKObcqT3JQ6S4rjHTln49I2sYTz
7JEH4TcplKjSjHyq5MhHfA+CV2/AB2BO6G8limu7SheXuvECQQDwOpZrZDeTOOBk
z1vercawd+J9ll/FZYttnrWYTI1sSF1sNfZ7dUXPyYPQFZ0LQ1bhZGmWBZ6a6wd9
R+PKlmJvAkEA6o32c/WEXxW2zeh18sOO4wqUiBYq3L3hFObhcsUAY8jfykQefW8q
yPuuL02jLIajFWd0itjvIrzWnVmoUuXydwJAXGLrvllIVkIlah+lATprkypH3Gyc
YFnxCTNkOzIVoXMjGp6WMFylgIfLPZdSUiaPnxby1FNM7987fh7Lp/m12QJAK9iL
2JNtwkSR3p305oOuAz0oFORn8MnB+KFMRaMT9pNHWk0vke0lB1sc7ZTKyvkEJW0o
eQgic9DvIYzwDUcU8wJAIkKROzuzLi9AvLnLUrSdI6998lmeYO9x7pwZPukz3era
zncjRK3pbVkv0KrKfczuJiRlZ7dUzVO0b6QJr8TRAA==
-----END RSA PRIVATE KEY-----
`)

// 公钥: 根据私钥生成
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDcGsUIIAINHfRTdMmgGwLrjzfM
NSrtgIf4EGsNaYwmC1GjF/bMh0Mcm10oLhNrKNYCTTQVGGIxuc5heKd1gOzb7bdT
nCDPPZ7oV7p1B9Pud+6zPacoqDz2M24vHFWYY2FbIIJh8fHhKcfXNXOLovdVBE7Z
y682X1+R1lRK8D+vmQIDAQAB
-----END PUBLIC KEY-----
`)

// 公钥加密私钥解密
func (en *RsaHelper) RsaEncryptWithPublic(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		{
			return nil, ErrorMsg("public key error", nil)
		}
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)

}

//公钥加密私钥解密文件
func (en *RsaHelper) RsaEncryptFileWithPublic(src, dst string) (bool, error) {

	r1, err0 := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err0 != nil {
		return false, err0
	}
	defer r1.Close()

	w1, err1 := en.fs.CreateOpenFile(dst)
	if err1 != nil {
		return false, ErrorMsg("CreateFileError:"+dst, err0)
	}
	defer w1.Close()

	bufReader := bufio.NewReader(r1)
	buf := make([]byte, 100)
	for {
		readNum, err := bufReader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return false, err
		}
		rsdata := buf[:readNum]
		rsaData, err := en.RsaEncryptWithPublic(rsdata)
		if err != nil {
			return false, err
		}
		w1.Write(rsaData)

	}
	return true, nil
}

// 公钥加密私钥解密
func (en *RsaHelper) RsaDecryptWithPrivte(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, ErrorMsg("private key error!", nil)
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

//公钥加密私钥解密文件
func (en *RsaHelper) RsaDecryptFileWithPrivte(src, dst string) (bool, error) {

	r1, err0 := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err0 != nil {
		return false, ErrorMsg("OpenFileError:"+src, err0)
	}
	defer r1.Close()

	w1, err1 := en.fs.CreateOpenFile(dst)
	if err1 != nil {
		return false, ErrorMsg("CreateFileError:"+dst, err1)
	}
	defer w1.Close()

	bufReader := bufio.NewReader(r1)
	buf := make([]byte, 128)
	for {
		readNum, err := bufReader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return false, err
		}
		rsdata := buf[:readNum]
		rsaData, err := en.RsaDecryptWithPrivte(rsdata)
		if err != nil {
			return false, err
		}
		w1.Write(rsaData)
	}
	return true, nil
}
