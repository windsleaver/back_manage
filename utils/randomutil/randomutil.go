package randomutil

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * @Time   : 2022/2/24 00:40
 * @Author : WindsLeaver
 * @File   : randomutil
 * @Version: 1.0.0
 * @Description:
 **/

// GetRandomSixCode 随机数6位验证码
func GetRandomSixCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", r.Int31n(1000000))
	return code
}

// GetRandomString 生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

// GetRandomNumString 生成随机数字字符串
func GetRandomNumString(length int) string {
	str := "0123456789"
	bytes := []byte(str)
	var result []byte
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

// GetRandomBase32String 生成base32随机密钥
func GetRandomBase32String(length int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	bytes := []byte(str)
	var result []byte
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
