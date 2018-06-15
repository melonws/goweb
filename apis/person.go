package apis

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"strconv"
	"log"
	."gorestful/models/person"
	"fmt"
	"gorestful/libs/logHelper"
)

/**
 * 这里类似于controller 通过之前的路由进入对应的函数
 * 然后获取参数,操作数据库，进行数据返回
 */


func IndexApi(c *gin.Context) {
	logHelper.WriteLog("[通知请求][请求数据][1,23,4,5]","notify/access")

	c.String(http.StatusOK,"It works")
}

func AddPersonApi(c *gin.Context) {

	//参数处理开始(后续应写全局的数据验证，避免每个函数都过滤)
	name := c.Request.FormValue("name")
	age := c.Request.FormValue("age")

	var errMsg = ""

	if name == "" {
		errMsg = "姓名不能为空"
	}

	if age == "" {
		errMsg = "年龄不能为空"
	}

	if errMsg != "" {

		c.JSON(http.StatusOK,gin.H{
			"stat" : 0,
			"msg" : errMsg,
			"data" : 0,
		})
		return
	}

	ageInt,err := strconv.Atoi(age)

	if err != nil {
		ageInt = 0
		log.Fatalln(err)
	}
	//参数处理完毕

	//这里的Person是从models里的person包引入的，
	//它就好比我们以前的model类，我们可以直接构造出来，就相当于new
	//然后调用它自有的方法，进行数据添加
	p := Person{ Name:name , Age:ageInt}

	//这个方法可以点进去看
	id,err := p.AddPerson()

	if err!= nil {
		log.Fatalln(err)
	}

	fmt.Println("insert person Id {}", id)
	msg := fmt.Sprint("数据插入成功")

	//最终把数据拼装格式 进行返回
	c.JSON(http.StatusOK,gin.H{
		"stat" : 1,
		"msg" : msg,
		"data" : id,
	})

}
