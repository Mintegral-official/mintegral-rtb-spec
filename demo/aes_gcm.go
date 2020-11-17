package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)
/*
	本代码仅作为demo使用
	在线上代码中， key将会支持可配置， 所有 err 也都会被正确处理
 */

func AESGCMEncode(plaintext, key string) string {
	fmt.Println("AESGCM encoding...")
	keyBytes, _ := base64.StdEncoding.DecodeString(key)
	fmt.Printf("workkey:%s, bytes:%v, length:%d\n", key, keyBytes, len(keyBytes))
	nonce := make([]byte, 12)
	io.ReadFull(rand.Reader, nonce) // random iv
	fmt.Printf("rand iv: %v, length: %d\n", nonce, len(nonce))
	block, _ := aes.NewCipher(keyBytes)
	aesgcm, _ := cipher.NewGCM(block)
	ciphertext := aesgcm.Seal(nil, nonce, []byte(plaintext), nil)
	ciphertext = append(nonce , ciphertext...)
	fmt.Printf("ciphertext before base64:%v, length %d\n", ciphertext, len(ciphertext))
	return base64.URLEncoding.EncodeToString(ciphertext)
}

func AESGCMDecode(ciphertext, key string) (plaintext string, err error) {
	fmt.Printf("AESGCM Decoding...  cipher:%s, withqkey:%s\n", ciphertext, key)
	keyBytes, _ := base64.StdEncoding.DecodeString(key)
	fullCipher, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	nonce, cipherBytes := fullCipher[:12], fullCipher[12:]
	fmt.Printf("nonce: %v, length %d\n", nonce, len(nonce))
	fmt.Printf("cipher: %v, length %d\n", cipherBytes, len(cipherBytes))
	block, _ := aes.NewCipher(keyBytes)
	aesgcm, _ := cipher.NewGCM(block)
	plaintextBytes, err := aesgcm.Open(nil, nonce, cipherBytes, nil)
	return string(plaintextBytes), err
}

// generateRandKey 生成一个随机的16bytekey, 并进行base64
func generateRandKey() string {
	key := make([]byte, 16)
	io.ReadFull(rand.Reader, key)
	return base64.StdEncoding.EncodeToString(key)
}

func main() {
	workKey := generateRandKey()
	price := "111111"
	ret := AESGCMEncode(price, workKey)
	raw,err := AESGCMDecode(ret, workKey)
	if err != nil {
		fmt.Println("AESGCM decode err:" ,err.Error())
	} else {
		fmt.Printf("AESGCM decode result: %s, length:%d \n", raw, len(raw))
	}
	fmt.Println("result should be true:", raw == price)


	/*如下是某次执行的console print， 供参考
	AESGCM encoding...
	workkey:z2n31d6lo2Pdfe4J/a9ClA==, bytes:[207 105 247 213 222 165 163 99 221 125 238 9 253 175 66 148], length:16
	rand iv: [122 232 163 28 122 195 138 104 51 94 142 173], length: 12
	ciphertext before base64:[122 232 163 28 122 195 138 104 51 94 142 173 223 87 207 65 157 240 213 215 8 232 126 253 207 171 106 166 90 146 104 27 201 33], length 34
	AESGCM Decoding...  cipher:euijHHrDimgzXo6t31fPQZ3w1dcI6H79z6tqplqSaBvJIQ==, withqkey:z2n31d6lo2Pdfe4J/a9ClA==
	nonce: [122 232 163 28 122 195 138 104 51 94 142 173], length 12
	cipher: [223 87 207 65 157 240 213 215 8 232 126 253 207 171 106 166 90 146 104 27 201 33], length 22
	AESGCM decode result: 111111, length:6
	result should be true: true
	 */
}

