package gotool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//AES加密,CBC
func AesEncrypt(text, key string) (string, error) {
	src := []byte(text)
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	src = pkcs7Padding(src, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, keyBytes[:blockSize])
	crypted := make([]byte, len(src))
	blockMode.CryptBlocks(crypted, src)
	return hex.EncodeToString(crypted), nil
}

//AES解密
func AesDecrypt(ciphertext, key string) (string, error) {
	crypted, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyBytes[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pkcs7UnPadding(origData)
	return string(origData), nil
}
