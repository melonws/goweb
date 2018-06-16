package mqHelperg

import (
	"log"
	"fmt"
	"github.com/streadway/amqp"
	"errors"
	"bytes"
	"github.com/melonws/goweb/config"
	"time"
)

var conn *amqp.Connection
var channel *amqp.Channel

var mqUrl string

func init() {
	config := Config.RabbitMqConfig
	mqUrl = fmt.Sprintf("amqp://%s:%s@%s:%s/%s",
		config["RabbitMqUser"], config["RabbitMqPass"], config["RabbitMqHost"],
			config["RabbitMqPort"], config["RabbitMqName"])
}


func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}


func mqConnect() {
	var err error
	conn, err = amqp.Dial(mqUrl)
	failOnErr(err, "failed to connect tp rabbitmq")

	channel, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
}



func Push(config map[string]string,msgContent string) (bool,error) {

	if channel == nil {
		mqConnect()
	}

	exchange,err := config["exchange"]
	if !err {
		return err,errors.New("没有exchange")
	}

	routingKey,err := config["routingKey"]
	if !err {
		return err,errors.New("没有routingKey")
	}

	errs := channel.Publish(exchange, routingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgContent),
	})

	if errs != nil {
		return false,errs
	}
	return true,nil
}

/**
 * 为了misson这个调度器重新写的
 * 进行mq连接
 * 连接好了以后设置每次消费的数量，根据协程数
 * 然后通过底层包读出数据 msgs
 * 并且写入另外一个传进来的无缓冲队列，该队列长度 = 协程数 = 读取消息数
 * 目的是跑指定个协程的时候数据就进不去了，会阻塞住，避免协程无限起
 */
func ReadMsg(config map[string]string,ds chan amqp.Delivery,count int) {
	if channel == nil {
		mqConnect()
	}

	channel.Qos(count,0,false)

	queueName,errs := config["queue"]

	if !errs {
		//return "",errors.New("没有queue")
	}

	msgs, err := channel.Consume(queueName, "", false, false, false, false, nil)

	failOnErr(err, "")

	forever := make(chan bool)

	for d := range msgs {
		ds <- d
	}

	<-forever

}


/**
 * 单进程、单条消费的接收函数
 * 最初的写法，后来因为协程的问题，重写了一个 ReadMsg
 */
func Receive(config map[string]string,callback func(delivery amqp.Delivery),count int) {
	if channel == nil {
		mqConnect()
	}
	channel.Qos(count,0,false)
	queueName,errs := config["queue"]
	if !errs {
		//return "",errors.New("没有queue")
	}

	msgs, err := channel.Consume(queueName, "", false, false, false, false, nil)
	failOnErr(err, "")

	//forever := make(chan bool)
loop :
	select { case d, ok := <- msgs :
		if ok {
			go callback(d)
			time.Sleep(5 * time.Millisecond)
		}
		goto loop
	}

	//for d := range msgs {
	//	fmt.Println(d)
	//	go callback(d)
	//	time.Sleep(5 * time.Millisecond)
	//	fmt.Println("d over")
	//}
	//
	//<-forever
}


func CreateQueue(config map[string]string) (amqp.Queue,error) {
	if channel == nil {
		mqConnect()
	}
	queue,err := channel.QueueDeclare(config["queue"],true,false,false,false,nil)

	if err != nil {
		return queue,err
	}

	return queue,nil
}

func BindQueue(config map[string]string) bool {
	if channel == nil {
		mqConnect()
	}
	err := channel.QueueBind(config["queue"],config["routingKey"],config["exchange"],false,nil)
	if err != nil {
		return false
	}
	return true
}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}

