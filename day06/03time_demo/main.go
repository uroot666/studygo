package main

import (
	"fmt"
	"time"
)

func f1() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())

	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	// time.Unix() 时间戳解析出时间
	ret := time.Unix(1564803667, 0)
	fmt.Println(ret)

	// 时间间隔
	fmt.Println(time.Second)

	// now + 1小时
	fmt.Println(now.Add(24 * time.Hour))

	// 定时器
	// timer := time.Tick(time.Second * 1)
	// for t := range timer {
	// 	fmt.Println(t)
	// }

	// 格式化时间 把语言中时间对象转换成字符串类型的时间
	// 2022-10-29
	fmt.Println(now.Format("2006-01-02"))
	// 2022/10/29 23:47:59
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	// 按照对应格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02 15:04:05", "1992-10-05 18:30:00")
	if err != nil {
		fmt.Printf("parse time failed, err: %v\n", err)
		return
	}

	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())

	// Sleep
	fmt.Println(time.Now().Unix())
	time.Sleep(6 * time.Second)
	fmt.Println(time.Now().Unix())

}

func main() {
	f1()
}
