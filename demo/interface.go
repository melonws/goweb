package main

import (
	"fmt"
	"reflect"
)

type KTV struct{
	Address string
	Humens
}

type Humens interface{
	Say()
}

type Limin struct{
	Name string
	Height int
}

func (l *Limin) Say(){
	fmt.Println("我今晚没有名字")
}


func main() {

	var inter interface{} = "abc"

	var inter2 interface{} = 123

	fmt.Println(inter)
	fmt.Println(inter2)

	//断言
	fmt.Println("值是:"+inter.(string))

	//
	b,ok:=inter.(string)

	if ok {
		fmt.Println(b)
	}

	x := reflect.TypeOf(inter)

	v := reflect.ValueOf(inter)

	fmt.Println(x,v)



	k := KTV{"北京",&Limin{"李敏",190}}

	fmt.Println(k)
}

