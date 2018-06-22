package main

import (
	"github.com/melonws/goweb/demo/pipelineDemo/pipeline"
	"fmt"
	"os"
	"bufio"
)

func main() {


}

func source () {
	p := pipeline.ArraySource(3,2,6,7,4,5,1,8,9)

	for {
		if num,ok := <-p; ok {
			fmt.Println(num)
		} else {
			break
		}
	}

	for v := range p {
		//如果使用 range 信道
		//那么 信道一定要 close 否则就无限循环
		fmt.Println(v)
	}
}

func mergeDemo () {
	p := pipeline.Merge(
		pipeline.InMemSort(pipeline.ArraySource(3,2,6,7,4)),
		pipeline.InMemSort(pipeline.ArraySource(5,1,9,8,10)),
	)

	for v := range p {
		//如果使用 range 信道
		//那么 信道一定要 close 否则就无限循环
		fmt.Println(v)
	}
}

func readAndWriter () {
	const filename = "large.in"
	const n  =100000000

	file,err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := pipeline.RandomSource(n)

	//read write 速度慢 bufio 包装加速
	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer,p)
	writer.Flush()

	file,err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReadSource(bufio.NewReader(file),-1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count ++
		if count > 100 {
			break
		}
	}
}

func expsort() {
	//p := createPipline()
	//wirteTofile(p)
	//printFile()
}


