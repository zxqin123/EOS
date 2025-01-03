package main

import "fmt"

// 定义一个接口
type Worker interface {
	// 定义接口方法
	work()
	Method2(arg string) string
	GetValue() int
}

// 定义一个结构体
type Worker struct {
	ID      int
	Address string
}

// 为结构体实现接口方法
func (ms *MyStruct) Method1() {
	fmt.Println("Method1 called")
}

func (ms *MyStruct) Method2(arg string) string {
	return "Hello " + arg
}

func (ms *MyStruct) GetValue() int {
	return ms.value
}

func main() {
	// 创建一个结构体实例
	ms := &MyStruct{value: 42}

	// 使用接口
	var iface MyInterface = ms

	// 调用接口的方法
	iface.Method1()
	fmt.Println(iface.Method2("Go"))
	fmt.Println("Value:", iface.GetValue())
}
