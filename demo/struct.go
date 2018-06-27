package main

import (
	"fmt"
	"github.com/melonws/goweb/libs/logHelper"
	"goweb/libs/esHelper"
)

type Humen struct {
	Name string
}

type Jump interface {
	jump()
} 


type Stundent struct {
	Humen
	Id int
	Grade string
	Class string
	Jump
}

type Cs struct {

}
type Chiji struct{}

func (cs *Cs) jump() {
	fmt.Println(444)
}


func (c *Chiji) jump() {
	fmt.Println(123)
}

func main() {

	h := Humen{"wangshu"}
	//c := &Chiji{}


	cs := &Cs{}

	s := Stundent{h,3,"一年级","二班",cs}


	s.say("同学们起立")
	s.hello()
}




func (this *Humen) say(str string){
	fmt.Println(this.Name+"说:"+str)
}


func (s *Stundent) hello(){
	s.Humen.say("我叫啥")//子类调用父类方法。
	fmt.Println(s.Name+"说:老师好")
}