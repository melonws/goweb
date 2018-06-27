package main

import "fmt"

func main () {
	sliceAppend2()
}

func array () {

	var arr [5]int

	arr = [5]int{1,2,3}

	fmt.Println(arr)

	var arr2 = [...]int{4,5,6}

	fmt.Println(arr2)
}


func slice () {

	var slice = []int{}//append 速度慢

	var slice2 = make([]int,10) //append 速度快 但是有10个0

	var slice3 = make([]int,0,10) //append 速度快 没有10个0 第三个参数是容量

	fmt.Println(slice2)
	fmt.Println(slice3)

	slice = []int{1,2,3}

	fmt.Println("slice:",slice)

	slice = append(slice,4)
	//当我们用append追加元素到切片时，如果容量不够，go就会创建一个新的切片变量


	fmt.Println("append:",slice)

	index := 3

	slice = append(slice[:index],slice[index+1:]...)

	fmt.Println("remove element:",slice)

}

func slice2 () {
	var osa = make ([]string,0);

	sa:=&osa;

	for i:=0;i<10;i++{

		*sa=append(*sa,fmt.Sprintf("%v",i))

		fmt.Printf("addr of osa:%p,\taddr:%p \t content:%v\n",osa,sa,sa);

	}

	fmt.Printf("addr of osa:%p,\taddr:%p \t content:%v\n",osa,sa,sa);
}


/*
 * 动态申请内存
 */
func sliceAppend()  {

	var s []int

	fmt.Printf("addr of osa:%p,\t content:%v\n",s,s);

	s = append(s,1)

	fmt.Printf("addr of osa:%p,\t content:%v\n",s,s);

	s = append(s,2)

	fmt.Printf("addr of osa:%p,\t content:%v\n",s,s);
}

/*
 * 定义容量，
 */
func sliceAppend2()  {

	var s = make([]int,2,3)

	fmt.Printf("addr of osa:%p,\t content:%v,\t length:%d ,\t capacity:%d \n",s,s,len(s),cap(s));

	s = append(s,1)

	fmt.Printf("addr of osa:%p,\t content:%v,\t length:%d,\t capacity:%d\n",s,s,len(s),cap(s));

	s = append(s,2)

	fmt.Printf("addr of osa:%p,\t content:%v,\t length:%d,\t capacity:%d\n",s,s,len(s),cap(s));

	s = append(s,3)

	fmt.Printf("addr of osa:%p,\t content:%v,\t length:%d,\t capacity:%d\n",s,s,len(s),cap(s));

	s = append(s,4)
	s = append(s,5)
	fmt.Printf("addr of osa:%p,\t content:%v,\t length:%d,\t capacity:%d\n",s,s,len(s),cap(s));
}