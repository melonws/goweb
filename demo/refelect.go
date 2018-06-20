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
 */

type Style struct{
	f string `species:"gopher" color:"blue"`
	a string `species:"ggg" color:"red"`
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
	structToJson()

}

func structToJson(){
	data := Data{1,"WangShu","Some things",true,false,time.Now()}

	jsonByte,_ := json.Marshal(data)

	fmt.Println(string(jsonByte))
}


func reflectTest() {
	s := Style{}

	st := reflect.TypeOf(s)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))
}