package dbHelper

import (
	"github.com/melonws/goweb/config"
	"github.com/gohouse/gorose"
	_"github.com/go-sql-driver/mysql"
	"log"
)


//创建一个model 后封装的。
func CreateModel(table string) (gorose.Connection,*gorose.Database,error) {

	//获取配置
	var dbConfig = Config.DBConfig
	//通过gorose包进行链接
	connection ,err := gorose.Open(dbConfig)

	if err != nil {
		log.Fatalln(err)
		return connection,nil,err
	}

	//链接后获取数据库
	db := connection.GetInstance()

	//创建对应的表的mode
	model  := db.Table(table)

	return connection,model,nil
}

