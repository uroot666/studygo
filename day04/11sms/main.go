package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	写一个系统能够查看、新增、删除学生
*/

type student struct {
	id   int64
	name string
}

var (
	allStudent map[int64]*student //变量声明
)

// newStudent 是student类型的构造函数
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudent() {
	// 打印所有的学生
	for k, v := range allStudent {
		fmt.Printf("学号: %d 姓名: %s \n", k, v.name)
	}
}

func addStudent() {
	// 向allStudent添加一个新的学生
	// 1. 创建一个新学生
	var (
		id   int64
		name string
	)
	fmt.Print("输入学生学号:")
	fmt.Scanln(&id)
	fmt.Print("输入学生姓名:")
	fmt.Scanln(&name)
	// 2. 追加到allStudent map中
	newStu := newStudent(id, name)
	allStudent[id] = newStu
}

func deleteStudent() {
	// 输入删除学生的序号
	var id int64
	fmt.Print("输入删除学生的学号:")
	fmt.Scanln(&id)
	// 2. allStudent这个map根据学号删除
	delete(allStudent, id)
}

func main() {
	allStudent = make(map[int64]*student, 50) // 初始化
	// 1. 打印菜单
	fmt.Println("欢迎使用学生管理系统！")
	fmt.Println(`
		1. 查看所有学生
		2. 新增学生
		3. 删除学生
	`)
	for {
		fmt.Printf("请输入操作: ")
		// 2. 等待用户选择要做什么
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了 %d 这个选项!\n", choice)
		// 3. 执行对应的函数
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("输入错误")
		}
	}
}
