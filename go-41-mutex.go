package main

import (
	"sync"
	"math/rand"
	"sync/atomic"
	"runtime"
	"time"
	"fmt"
)

func main() {
	var state = make(map[int]int)
	// 这个`mutex`将同步对状态的访问
	var mutex = &sync.Mutex{}
	var ops int64 = 0
	// 这里我们启动100个协程来不断地读取这个状态
	for r := 0; r < 100; r++ {

		for {
			total := 0
			// 对于每次读取，我们选取一个key来访问，
			// mutex的`Lock`函数用来保证对状态的
			// 唯一性访问，访问结束后，使用`Unlock`
			// 来解锁，然后增加ops计数器

			key := rand.Intn(5)
			mutex.Lock()
			total += state[key]
			mutex.Unlock()
			atomic.AddInt64(&ops, 1)

			// 为了保证这个协程不会让调度器出于饥饿状态，
			// 我们显式地使用`runtime.Gosched`释放了资源
			// 控制权，这种控制权会在通道操作结束或者
			// time.Sleep结束后自动释放。但是这里我们需要
			// 手动地释放资源控制权
			runtime.Gosched()
		}
	}

	for w:=0; w<10;w++{
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
