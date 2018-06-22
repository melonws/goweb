package pipeline

import (
	"sort"
	"io"
	"encoding/binary"
	"math/rand"
)

//可变长参数
func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	//channel 不能自己给自己写进去 要用协程写

	go func() {
		for _,v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}
				//只写 			//使用者只读
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		a := []int{}
		for v := range in {
			a = append(a,v)
		}
		//Sort
		sort.Ints(a)

		for _,v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func Merge(in1,in2 <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		v1,ok1 := <-in1
		v2,ok2 := <-in2
		//ok用来检测 channel 获取成功与否
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1,ok1 = <-in1
			} else {
				out <- v2
				v2,ok2 = <-in2
			}
		}
		close(out)
	}()

	return out
}

func ReadSource(reader io.Reader,chunkSize int) <-chan int {
	out := make(chan int)

	go func() {
		buffer := make([]byte,8)
		bytesRead := 0
		for {
			n,err := reader.Read(buffer)
			bytesRead += n
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil || (chunkSize != -1 && bytesRead > chunkSize) {
				break
			}
		}
		close(out)
	}()

	return out
}

func WriterSink ( writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte,8)
		binary.BigEndian.PutUint64(buffer,uint64(v))
		writer.Write(buffer)
	}
}


func RandomSource (count int) <-chan int {
	out := make(chan int)

	go func() {
		for i:=0;i<count;i++ {
			out <- rand.Int()
		}
		close(out)
	}()

	return out
}

//两两递归 归并排序
func MergeN (inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}
	m := len(inputs) / 2
	//inputs[0..m) and inputs[m..end)
	return Merge(MergeN(inputs[:m]...),MergeN(inputs[m:]...))
}