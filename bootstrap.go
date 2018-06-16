package main

import (
	"github.com/melonws/goweb/cmds"
	"reflect"
	"errors"
	"flag"
	"fmt"
	"github.com/melonws/goweb/cmds/mqOrderCourse"
)

var funcMap map[string]interface{}

var funcDesc map[string]string

func init() {
	funcMap = map[string]interface{}{
		"custom_reg": cmd_custom_reg.Start,
		"order_course":mqOrderCourse.Start,
	}

	funcDesc = map[string]string{
		"custom_reg" : "用户注册",
		"order_course":"成课数据",
	}
}

func main() {

	var name string
	var count int

	flag.StringVar(&name,"name","default","脚本名称")
	flag.IntVar(&count,"count",1,"脚本启动数量")
	flag.Parse()
	if name == "default" {
		fmt.Println("请使用 -h 参数查看命令参数")

		fmt.Println("例如 -name 脚本名 -count 执行数量")

		fmt.Println("脚本列表:")
		for key,value := range funcDesc {
			fmt.Println(key+":"+value)
		}
		return
	}

	if _,ok := funcMap[name]; !ok{
		fmt.Println("不存在这个方法")

		for key,value := range funcDesc {
			fmt.Println(key+":"+value)
		}
		return
	}
	Call(funcMap,name,count)

}



func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}