package main

import (
	"time"
	"fmt"
)

func main() {
	//orderCourse.AddToEs(18)
	//a := 2
	//b := 3
	//sumNum := make(chan int)
	//productNum := make(chan int)
	////empty
	//
	//go sum(a,b,sumNum)
	//go product(a,b,productNum)
	//fmt.Println(<-sumNum,<-productNum)

	channel := make(chan int,100)

	//go func() {
	//	for {
	//
	//
	//	}
	//}()//阻塞



	time.Sleep(10 * time.Second)
	//sum,ok := sumA(a,b)
	//
	//if ok {
	//	fmt.Println(sum)
	//}
	//
	//x,_ := productA(a,b)
	//fmt.Println(x)

}

func sum (a int , b int , ch chan int) {
	time.Sleep( 5 * time.Second)
	ch <- a+b
}

func product(a int ,b int ,ch chan int) {
	time.Sleep( 5 * time.Second)
	ch <- a*b
}

func sumA (a int , b int ) (int ,bool) {

	time.Sleep( 5 * time.Second)
	return a + b, true
}

func productA(a int ,b int ) (int ,bool) {
	time.Sleep( 5 * time.Second)
	return a*b,true
}