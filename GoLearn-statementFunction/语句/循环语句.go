/**
1. for 循环 Go语言的For循环有3中形式，只有其中的一种使用分号
for init; condition; post { }
for condition { }
for { }

    init： 一般为赋值表达式，给控制变量赋初值；
    condition： 关系表达式或逻辑表达式，循环控制条件；
    post： 一般为赋值表达式，给控制变量增量或减量。


2. 嵌套循环
 */
package main

import "fmt"

// 1. for循环
func forCycle() {
	var b int = 15
	// 默认值为0
	var a int

	// 标准
	for a := 0; a < 10; a++ {
		fmt.Println("a 的值为:", a)
	}

	// 关系表达式
	for a < b {
		a++
		fmt.Println("a 的值为:", a)
	}
	// 相当于 while(true)
	for {
		fmt.Println("this is a deadLoop")
	}
}

// 2. 嵌套循环
func nesting() {
	for a := 0; a < 10; a++ {
		for b := 0; b < 5; b++ {
			fmt.Println("b的值: ", b)
		}
		fmt.Println("a的值: ",a)
	}
}

func main() {
	//forCycle()
	nesting()
}
