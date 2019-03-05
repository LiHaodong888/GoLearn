/**
常量是一个简单值的标识符，在程序运行时，不会被修改的量。
常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
常量的定义格式：const identifier [type] = value
常量分为两种
1. 指定类型定义： const b string = "abc"
2. 推断类型定义： const b = "abc"

 */
package main

import "fmt"

// 常量
func consts() {
	// 单个常量
	const A string = "1"

	// 多个常量
	const B, C, D string = "2", "3", "4"
	// 推断定义常量
	const E = 1
	fmt.Println(A, B, C, D, E)
}

// 枚举
func enums()  {
	// 使用常量枚举
	const (
		Unknown = 0
		Female = 1
		Male = 2
	)
	fmt.Println(Unknown, Female, Male) // 0 1 2

	// 使用 iota 块来实现自增枚举
	// iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，
	// const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)
	const (
		a = iota
		b
		c
	)
	fmt.Println(a,b,c)
}

// 依次调用
func main() {
	consts()
	enums()
}
