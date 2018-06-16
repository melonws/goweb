package main

import "fmt"

type Humen struct {
	Name string
}

type Stundent struct {
	Humen
	Id int
	Grade string
	Class string
}

func main() {

	h := Humen{"wangshu"}

	s := Stundent{h,3,"一年级","二班"}



	s.say("同学们起立")
	s.hello()
}

func (h *Humen) say(str string){
	fmt.Println(h.Name+"说:"+str)
}

func (s *Stundent) hello(){
	fmt.Println(s.Name+"说:老师好")
}