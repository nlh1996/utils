package utils

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"log"
	"net/url"
	"strings"
)

// Jiami md5加密
func Jiami(arg1 string, arg2 string) string {
	h := md5.New()
	h.Write([]byte(arg1))
	result := fmt.Sprintf("%x", h.Sum(nil))
	h.Write([]byte(result + arg2))
	result = fmt.Sprintf("%x", h.Sum(nil))
	h.Write([]byte(result + "你妈嗨"))
	result = fmt.Sprintf("%x", h.Sum(nil))
	return result
}

// MD5 .
func MD5(arg string) string {
	h := md5.New()
	h.Write([]byte(arg))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// URLDecode .
func URLDecode(d string) string {
	res, err := url.QueryUnescape(d)
	if err != nil {
		log.Println(err)
		return ""
	}
	return res
}

// Base64Decode .
func Base64Decode(d string) string {
	res, err := base64.StdEncoding.DecodeString(d)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(res)
}

// GenRsaKey RSA公钥私钥产生
func GenRsaKey() (prvkey, pubkey []byte) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prvkey = pem.EncodeToMemory(block)
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubkey = pem.EncodeToMemory(block)
	return
}

// RsaSignWithSha256 签名
func RsaSignWithSha256(data []byte, keyBytes []byte) []byte {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("private key error"))
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey err", err)
		panic(err)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		panic(err)
	}

	return signature
}

// RsaVerySignWithSha256 验证
func RsaVerySignWithSha256(data, signData, keyBytes []byte) bool {
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signData)
	if err != nil {
		panic(err)
	}
	return true
}

// RsaEncrypt 公钥加密
func RsaEncrypt(data, keyBytes []byte) []byte {
	//解密pem格式的公钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		panic(err)
	}
	return ciphertext
}

// RsaDecrypt 私钥解密
func RsaDecrypt(ciphertext, keyBytes []byte) []byte {
	// 获取私钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("private key error!"))
	}
	// 解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		panic(err)
	}
	return data
}

// VerifyRsaSign .
func VerifyRsaSign(content string, sign string, publicKey string) error {
	publicKeyByte, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return err
	}
	pub, err := x509.ParsePKIXPublicKey(publicKeyByte)
	if err != nil {
		return err
	}
	hashed := sha256.Sum256([]byte(content))
	signature, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, hashed[:], signature)
}

// VerifyRsa1Sign .
func VerifyRsa1Sign(content string, sign string, publicKey string) error {
	publicKeyByte, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return err
	}
	pub, err := x509.ParsePKIXPublicKey(publicKeyByte)
	if err != nil {
		return err
	}
	h := crypto.Hash.New(crypto.SHA1) //进行SHA1的散列
	h.Write([]byte(content))
	hashed := h.Sum(nil)
	signature, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA1, hashed[:], signature)
}

func verifySign(signData, sign, signType, aliPayPublicKey string) (err error) {
	var (
		h         hash.Hash
		hashs     crypto.Hash
		block     *pem.Block
		pubKey    interface{}
		publicKey *rsa.PublicKey
		ok        bool
	)
	signBytes, _ := base64.StdEncoding.DecodeString(sign)
	if block, _ = pem.Decode([]byte(aliPayPublicKey)); block == nil {
		return errors.New("支付宝公钥Decode错误")
	}
	if pubKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		return fmt.Errorf("x509.ParsePKIXPublicKey：%w", err)
	}
	if publicKey, ok = pubKey.(*rsa.PublicKey); !ok {
		return errors.New("支付宝公钥转换错误")
	}
	switch signType {
	case "RSA":
		hashs = crypto.SHA1
	case "RSA2":
		hashs = crypto.SHA256
	default:
		hashs = crypto.SHA256
	}
	h = hashs.New()
	h.Write([]byte(signData))
	return rsa.VerifyPKCS1v15(publicKey, hashs, h.Sum(nil), signBytes)
}

// FormatPublicKey .
func FormatPublicKey(publicKey string) (pKey string) {
	var buffer strings.Builder
	buffer.WriteString("-----BEGIN PUBLIC KEY-----\n")
	rawLen := 64
	keyLen := len(publicKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(publicKey[start:])
		} else {
			buffer.WriteString(publicKey[start:end])
		}
		buffer.WriteByte('\n')
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END PUBLIC KEY-----\n")
	pKey = buffer.String()
	return
}
