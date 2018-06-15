package cmd_custom_reg

import (
	"goweb/libs/mqHelper"
	"encoding/json"
	"fmt"
	"goweb/libs/logHelper"
	"github.com/streadway/amqp"
	"bytes"
	"time"
)

type UserReg struct {
	UId  int `json:"uid,string"`
	UserId  int `json:"user_id,string"`
}

/**
 * 脚本启动函数，建议统一命名
 * 设置了配置文件、回调函数、然后就调起Mq的调度器创建任务开始执行
 */
func Start(count int){
	config := map[string]string{
		"exchange":"bizcrm-customer",
		"routingKey":"bizcrm-customer-reg",
		"queue":"bizcrm-customer-reg",
	}
	callback := mqCallBack

	mqHelperg.NewService(config,callback,count).Run()
}

/**
 * 回调函数，也就是业务逻辑
 * 不管底层的话 就照着这个文件写就行
 */
func mqCallBack(delivery amqp.Delivery){

	//构造一个数据对象 最上方 type UserReg
	u := new(UserReg)

	//把delivery中的数据解析到构造的对象u中 u数据类型是UserReg
	err := json.Unmarshal(delivery.Body,u)

	//再转成一个json字符串为了写日志用 jsonData数据类型是 string
	jsonData,_ :=json.Marshal(u)


	if err != nil {
		//异常情况
		fmt.Println("异常了")
	}


	logHelper.WriteLog("请求到了\n"+*BytesToString(&jsonData),"mq/receive")

	//这附近都可以拿到数据做数据的业务处理
	time.Sleep(1 * time.Second)


	logHelper.WriteLog("休息完了"+*BytesToString(&jsonData),"mq/receive")

	//处理完要把消费Ack 这样才会进来新的消息,但是整体底层还缺少了异常处理机制，对于异常数据也同样需要Ack否则会阻塞后续的正常数据
	//待优化完善
	delivery.Ack(true)


}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}