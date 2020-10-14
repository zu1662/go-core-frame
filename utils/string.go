package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// StringToInt64 StringToInt64
func StringToInt64(e string) (int64, error) {
	return strconv.ParseInt(e, 10, 64)
}

// StringToInt StringToInt
func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}

// StringToBool StringToBool
func StringToBool(e string) (bool, error) {
	return strconv.ParseBool(e)
}

// GetCurrentTimeStr GetCurrentTimeStr
func GetCurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetCurrentTime GetCurrentTime
func GetCurrentTime() time.Time {
	return time.Now()
}

// FormatTimeStr FormatTimeStr
func FormatTimeStr(timeStr string) (string, error) {
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation("2006-01-02T15:04:05.000Z", timeStr, loc)
	return theTime.Format("2006/01/02 15:04:05"), err
}

// StructToJsonStr StructToJsonStr
func StructToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

// GetBodyString GetBodyString
func GetBodyString(c *gin.Context) (string, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return string(body), nil
	} else {
		return "", err
	}
}

// JsonStrToMap JsonStrToMap
func JsonStrToMap(e string) (map[string]interface{}, error) {
	var dict map[string]interface{}
	if err := json.Unmarshal([]byte(e), &dict); err == nil {
		return dict, err
	} else {
		return nil, err
	}
}

// StructToMap StructToMap
func StructToMap(data interface{}) (map[string]interface{}, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}

// MobileSecurity 手机号中间加密
// 19951266666   --->  199****6666
func MobileSecurity(mobile string) string {
	if len(mobile) < 10 {
		return mobile
	}
	secMobile := mobile[:3] + "****" + mobile[6:]
	return secMobile
}

// GetSHA256HashCode SHA256生成哈希值
func GetSHA256HashCode(message []byte) string {
	//方法一：
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write(message)
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode
}

// GetMD5HashCode MD5生成哈希值
func GetMD5HashCode(message []byte) string {
	//方法一：
	//创建一个基于md5算法的hash.Hash接口的对象
	hash := md5.New()
	//输入数据
	hash.Write(message)
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode
}

// GetUUID 获取UUID
func GetUUID() string {
	uid, _ := uuid.NewV4()
	return uid.String()
}
