package orderCourse

import (
	"gorestful/libs/dbHelper"
	"gorestful/libs/logHelper"
	"time"
	"gorestful/libs/esHelper"
	"fmt"
)

type orderCourse struct {

}

var orderInfoFilter = []string{
	"start_day","course_begin_time","course_end_time","join_time","leave_time","created_at","updated_at","refound_time","end_time",
}


func AddToEs(courseOrderInfoId int) bool {

	connect,model,err := dbHelper.CreateModel("course_order_infos")
	if err != nil {
		logHelper.WriteLog("数据库连接错误"+err.Error(),"orderCourseInfo/addtoes/error")
		return false
	}
	defer connect.Close()

	orderinfo,err := model.Where("id","in",[]interface{}{courseOrderInfoId}).First()
	fmt.Println(orderinfo)
	if err != nil {
		logHelper.WriteLog("数据库查询出错、或订单不存在"+err.Error(),"orderCourseInfo/addtoes/error")
		return false
	}

	filter(&orderinfo,orderInfoFilter)

	prefix := "xes_od_"

	var newData  map[string]interface{}
	newData = make(map[string]interface{})
	for k,v := range orderinfo {
		key := prefix+k
		newData[key] = v
	}


	Es := esHelper.CreateEs()
	//x := a.Add("wangshu","1","",data)
	result := Es.Edit("wangshu","1","",newData)
	return result
}

func filter(data *map[string]interface{},filter []string){
	ds := *data
	for _,k := range filter {
		if d,ok := ds[k]; ok {
			if d == "0000-00-00 00:00:00" {
				ds[k] = "2001-01-01 00:00:00"
				continue
			}
			timeStr := ds[k].(string)
			lenth:=len([]rune(timeStr))

			if lenth == 10 {
				timeStr += " 00:00:00"
			}else if lenth == 19{

			}else {
				timeStr = "2001-01-01 00:00:00"
			}
			unix,_ := time.Parse("2006-01-02 15:04:05","2018-06-14 17:50:55")
			formatTime := unix.Format("2006-01-02 15:04:05")
			ds[k] = formatTime
		}
	}
}