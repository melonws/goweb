package mqHelperg

import (
	//"runtime"
	"sync/atomic"
	"github.com/streadway/amqp"
	//"os"
	//"os/signal"
	//"syscall"
	"time"
)

/**
 *  任务获取的对象
 */
type TaskData struct{
	amqp.Delivery
}

/**
 * 整体消费服务的对象
 */
type Service struct {
	capacity int
	tasks chan *TaskData
	numThread int
	closeChans chan struct{}
	stopFlag int32
	loopStopChan chan struct{}
	callback func(delivery amqp.Delivery)
	config map[string]string
}

/**
 * 创建一个消费进程对象
 * 需要传来mq队列配置,拿到消息后的回调函数,协程数量
 */
func NewService(config map[string]string , callback func(delivery amqp.Delivery),capacity int) *Service{
	service := &Service{}
	service.capacity = capacity
	service.numThread = capacity
	//service.numThread = runtime.NumCPU() * 2
	service.tasks = make(chan *TaskData,capacity)
	service.stopFlag = 0
	service.closeChans = make(chan struct{},service.numThread)
	service.loopStopChan = make(chan struct{})
	service.callback = callback
	service.config = config
	return service

}

/**
 * 停止所有的协程
 */
func (this *Service) Stop() {
	atomic.StoreInt32(&this.stopFlag,1)
	<-this.loopStopChan
	close(this.tasks)
	for i := 0; i < this.numThread; i++ {
		<-this.closeChans
	}
}


/**
 * 开始脚本进程的入口
 * 通过设置的协程数量，开始启动run方法,处理拿到消息的回调函数
 * 开启 LoopConsume方法，看名字就知道是循环消费
 */
func (this *Service) Run() {
	for i := 0; i < this.numThread; i++ {
		go this.run(i)
	}
	go this.LoopConsume()

	forever := make(chan bool)
	<-forever
}

/**
 * 从tasks 无缓冲信道中拿消费到的delivery(MQ文档)
 * 然后调用callback函数进行业务操作
 * 记得callback中要 Ack 否则会阻塞
 */
func (this *Service) run(i int) {
loop:
	for {
		select { case task, ok := <-this.tasks:
			if ok {
				//#TODO process
				this.callback(task.Delivery)
			} else {
				break loop
			} }
	}
	this.closeChans <- struct{}{}
}


/**
 * 循环从mq中消费数据底层通过Mq的包实现
 * 单协程唤起读取数据的函数，该函数最底层是循环获取数据。
 */
func (this *Service) LoopConsume() {

	deliverys := make(chan amqp.Delivery,this.capacity)

	//由于ReadMsg所在文件mqHelper 和 misson是一个包 所以直接调用
	go ReadMsg(this.config,deliverys,this.capacity)

	for atomic.LoadInt32(&this.stopFlag) == 0 {
		//TODO ReadData
		//将消费到的数据读出，并且插入tasks无缓冲信道
		d := <- deliverys
		//TODO ReadData
		task := &TaskData{d}
		this.tasks <- task
		//5ms的间隔建立协程，否则mq消费连接中断，至于为什么我也再探索
		time.Sleep(time.Millisecond*5)
	}
	this.loopStopChan <- struct{}{}
}



//func main(){
//	service := NewService(3 ,callback)
//	go service.Run()//启动程序处理
//
//	forever := make(chan bool)
//	//service.Stop()
//	<-forever
//	//c := make(chan os.Signal)
//	//signal.Notify(c,os.Interrupt,syscall.SIGINT,syscall.SIGTERM)
//	//s := <-c
//	//fmt.Println(s)
//	//service.Stop()
//	//fmt.Println("exit :D")
//}