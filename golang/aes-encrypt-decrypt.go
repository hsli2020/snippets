// 原文链接：https://blog.csdn.net/qq_25504271/article/details/79454648
package main

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "fmt"
)

const (
    key = "2018201820182018"
    iv  = "1234567887654321"
)

func main() {
    str := "我勒个去"
    es, _ := AesEncrypt(str, []byte(key))
    fmt.Println(es)

    ds, _ := AesDecrypt(es, []byte(key))
    fmt.Println(string(ds))
}

func AesEncrypt(encodeStr string, key []byte) (string, error) {
    encodeBytes := []byte(encodeStr)
    //根据key 生成密文
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    blockSize := block.BlockSize()
    encodeBytes = PKCS5Padding(encodeBytes, blockSize)

    blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
    crypted := make([]byte, len(encodeBytes))
    blockMode.CryptBlocks(crypted, encodeBytes)

    return base64.StdEncoding.EncodeToString(crypted), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    //填充
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)

    return append(ciphertext, padtext...)
}

func AesDecrypt(decodeStr string, key []byte) ([]byte, error) {
    //先解密base64
    decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
    if err != nil {
        return nil, err
    }
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
    origData := make([]byte, len(decodeBytes))

    blockMode.CryptBlocks(origData, decodeBytes)
    origData = PKCS5UnPadding(origData)
    return origData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
    length := len(origData)
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}
