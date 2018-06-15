package person

import (
	"goweb/libs/dbHelper"
	"log"
)

//这里的type Person 就好比 class person
//里面的 id name age  就好比属性
type Person struct {
	Id int `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Age int `json:"age" form:"age"`
}

//而这个方法 就好比 class 中的非静态方法，方法名首字母大写相当于public
//这个方法一定要 构造出person 然后 person.AddPerson() 而不能直接靠包名调用
func (p *Person) AddPerson() (id int , err error) {

	//相比AddPerson不同，CreateModel就好比静态方法，可以直接用包名.方法名调用。
	connection,personModel,err :=dbHelper.CreateModel("test")

	if err!= nil {
		log.Fatalln(err)
		return 0,err
	}

	defer connection.Close()

	data := map[string]interface{}{
		"name":p.Name , "age":p.Age,
	}

	_,err = personModel.Data(data).Insert()

	return personModel.LastInsertId,err
}