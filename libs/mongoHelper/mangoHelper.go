package mongoHelper

import (
	"gopkg.in/mgo.v2"
	"github.com/melonws/goweb/libs/logHelper"
	"time"
)

/**
 * 建立数据库链接
 */
func connect() *mgo.Session {

	mongoDialInfo := &mgo.DialInfo{
		Addrs:[]string{"localhost"},
		Timeout:10 * time.Second,
		Database:"testdb",
		Username:"",
		Password:"" }

	session,err := mgo.DialWithInfo(mongoDialInfo)

	if err != nil {
		logHelper.WriteLog("mango 链接失败"+err.Error(),"mangodb/error")
		return nil
	}

	return session
}

/**
 * 通过数据库连接后，拿到对应的session(会话) 和 collection(集合)
 * 便于操作
 * 返回后的操作，记得先 defer session.Close()
 */
func CreateModel(collectionName string) (*mgo.Session,*mgo.Collection) {

	session := connect()

	if session == nil {
		return nil,nil
	}

	session.SetMode(mgo.Monotonic,true)

	collection := session.DB("testdb").C(collectionName)

	return session,collection
}