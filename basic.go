package main

import (
	"fmt"
	"unsafe"
)

//常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数
const (
	c1 = "abc"
	c2 = len(c1)
	c3 = unsafe.Sizeof(c2)
)

func main() {

	// 声明一个变量并初始化
	var a = "Hello"
	fmt.Println(a)

	// 没有初始化就为零值
	var b int
	fmt.Println(b)

	// bool 零值为 false
	var c bool
	fmt.Println(c)

	var d = 100
	fmt.Println(d)

	f := 1000
	fmt.Println(f)

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)

}
