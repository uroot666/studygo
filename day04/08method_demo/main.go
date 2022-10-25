package main

import "fmt"

// 方法
// go语言中如果标识符首字母大写，就表示对外部可见

type dog struct {
	name string
	age  int
}

// 构造函数
func newDog(name string, age int) dog {
	return dog{
		name: name,
		age:  age,
	}
}

// 方法是作用于特定类型的函数
// 接受者表示的是调用该方法的具体类型变量，多用类型名手字母小写表示
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪～\n", d.name)
}

func (d *dog) add() {
	d.age++
}

func (p *dog) dream() {
	fmt.Println("abc")
}

func main() {
	d1 := newDog("a", 1)
	d1.wang()
	fmt.Println(d1.age)
	d1.add()
	fmt.Println(d1.age)
	d1.dream()
}
