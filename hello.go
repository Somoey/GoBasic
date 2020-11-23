package main

// Go 程序是通过 package 来组织的。
// 只有 package 名称为 main 的源码文件可以包含 main 函数。
// 一个可执行程序有且仅有一个 main 包。
// 通过 import 关键字来导入其他非 main 包。
// Go 语言的包引入一般为: 项目名/包名
// import "test/controllers"

import "fmt"

// 通过 const 关键字来进行常量的定义。
// 通过在函数体外部使用 var 关键字来进行全局变量的声明和赋值。
// 通过 type 关键字来进行结构(struct)和接口(interface)的声明。
// 通过 func 关键字来进行函数的声明。
const age = 20

var name = "Somoey"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("My name is " + name)
}

//首字母小写--private；大写--public
// 方法的调用为: 包名.方法名()
// controllers.Test()
// 本包内方法名可为小写，包外调用方法名首字母必须为大写。
