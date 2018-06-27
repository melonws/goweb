package main

import (
	"time"
	"strconv"
	"fmt"
)

func main() {
	cushion()
}

//写的快 读的慢 但是有缓冲
func cushion(){

	channel := make(chan int ,50)

	//协程写数据
	go func() {
		i := 0
		for{
			i++
			channel<- i
			fmt.Println("协程写入信道成功")
			if i == 52 {
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
		fmt.Println("主进程读出信道成 功"+strconv.Itoa(num))
	}

}
