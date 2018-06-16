package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	block2()

}

func deadlock() {
	//无缓冲信道
	channel := make(chan bool)

	//写数据
	channel<- true

	//读数据
	<-channel

	//死锁
	//信道是主进程和协程、或者协程之间通信用的
}


func dontdeadlock() {
	//无缓冲信道
	channel := make(chan bool)

	//协程写数据
	go func() {
		fmt.Println("这回不死了")
		channel<- true
	}()

	//读数据
	<-channel

}

//写的慢 读阻塞
func block1 () {

	channel := make(chan int)

	//协程写数据
	go func() {
		i := 0
		for{
			i++
			time.Sleep(5 * time.Second)
			channel<- i
			fmt.Println("协程写入信道成功")

		}
	}()

	var num int
	//读数据
	for{
		num = <-channel
		fmt.Println("主进程读出信道成功"+strconv.Itoa(num))
	}

}

//读的慢 写阻塞
func block2 () {

	channel := make(chan int)

	//协程写数据
	go func() {
		i := 0
		for{
			i++
			channel<- i
			fmt.Println("协程写入信道成功")
			if i == 50 {
				fmt.Println("写入完毕")
				break
			}
		}
	}()

	var num int
	//读数据
	for{
		num = <-channel
		time.Sleep(5 * time.Second)
		fmt.Println("主进程读出信道成功"+strconv.Itoa(num))
	}

}