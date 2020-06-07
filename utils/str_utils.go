//author xinbing
//time 2018/8/28 14:18
//字符串工具
package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"
)

var randomStrSource = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//获取随机字符串
func GetRandomStr(length int) string {
	result := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63())) //增大随机性
	for i := 0; i < length; i++ {
		result[i] = randomStrSource[r.Intn(len(randomStrSource))]
	}
	return string(result)
}

//生成纯数字的随机字符串
func GetRandomNumStr(length int) string {
	result := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63())) //增大随机性
	for i := 0; i < length; i++ {
		result[i] = byte('0' + r.Intn(10)) //0 - 9
	}
	return string(result)
}

func UUID() (string, error) {
	id, err := uuid.NewV1()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// bool to int64(1,0)
func Uint8ToBool(data uint8) bool {
	if data > 0 {
		return true
	}
	return false
}

func ToString(i interface{}) string {
	if b, err := json.Marshal(i); err == nil {
		return string(b)
	}
	return "转换失败..."
}

func MakeUidWithTime(prefix string, random_len int) string {
	const time_base_format = "060102030405"
	uid := prefix + time.Now().Format(time_base_format) + GetRandomNumStr(random_len)
	return uid
}

func GetSha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
