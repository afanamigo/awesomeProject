package main

import (
	"fmt"
	"os"
	"time"
)

var a = 17
var f = 3.14

func main() {

	// 内置函数
	print("what the hack")
	fmt.Println()

	// 还没研究到os包
	writeString, err := os.Stdout.WriteString("Hello World from OS")
	if err != nil {
		print(writeString)
	}

	fmt.Println("Time is ", time.Now())

	// 基数初级类型定义
	// 包级变量： var声明
	// var variableName = InitExpression (自动推导表达式类型、以及默认类型)
	fmt.Println("var a is ", a)
	fmt.Println("var f is ", f)

	// 局部变量
	var a1 int32 = 17
	var f1 float32 = 3.14
	fmt.Println()
	fmt.Println("var a1 is ", a1)
	fmt.Println("var f1 is ", f1)

	var a2 = int32(17)
	var f2 = float32(3.14)
	fmt.Println()
	fmt.Println("var a2 is ", a2)
	fmt.Println("var f2 is ", f2)

	// var 声明代码块
	var (
		a3 = 17
		f3 = float32(3.14)
	)
	fmt.Println()
	fmt.Println("var a3 is ", a3)
	fmt.Println("var f3 is ", f3)

	// but not 看起来形式不一致
	var (
		a4         = 17
		f4 float32 = 3.14
	)
	fmt.Println()
	fmt.Println("var a4 is ", a4)
	fmt.Println("var f4 is ", f4)

	// 延迟初始化，变量拥有“0值”
	var a5 int32
	var f5 float32
	fmt.Println()
	fmt.Println("var a5 is ", a5)
	fmt.Println("var f5 is ", f5)

	//  命名习惯：声明聚类&就近原则
	// 同一类的变量声明放在 同一var块
}
