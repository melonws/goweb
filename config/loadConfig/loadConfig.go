package loadConfig

import (
	"github.com/Unknwon/goconfig"
	"os"
	"gorestful/libs/logHelper"
	"log"
)

/**
 * 通过env文件获取配置项
 */

var cfg *goconfig.ConfigFile

func init() {

	//获取配置文件的路径
	filePath := getConfigPath()

	config,err := goconfig.LoadConfigFile(filePath)

	if err != nil {
		logHelper.WriteLog("[读取配置文件出错]"+err.Error(),"system/error")
		log.Fatalln("[读取配置文件出错]"+err.Error())
	}

	cfg = config
}

/**
 * 获取当前脚本的路径下env文件
 */
func getConfigPath() string {

	var currentDir string

	//获取当前脚本运行的路径
	currentDir, err := os.Getwd()

	if err != nil {
		logHelper.WriteLog("[读取配置文件出错]"+err.Error(),"system/error")
		log.Fatalln("[读取配置文件出错]"+err.Error())
	}

	var logFile string

	logFile = currentDir + "/.env"

	return logFile
}

/**
 * 公共方法，可以通过该方法获取对应env里的配置
 */
func Get(selection string,name string,defaultValue string) string {

	value,err := cfg.GetValue(selection,name)

	if err != nil || value == "" {
		return defaultValue
	}

	return value
}