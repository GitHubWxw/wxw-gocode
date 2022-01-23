/*
@Time : 2022/1/23 15:01
@Author : weixiaowei
@File : aes_key16
*/
package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

// 改动点：
//  1. 将base64编码改为hex编码 原因是 通过程序参数执行 base64解码时位数不一致导致解密失败
// 相关文章
// 1. https://blog.csdn.net/u014270740/article/details/91038606

func main() {
	//m := md5.Sum([]byte("cvrhsdftredhghgfjhgwsfresdsfhjk"))
	//key := hex.EncodeToString(m[:])[0:16]

	key := "0123456789ABCDEF"
	content := "weixiaowei@qoogle.com"

	// 加密
	encryptResult := AesEncrypt02(content, key)
	fmt.Println("加密后：", encryptResult)

	decryptResult := AesDecrypt02(encryptResult, key)
	fmt.Println("解密后：", decryptResult)

}

// AES 加密
func AesEncrypt02(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)

	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding02(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return hex.EncodeToString(cryted)
}

// AES 解密
func AesDecrypt02(crypto string, key string) string {
	// 转成字节数组
	// cryptoByte, _ := base64.StdEncoding.DecodeString(crypto) //
	cryptoByte, _ := hex.DecodeString(crypto)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(cryptoByte))
	// 解密
	blockMode.CryptBlocks(orig, cryptoByte)
	// 去补全码
	orig = PKCS7UnPadding02(orig)
	return string(orig)
}

//补码
func PKCS7Padding02(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

//去码
func PKCS7UnPadding02(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
