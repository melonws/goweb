package logHelper

import (
	"os"
	"fmt"
	"time"
	"log"
)

/**
 * 写日志方法
 * 最终把参数1 要写入的内容
 * 记录在参数2 当前脚本运行目录下logs文件夹下/filepath目录里，通过日期分割日志
 */
func WriteLog(message string , filePath string){

	//获取日志文件名绝对路径
	logFile := setPath(filePath)

	have,err := PathExists(logFile)
	if err !=nil {
		fmt.Println(err)
	}

	if !have {
		result := createPath(logFile)
		log.Println("日志目录创建:",result)
	}

	//向文件写入数据
	write(message,logFile)
}

/**
 * 文件读写操作
 */
func write(message string ,logFile string){

	currentDate := time.Now().Format("2006-01-02")

	logFile = logFile + "/" + currentDate + ".log"

	logStream, err := os.OpenFile(logFile, os.O_RDWR | os.O_CREATE |os.O_APPEND, 0777)

	if err != nil {
		fmt.Printf("open file error=%s\r\n", err.Error())
		os.Exit(-1)
	}

	defer logStream.Close()

	logger := log.New(logStream,"\r\n",log.Ldate | log.Ltime | log.Lshortfile)

	logger.Println(message)
}

/**
 * 获取脚本当前目录拼接日志目录
 */
func setPath(filePath string) string {

	var currentDir string

	currentDir, err := os.Getwd()

	if err != nil {

	}

	var logFile string

	logFile = currentDir + "/logs"

	if filePath != "" {
		logFile = currentDir + "/logs/" + filePath
	}

	return logFile
}

/**
 * 判断目录是否存在
 */
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/**
 * 创建目录
 */
func createPath(path string) bool {
	err := os.MkdirAll(path,os.ModePerm)

	if err != nil {
		return false
	}

	return true
}