package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	value int64
}

type result struct {
	*job
	result int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

func t1(ch1 chan<- *job) {
	wg.Done()
	// 循环生成int64类型随机数，发送到jobChan
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		ch1 <- newJob
		time.Sleep(time.Microsecond * 500)
	}
}

func t2(ch1 <-chan *job, resultChan chan<- *result) {
	wg.Done()
	// 从jobChan中取出随机数计算各位数的和，将结果发送给resultChan
	for {
		job := <-ch1
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job:    job,
			result: sum,
		}
		resultChan <- newResult
	}
}

func main() {
	wg.Add(1)
	go t1(jobChan)
	// 开启24个gouroutine执行job
	for i := 0; i < 24; i++ {
		go t2(jobChan, resultChan)
		wg.Add(1)
	}

	// 主goroutine从resultChan取出结果并打印到终端输出
	for result := range resultChan {
		fmt.Printf("value: %d sum: %d\n", result.job.value, result.result)
	}
	wg.Wait()
}
