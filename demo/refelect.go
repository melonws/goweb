package main

import	(
	"reflect"
	"fmt"
)

type Style struct{
	f string `species:"gopher" color:"blue"`
	a string `species:"ggg" color:"red"`
}


func main() {

   s := Style{}

   st := reflect.TypeOf(s)
   field := st.Field(0)
   fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))
}
