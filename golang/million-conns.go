package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	MaxWorker = 100 // 随便设置值
	MaxQueue  = 200 // 随便设置值
)

// 一个可以发送工作请求的缓冲 channel
var JobQueue chan Job

func init() {
	JobQueue = make(chan Job, MaxQueue)
}

type Payload struct{}

type Job struct {
	PayLoad Payload
}

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

// Start 方法开启一个 worker 循环，监听退出 channel，可按需停止这个循环
func (w Worker) Start() {
	go func() {
		for {
			// 将当前的 worker 注册到 worker 队列中
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				//   真正业务的地方
				//  模拟操作耗时
				time.Sleep(500 * time.Millisecond)
				fmt.Printf("上传成功:%v\n", job)
			case <-w.quit:
				return
			}
		}
	}()
}

func (w Worker) stop() {
	go func() {
		w.quit <- true
	}()
}

// 初始化操作

type Dispatcher struct {
	// 注册到 dispatcher 的 worker channel 池
	WorkerPool chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool}
}

func (d *Dispatcher) Run() {
	// 开始运行 n 个 worker
	for i := 0; i < MaxWorker; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				// 尝试获取一个可用的 worker job channel，阻塞直到有可用的 worker
				jobChannel := <-d.WorkerPool
				// 分发任务到 worker job channel 中
				jobChannel <- job
			}(job)
		}
	}
}

// 接收请求，把任务筛入JobQueue。
func payloadHandler(w http.ResponseWriter, r *http.Request) {
	work := Job{PayLoad: Payload{}}
	JobQueue <- work
	_, _ = w.Write([]byte("操作成功"))
}

func main() {
	// 通过调度器创建worker，监听来自 JobQueue的任务
	d := NewDispatcher(MaxWorker)
	d.Run()
	http.HandleFunc("/payload", payloadHandler)
	log.Fatal(http.ListenAndServe(":8099", nil))
}

/**
最终采用的是两级 channel，一级是将用户请求数据放入到 chan Job 中，这个 channel job
相当于待处理的任务队列。

另一级用来存放可以处理任务的 work 缓存队列，类型为 chan chan Job。调度器把待处理的
任务放入一个空闲的缓存队列当中，work 会一直处理它的缓存队列。通过这种方式，实现了
一个 worker 池。

首先我们在接收到一个请求后，创建 Job 任务，把它放入到任务队列中等待 work 池处理。

调度器初始化work池后，在 dispatch 中，一旦我们接收到 JobQueue 的任务，就去尝试获取
一个可用的 worker，分发任务给 worker 的 job channel 中。 注意这个过程不是同步的，
而是每接收到一个 job，就开启一个 G 去处理。这样可以保证 JobQueue 不需要进行阻塞，
对应的往 JobQueue 理论上也不需要阻塞地写入任务。

这里"不可控"的 G 和上面还是又所不同的。仅仅极短时间内处于阻塞读 Chan 状态， 当有
空闲的 worker 被唤醒，然后分发任务，整个生命周期远远短于上面的操作。
*/
