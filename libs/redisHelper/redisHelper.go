package redisHelper

import (
	"github.com/garyburd/redigo/redis"
	"../../config"
	"../../libs/logHelper"
	"strconv"
)

func connect() (redis.Conn,error) {
	config := Config.RedisConfig

	redisHost := config["Redis_Host"].(string) + ":" + config["Redis_Port"].(string)

	var dailOption redis.DialOption
	var err error
	var connection redis.Conn

	if config["Redis_Pass"] != nil {

		dailOption = redis.DialPassword(config["Redis_Pass"].(string))

		connection,err = redis.Dial("tcp",redisHost,dailOption)

	}else{
		connection,err = redis.Dial("tcp",redisHost)
	}

	if err != nil {
		return nil,err
	}

	return connection,nil
}

//time 传0 代表无限 单位秒
func Set(key string,value string , time int) bool {

	var err error

	connection,err  := connect()

	if err != nil {
		data := "key:" + key + ", value:" + value + ", time:" + strconv.Itoa(time)+"s"
		logHelper.WriteLog("redis connect error sourceData:" + data,"redis/error")
		return false
	}

	defer connection.Close()


	_,err = connection.Do("SET" , key , value ,"EX",strconv.Itoa(time))

	if err != nil {
		data := "key:" + key + ", value:" + value + ", time:" + strconv.Itoa(time)+"s"
		logHelper.WriteLog("redis SET error sourceData:" + data,"redis/error")
		return false
	}
	return true
}


func Get(key string) string {

	var err error

	connection,err := connect()

	if err != nil {
		data := "key:" + key
		logHelper.WriteLog("redis connect error sourceData:" + data,"redis/error")
		return ""
	}

	defer connection.Close()

	data,err := redis.String(connection.Do("GET",key))

	if err != nil {
		data := "key:" + key
		logHelper.WriteLog("redis GET error sourceData:" + data,"redis/error")
		return ""
	}

	return data
}

//todo 继续封装 https://studygolang.com/articles/7104