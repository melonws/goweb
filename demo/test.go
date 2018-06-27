package main

import (
	"fmt"
	"math/rand"
	"github.com/melonws/goweb/libs/logHelper"
)

func main()  {

	*logHelper.Name = "123"

	fmt.Println(logHelper.Name)

	*logHelper.Name = "456"

	fmt.Println(logHelper.Name)
	//mapTest()
	//iftest()
	//age := []int{18}
	//updateAgeInt(age)
	//fmt.Println(age)
	//updateAgeIntx(&age)
	//fmt.Println(age)
}

func updateAgeInt ( a []int){
	a[0] = 10
}

func updateAgeIntx (a *[]int) {
	(*a)[0] = 20
}

func maptest (){
	userData := map[string]string{
		"name" : "wangshu",
		"age" : "18",
	}

	updateName(&userData)
	fmt.Println(userData)
	updateAge(userData)
	fmt.Println(userData)
}

func updateName(user *map[string]string){

	if _,ok := (*user)["name"]; ok {
		(*user)["name"] = "王澍"
	}
}

func updateAge(user map[string]string){
	if _,ok := user["age"]; ok{
		user["age"] = "20"
	}
}


func iftest() {
	a := rand.Int()
	b := 3
	d := map[string]int{
		"a" : 2,
		"b" : 3,
	}

	if a > b {
		fmt.Println(a)
	}

	if v,ok := d["a"];ok{
		fmt.Println(v)
	}

}

func fortest()  {

	d := map[string]int{
		"a" : 2,
		"b" : 3,
	}

	for i:=0;i<10;i++ {
		fmt.Println(i)
	}

	for k,v := range d {
		fmt.Println(k,v)
	}

	for {
		fmt.Println(666)
	}

loop:
	count := 0
	for {
		count++
		if count == 250 {
			continue
		}
		fmt.Println(666)
		if count > 666 {
			break
		}
		if count == 665 {
			goto loop
		}
	}
}

var prop map[string]int = map[string]int{
	"score" : 80,
}

func Set(key string, value int){

	update(key,value)
}

func update(key string ,value int){
	if _,ok := prop[key]; ok {
		prop[key] = value
	}
}

func mapTest() {
	var dates map[string]interface{}

	dates = map[string]interface{}{
		"name":"wangshu",
		"age" : 18,
		"goods" : true,
		"buyRate" : 0.7,
	}

	fmt.Println(dates)
}