package Config

import "goweb/config/loadConfig"

/**
 * 这个Config的包的目的是为了统一查看管理
 */

var DBConfig = map[string]interface{}{

	"Default":         "mysql_dev",// 默认数据库配置
	"SetMaxOpenConns": 20,          // (连接池)最大打开的连接数，默认值为0表示不限制
	"SetMaxIdleConns": 10,          // (连接池)闲置的连接数, 默认1

	"Connections":map[string]map[string]string{
		"mysql_dev": {// 定义名为 mysql_dev 的数据库配置
			"host": loadConfig.Get("MySql","DB_HOST",""), 	// 数据库地址
			"username": loadConfig.Get("MySql","DB_USERNAME",""),       // 数据库用户名
			"password": loadConfig.Get("MySql","DB_PASSWORD",""),       // 数据库密码
			"port": loadConfig.Get("MySql","DB_PORT","3306"),            // 端口
			"database": loadConfig.Get("MySql","DB_DATABASE",""),        // 链接的数据库名字
			"charset": "utf8",         // 字符集
			"protocol": "tcp",         // 链接协议
			"prefix": loadConfig.Get("MySql","DB_PREFIX",""),              // 表前缀
			"driver": "mysql",         // 数据库驱动(mysql,sqlite,postgres,oracle,mssql)
		},
		"sqlite_dev": {
			"database": "./foo.db",
			"prefix": "",
			"driver": "sqlite3",
		},
	},

}

var RedisConfig = map[string]interface{}{
	"Redis_Host" :loadConfig.Get("Redis","REDIS_HOST","10.99.2.21"),
	"Redis_Port" : loadConfig.Get("Redis","REDIS_PORT","6379"),
	"Redis_Pass" : nil,
}

var RabbitMqConfig = map[string]string{
	"RabbitMqHost" : loadConfig.Get("RabbitMq","MQ_HOST",""),
	"RabbitMqPort" : loadConfig.Get("RabbitMq","MQ_PORT","5672"),
	"RabbitMqUser" : loadConfig.Get("RabbitMq","MQ_USER",""),
	"RabbitMqPass" : loadConfig.Get("RabbitMq","MQ_PASS",""),
	"RabbitMqName" : loadConfig.Get("RabbitMq","MQ_NAME",""),
}

var EsConfig = map[string]string{
	"EsAddr" : loadConfig.Get("Es","ES_HOST",""),
	"EsIndex" : loadConfig.Get("Es","ES_INDEX",""),
}