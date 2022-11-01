package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

// ini配置文件解析器

// MysqlConfig MySQL配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `int:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

// Config ...
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 0. 参数的校验
	// 0.1 传进来的data参数必须是指针类型，以为需要在函数中对其赋值
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer") // 格式化输出之后返回一个error类型
		return
	}
	// 0.2 传健来的data参数必须是结构体类型指针，以为配置文件中各种键值对需要赋值给结构体的字段
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct point")
		return
	}
	// 1. 读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	// string(b) 将字节类型的内容转换为字符串
	lineSlice := strings.Split(string(b), "\n")
	fmt.Println(lineSlice)
	// 2. 一行一行的读数据
	var structName string
	for idx, line := range lineSlice {
		line = strings.TrimSpace(line)
		// 如果是空行就跳过
		if len(line) == 0 {
			continue
		}
		// 2.1 如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 2.2 如果是[开头的表示节
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line: %d syntax error", idx+1)
				return
			}
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line: %d syntax error", idx+1)
				return
			}
			// 根据字符串sectionName去data里面根据反射找队友的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 说明找到了对应的嵌套结构体，把字段名记录下来
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			// 2.3 如果不是[开头就是=分割的键值对
			// 1. 以等号分割这一行，左边key、右边value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line: %d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			// value := strings.TrimSpace(line[index:])
			// 2. 根据structName 去data里面把对应的嵌套结构体取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) // 拿到嵌套结构体的值信息
			sType := sValue.Type()                     //拿到嵌套结构体的类型信息
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			// 3. 遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				filed := sType.Field(i) // tag信息是存储在类型信息中的
				if filed.Tag.Get("ini") == key {
					// 找到对应的字段
					fieldName = filed.Name
				}
			}
			//  4. 如果key = tag,给这个字段赋值
			// 4.1 根据fieldName 去取出这个字段
			fileObj := sValue.FieldByName(fieldName)

			// 4.2 对其赋值
			fmt.Println(fileObj.Type().Kind())
			return
		}

	}
	return nil
}

func main() {
	var cfg Config
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed, err: %v\n", err)
		return
	}
	fmt.Println(cfg)
}
