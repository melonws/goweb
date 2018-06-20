package mqOrderCourse

import (
	"github.com/melonws/goweb/libs/mqHelper"
	"github.com/streadway/amqp"
	"github.com/melonws/goweb/libs/logHelper"
	"fmt"
	"encoding/json"
	"bytes"
	"github.com/melonws/goweb/models/orderCourse"
	"strconv"
)

type message struct {
	Type string `json:"type" from:"type"`
	Id int `json:"id" from:"id"`
}

func Start(count int){
	config := map[string]string{
		"exchange":"bizcrm-es",
		"routingKey":"bizcrm-es-sync-data-wangshu2",
		"queue":"bizcrm-es-sync-data-wangshu6",
	}
	callback := mqCallBack

	//mqHelperg.NewService(config,callback,count).Run()
	mqHelperg.Receive(config,callback,count)
}


func mqCallBack(delivery amqp.Delivery){

	//构造一个数据对象 最上方 type UserReg
	msg := new(message)

	//把delivery中的数据解析到构造的对象u中 u数据类型是UserReg
	err := json.Unmarshal(delivery.Body,msg)

	//再转成一个json字符串为了写日志用 jsonData数据类型是 string
	jsonData,_ :=json.Marshal(msg)


	if err != nil {
		//异常情况
		fmt.Println("异常了",err)
	}


	//这附近都可以拿到数据做数据的业务处理
	if msg.Type == "course_order_info" {
		result := orderCourse.AddToEs(msg.Id)
		logHelper.WriteLog("处理成课数据:"+*BytesToString(&jsonData)+",结果"+strconv.FormatBool(result),"mq/receive/access")
	}



	//处理完要把消费Ack 这样才会进来新的消息,但是整体底层还缺少了异常处理机制，对于异常数据也同样需要Ack否则会阻塞后续的正常数据
	//待优化完善
	delivery.Ack(true)


}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}