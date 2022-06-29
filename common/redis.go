package common

import (
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"log"
)

var redisConn redis.Conn

func InitRedis() redis.Conn {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")

	password := viper.GetString("password")

	con, err := redis.Dial("tcp", host+":"+port)

	if err != nil {
		log.Println("redis 创建链接异常")
		panic(err)
	}

	if len(password) != 0 {
		_, err := con.Do("Auth", password)
		if err != nil {
			log.Println("Redis 鉴权失败！")
			panic(err)
		}
	}

	//defer func(con redis.Conn) {
	//	err := con.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(con)
	redisConn = con
	return con
}

// SetStringValue 向Redis存入字符类型数据
// key 键
// value 值
func SetStringValue(key string, value string) bool {
	_, err := redisConn.Do("set", key, value)
	if err != nil {
		panic(err)
		return false
	}
	return true
}

// GetStringValue 根据 key 获取 value
func GetStringValue(key string) string {
	if !HasKey(key) {
		return ""
	}

	value, err := redisConn.Do("get", key)
	if err != nil {
		panic(err)
	}
	return Strval(value)
}

// HasKey key 是否存在，存在 return true；否则 return false
func HasKey(key string) bool {
	hasKey, err := redis.Bool(redisConn.Do("EXISTS", key))
	if err != nil {
		panic(err)
		return false
	}
	return hasKey
}

// SetStringValueExpire 存储value 设置过期时间
// key 键
// value 值
// expirationTime 过期时间 （秒）
func SetStringValueExpire(key string, value string, expirationTime int) {
	if SetStringValue(key, value) {
		SetKeyExpire(key, expirationTime)
	}
}

func SetKeyExpire(key string, expirationTime int) {
	_, err := redisConn.Do("EXPIRE", key, expirationTime)
	if err != nil {
		panic(err)
	}
}

func SetJSONValue(key string, value map[string]string) bool {
	jsonValue := MapConversionJSONByte(value)
	n, err := redisConn.Do("SETNX", jsonValue)
	if err != nil {
		log.Println(err)
	}
	if n == int64(1) {
		return true
	}
	return false
}

func GetJSONValue(key string) map[string]string {
	value, err := redis.Bytes(redisConn.Do("GET", key))
	if err != nil {
		log.Println(err)
	}
	return JSONByteConversionMap(value)
}

func MapConversionJSONByte(value map[string]string) []byte {
	if len(value) < 0 {
		return nil
	}
	jsonValue, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return jsonValue
}

func JSONByteConversionMap(jsonValue []byte) map[string]string {
	var value map[string]string
	err := json.Unmarshal(jsonValue, &value)
	if err != nil {
		panic(err)
	}
	return value
}
