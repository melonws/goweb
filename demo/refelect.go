package main



import	(
	"reflect"
	"fmt"
	"time"
	"encoding/json"
)

/**
 * 例子：https://www.cnblogs.com/WayneZeng/p/7606126.html
 * 		https://www.jianshu.com/p/c4ec92afeca8
 * 		https://www.cnblogs.com/skymyyang/p/7690837.html
 */

type Style struct{
	F string `species:"gopher" color:"blue"`
	A string `species:"ggg" color:"red"`
}


type Data struct {
	ID 			int 		`json:"id"`
	Name 		string 		`json:"name"`
	Bio			string 		`json:"about,omitempty"`
	Active  	bool 		`json:"active"`
	Admin 		bool 		`json:"-"`
	CreateAt	time.Time   `json:"create_at"`
}


func main() {
	reflectTest()

}

func structToJson(){
	data := Data{1,"WangShu","Some things",true,false,time.Now()}

	jsonByte,_ := json.Marshal(data)

	fmt.Println(string(jsonByte))
}


func reflectTest() {
	s := Style{"wangshu","shiwo"}

	st := reflect.TypeOf(s)
	fmt.Println(st.Field(0).Tag.Get("color"))

	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))
}