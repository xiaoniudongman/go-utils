package security

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"

	"github.com/xndm-recommend/go-utils/errs"
)

// 用aes加密算法对数据进行加密，加密模式为ecb，加密之后用base64进行编码
func AesEcbEncryptBASE64(str2encode, key string) string {
	byte2encode := PKCS7Pad([]byte(str2encode))
	encryptByte := encrypt(byte2encode, key)
	return base64.StdEncoding.EncodeToString(encryptByte)
}

// 对加密数据进行解码，(aes，加密模式ecb，base64编码）
func AesEcbDecryptBASE64(encryptData, key string) string {
	encryptByte, err := base64.StdEncoding.DecodeString(encryptData)
	errs.CheckCommonErr(err)
	return string(decrypt(encryptByte, key))
}

// 加密
func encrypt(plaintext []byte, key string) []byte {
	cipher, err := aes.NewCipher([]byte(key[:aes.BlockSize]))
	if err != nil {
		panic(err.Error())
	}

	if len(plaintext)%aes.BlockSize != 0 {
		panic("Need a multiple of the blocksize 16")
	}

	ciphertext := make([]byte, 0)
	text := make([]byte, 16)
	for len(plaintext) > 0 {
		// 每次运算一个block
		cipher.Encrypt(text, plaintext)
		plaintext = plaintext[aes.BlockSize:]
		ciphertext = append(ciphertext, text...)
	}
	return ciphertext
}

// 解密
func decrypt(ciphertext []byte, key string) []byte {
	cipher, err := aes.NewCipher([]byte(key[:aes.BlockSize]))
	if err != nil {
		panic(err.Error())
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("Need a multiple of the blocksize 16")
	}

	plaintext := make([]byte, 0)
	text := make([]byte, 16)
	for len(ciphertext) > 0 {
		cipher.Decrypt(text, ciphertext)
		ciphertext = ciphertext[aes.BlockSize:]
		plaintext = append(plaintext, text...)
	}
	return plaintext
}

// Padding补全
func PKCS7Pad(data []byte) []byte {
	padding := aes.BlockSize - len(data)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func PKCS7UPad(data []byte) []byte {
	padLength := int(data[len(data)-1])
	return data[:len(data)-padLength]
}
