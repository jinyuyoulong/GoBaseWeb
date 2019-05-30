# Beanstalk

Go client for [beanstalkd](https://beanstalkd.github.io).

## Install

    $ go get github.com/beanstalkd/go-beanstalk

## Use

Produce jobs:

    c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
    id, err := c.Put([]byte("hello"), 1, 0, 120*time.Second)

Consume jobs:

    c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
    id, body, err := c.Reserve(5 * time.Second)

---

`Beanstalkd`是一个简单、高效的工作队列系统，其最初设计目的是通过后台异步执行耗时任务方式降低高容量Web应用的页面延时。而其简单、轻量、易用等特点，和对任务优先级、延时 超时重发等控制，以及众多语言版本的客户端的良好支持，使其可以很好的在各种需要队列系统的场景中应用。

我们使用的第三方库来自： "github.com/beanstalkd/go-beanstalk"

下面是Beanstalkd官方的状态图：

```clike
  //------------------- 状态图来自官方文档 -------------------//

   put with delay               release with delay
  ----------------> [Delayed] <------------.
                        |                   |
                 kick   | (time passes)     |
                        |                   |
   put                  v     reserve       |       delete
  -----------------> [Ready] ---------> [Reserved] --------> *poof*
                       ^  ^                |  |
                       |   \  release      |  |
                       |    `-------------'   |
                       |                      |
                       | kick                 |
                       |                      |
                       |       bury           |
                    [Buried] <---------------'
                       |
                       |  delete
                        `--------> *poof*
```

示例
```
package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	beanstalk "myvendor/go-beanstalk" //第三方使用相对路径
)

const (
	// TubeName1 channel 通道 value 不能是非法格式: 不能是中文，不能有空格，等
	TubeName1 = "channel1"
	// TubeName2 channel2
	TubeName2 = "channel2"
	// ip docker 对外 ip
	ip = "192.168.99.100"
)

// Producer 生产者 发布任务
func Producer(fname, tubeName string) {
	if fname == "" || tubeName == "" {
		return
	}

	c, err := beanstalk.Dial("tcp", ip+":11300")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	c.Tube.Name = tubeName
	c.TubeSet.Name[tubeName] = true
	fmt.Println(fname, " [Producer] 等着，我要开始生产啦 tubeName:", tubeName, " c.Tube.Name:", c.Tube.Name)

	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("massage for %s %d", tubeName, i)
		// 生产数据
		c.Put([]byte(msg), 30, 0, 120*time.Second)
		fmt.Println(fname, " [Producer] beanstalk put body:", msg)
		//time.Sleep(1 * time.Second)
	}

	c.Close()
	fmt.Println("Producer() end.")
}

// Consumer 消费者
func Consumer(fname, tubeName string) {
	if fname == "" || tubeName == "" {
		return
	}

	c, err := beanstalk.Dial("tcp", ip+":11300")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	c.Tube.Name = tubeName
	c.TubeSet.Name[tubeName] = true

	fmt.Println(fname, " [Consumer] 注意 我要开始消费啦 ———— tubeName:", tubeName, " c.Tube.Name:", c.Tube.Name)

	substr := "timeout"
	for {
		fmt.Println(fname, " [Consumer]------------------ ")
		//从队列中取出
		id, body, err := c.Reserve(1 * time.Second)
		if err != nil {
			if !strings.Contains(err.Error(), substr) {
				fmt.Println(fname, " [Consumer] 慢着 有错 [", c.Tube.Name, "] err:", err, " id:", id)
			}
			continue
		}
		fmt.Println(fname, " [Consumer] 收到消费数据 [", c.Tube.Name, "] job:", id, " body:", string(body))

		//从队列中清掉
		err = c.Delete(id)
		if err != nil {
			fmt.Println(fname, " [Consumer] [", c.Tube.Name, "] Delete err:", err, " id:", id)
		} else {
			fmt.Println(fname, " [Consumer] 清除 成功[", c.Tube.Name, "] Successfully deleted. id:", id)
		}
		fmt.Println(fname, " [Consumer]------------------")
		//time.Sleep(1 * time.Second)
	}
	fmt.Println("Consumer() end. ")
}

func main() {
	// 根据本机CPU核数 设置 go max procs
	// runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(1)

	go Producer("PA", TubeName1)
	go Producer("PB", TubeName2)

	go Consumer("CA", TubeName1)
	go Consumer("CB", TubeName2)

	time.Sleep(10 * time.Second)
}

```