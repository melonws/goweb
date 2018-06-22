package main

import (
	"strings"
	"fmt"
	"time"
)

func main() {

	r := &ReadFromFile{
		path : "/tmp/access.log",
	}

	w := &WriteToDb{
		"username&password..",
	}

	lp := &LogProcess{
		reader : r,
		rc : make(chan string),
		wc : make(chan string),
		writer : w,
	}
	go lp.reader.Read(lp.rc)
	go lp.Process()
	go lp.writer.Write(lp.wc)

	time.Sleep(1 * time.Second)
}



type LogProcess struct {
	rc chan string
	wc chan string
	reader Reader
	writer Writer
}

type Reader interface {
	Read(rc chan string)
}

type Writer interface {
	Write(wc chan string)
}

type ReadFromFile struct {
	path string
}

type WriteToDb struct {
	influxDBDsn string
}

func (this *ReadFromFile) Read(rc chan string){
	//读取模块
	line := "message"
	rc <- line
}

func (this *WriteToDb) Write(wc chan string){
	//写入模块
	fmt.Println(<-wc)
}

func (l *LogProcess) Process()  {
	//解析模块
	data := <-l.rc
	l.wc <- strings.ToUpper(data)
}

