package main

import (
	"fmt"
	"math/big"
	"reflect"
	"unicode/utf8"
)

func main() {

	// _ 为匿名引用或命名，go编译器不允许无效引用

	// all data type in Go
	// 总的来说有4类，1. 数字型 整型和浮点 2.布尔类型 3.字符型&字符串类型 4.数组类型

	// 整型&浮点
	// 整型 数字代表二进制数值所占的长度（默认为cpu字长：cpu一次性存取数据的最大长度）
	// 默认情况下为cpu位数（32、64），32位cpu一样可以定义64位字长的数据类型
	var _ int8
	var _ int16
	var _ int32
	var _ int64
	var _ int
	// 无符号（unsigned）
	var _ uint8
	var _ uint16
	var _ uint32
	var _ uint64
	var _ uint
	//浮点 Go中没有float关键字（考虑到float浮点类型不能精确表示数字，浮点的位数决定了精度）
	var _ float32
	var _ float64
	var _ complex64
	var _ complex128

	// 字符串类型，UTF-8编码
	var _ string
	_ = "string"
	str := "str is string"
	fmt.Println(str)

	// 字符类型
	// 20013
	var _str = '中'
	fmt.Println(_str)
	// Printf 格式化显示
	fmt.Printf("%c", _str)
	fmt.Println()
	fmt.Println(string(_str))

	// len函数，查看实际占用的存储空间，底层字节数组长度
	var _chars = "字符串"
	len := len(_chars)
	fmt.Println(len)

	for i := 0; i < len; i++ {
		// 16进制格式化
		fmt.Printf("%X", _chars[i])
		//字节数组元素
		//fmt.Println(_chars[i])
	}

	fmt.Println()
	fmt.Println("_chars字符串的长度为：", utf8.RuneCountInString(_chars))
	for _, z := range _chars {
		fmt.Printf("%T, %X\n", z, z)
	}

	// rune （int32）

	// byte （uint8）

	var _ bool

	_ = true
	_ = false

	tr := true
	fs := false
	fmt.Println("tr is", tr)
	fmt.Println("fs is", fs)
	//a = int, & int32?int64
	var a = 1
	var _a = "1"
	var b int32 = 2
	var c int64 = 3
	var d = 3.14
	var e float32 = 3
	fmt.Println(_a)
	fmt.Println(reflect.TypeOf(_a))
	// int
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	// int32
	fmt.Println(reflect.TypeOf(b))
	// int64
	fmt.Println(reflect.TypeOf(c))
	// float64 默认长度
	fmt.Println(reflect.TypeOf(d))
	// float32
	fmt.Println(reflect.TypeOf(e))

	// 浮点的比较，不能使用 “==”、“>”、“<”
	// 使用big包中的函数，-1小于 0等于 1大于
	f := 3.14
	g := 3.15
	result := big.NewFloat(f).Cmp(big.NewFloat(g))
	fmt.Println(result)

	// 数组，Go中数组大小固定，生命时必须指定大小
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println(reflect.TypeOf(arr))
	arr2 := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr2)
	fmt.Println(reflect.TypeOf(arr2))
	// Go中没有隐式类型转换，所以类型 [5]int 和 [7]int 不一样

	// 通过下标改变数组元素的值
	arr2[6] = 5
	fmt.Println(arr2)

	// 不像java有自上而下的类型转换和继承关系
	// 继承虽然便利，但是带来了复杂的类型层次，也导致了语言和系统设计的复杂性
	// 复杂的继承和类型绑定转换也会增加运行时性能开销
	// 所以让我们忘掉什么背景语言（尤其是狗屎java。），从go重新开始

	// to do 补充
	// 字符和字符串以及unicode、utf-8、ASCII码
	// 可变大小的数组：Slice 切片

}
