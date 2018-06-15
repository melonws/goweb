package main

import (
	"fmt"
	."goweb/router"
)

/**
 * 用来测试mangodb插入数据用的对象
 */
type user struct{
	Id int
	Name string
	Description string
}

func test( a *map[string]interface{}) {
	x := *a
	for k,v := range x {
		(*a)[k] = v.(int) + 1
	}
}

/**
 * 主函数，整个程序启动入口
 */
func main() {

	//a := map[string]interface{}{
	//	"a":1,
	//	"b":2,
	//	}
	//
	//test(&a)
	//fmt.Println(a)
	//orderCourse.AddToEs(19)
	//data := user{1,"sumail","13"}
	//a := esHelper.CreateEs()
	////x := a.Add("wangshu","1","",data)
	//x := a.Edit("wangshu","1","",data)
	//fmt.Println(x)
	//config := Config.RedisConfig
	//
	//fmt.Println(config)

	//result := redisHelper.Set("test","what",1)
	//
	//data := redisHelper.Get("test")
	//fmt.Println(result)
	//fmt.Println(data)

	/**
	 * 初始化路由,是由于引入了 router包才有的方法，在包之前有个 .
	 * 这样直接把函数加载进来了，类似于require include
	 */
	router := InitRouter()

	router.Run(":8000")



	//session,collection :=mongoHelper.CreateModel("user")
	//
	//if session == nil {
	//	fmt.Println("链接失败")
	//	return
	//}
	//
	//defer session.Close()
	//
	//doc := user{
	//	bson.NewObjectId(),
	//	"test",
	//	"wangshu",
	//}
	//
	//err := collection.Insert(doc)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
}

func Count(ch chan int) {

	fmt.Println("Counting")
	ch <- 1
}