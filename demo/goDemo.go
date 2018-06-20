package main

import (
	"time"
	"fmt"
)

func main() {
	linear()
	parallel()
}

//-------------------------------------------
//类似php 等待一个函数返回再运行下一个函数
func linear() {
	a := 2
	b := 3
	sumNum,ok := sumA(a,b)
	if ok {
		fmt.Println(sumNum)
	}
	productNum,_ := productA(a,b)
	fmt.Println(productNum)
}

func sumA (a int , b int ) (int ,bool) {

	time.Sleep( 5 * time.Second)
	return a + b, true
}

func productA(a int ,b int ) (int ,bool) {
	time.Sleep( 5 * time.Second)
	return a*b,true
}

//类似php 等待一个函数返回再运行下一个函数
//-------------------------------------------


//-------------------------------------------
// 通过协程 并行跑两个函数
func parallel(){
	a := 2
	b := 3
	sumNum := make(chan int)
	productNum := make(chan int)
	//empty

	go sum(a,b,sumNum)
	go product(a,b,productNum)
	fmt.Println(<-sumNum,<-productNum)
}


func sum (a int , b int , ch chan int) {
	time.Sleep( 5 * time.Second)
	ch <- a+b
}

func product(a int ,b int ,ch chan int) {
	time.Sleep( 5 * time.Second)
	ch <- a*b
}
//通过协程 并行跑两个函数
//-------------------------------------------