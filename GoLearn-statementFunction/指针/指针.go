/**
我们都知道，变量是一种使用方便的占位符，用于引用计算机内存地址。
Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
1. 什么是指针
2. 如何去访问指针
3. 如何使用指针
4. Go 空指针
5. 判断是否为空指针
 */
package main

import "fmt"

/**
1.什么是指针
一个指针变量指向了一个值的内存地址。
类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：
var var_name *var-type
var ip *int        指向整型
var fp *float32    指向浮点型
 */


// 2. 访问变量在内存中地址 Go 语言的取地址符是 &
// 打印 0xc42006c188
func accessPointer()  {
	var a int = 20
	fmt.Println("变量的地址:", &a)
}
// 3. 如何使用指针
func usePointer() {
	var a int= 20   /* 声明实际变量 */
	var ip *int        /* 声明指针变量 */

	ip = &a  /* 指针变量的存储地址 */

	fmt.Println("a 变量的地址是:", &a  )

	/* 指针变量的存储地址 */
	fmt.Println("ip 变量储存的指针地址:", ip )

	/* 使用指针访问值 */
	fmt.Println("*ip 变量的值:", *ip )
}

/**
4. 空指针
当一个指针被定义后没有分配到任何变量时，它的值为 nil。
nil 指针也称为空指针。
nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
一个指针变量通常缩写为 ptr。
 */
func nullPointer() {
	var  ptr *int
	fmt.Println("ptr 的值为:", ptr )
}

// 5. 判断空指针
func judge()  {
	var a int  = 10
	var  ptr *int
	if (ptr != nil) {
		fmt.Println("不是空指针")
	} else if (ptr == nil) {
		fmt.Println("是空指针")
	}

	// 指针变量的存储地址
	ptr = &a

	if (ptr != nil) {
		fmt.Println("不是空指针")
	}

}
func main() {
	accessPointer()
	usePointer()
	nullPointer()
	judge()
}