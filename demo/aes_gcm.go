package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)
var key = "0123456789ABCDEF0123456789ABCDEF" // 32 bit hex
var nonceStr = "000000000000000000000000"  // hex decode的结果是一个12位的[]byte,
func AESGCMEncode(plaintext, key string) string {
	nonce, _ := hex.DecodeString(nonceStr)
	block, _ := aes.NewCipher([]byte(key))
	aesgcm, _ := cipher.NewGCM(block)
	ciphertext := aesgcm.Seal(nil, nonce, []byte(plaintext), nil)
	return base64.URLEncoding.EncodeToString(ciphertext)
}

func AESGCMDecode(ciphertext, key string) string {
	cipherStr, _ := base64.URLEncoding.DecodeString(ciphertext)
	nonce, _ := hex.DecodeString(nonceStr)
	block, _ := aes.NewCipher([]byte(key))
	aesgcm, _ := cipher.NewGCM(block)
	plaintext, _ := aesgcm.Open(nil, nonce, cipherStr, nil)
	return string(plaintext)
}

func main() {

	price := "4.56"
	cipher := AESGCMEncode(price, key)
	fmt.Println(cipher)
	raw := AESGCMDecode(cipher, key)
	fmt.Println(raw)
}

