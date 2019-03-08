/**
Go 语言中变量可以在三个地方声明：
    1. 函数内定义的变量称为局部变量
    2. 函数外定义的变量称为全局变量
	3. 局部变量与全局变量名称相同,会如何选择?
    3. 函数定义中的变量称为形式参数

 */
package main

import "fmt"

// 1. 局部变量
func localVariable() {
	// 声明局部变量
	var a int;
	// 初始化值
	a = 10
	fmt.Println("a的值: ", a)
}

// 声明全局变量
var b int;

// 2. 全局变量
func globalVariable() {
	// 全局变量初始化值
	b = 20
	fmt.Println("b的值: ", b)

}

// 3. 局部变量与全局变量名称相同情况下
// 打印的值为50 函数内的局部变量会被优先考虑
func localGloba() {
	var b = 50
	fmt.Println("b的值: ", b)
}

// 4. 形式参数
func formalParameter() {
	/* main 函数中声明局部变量 */
	var a int = 10
	var b int = 20
	var c int = 0

	fmt.Println("main()函数中 a =", a);
	c = sum(a, b);
	fmt.Println("main()函数中 c =", c);
}

/* 函数定义-两数相加 */
func sum(a, b int) int {
	fmt.Println("sum() 函数中 a =", a);
	fmt.Println("sum() 函数中 b =", b);
	return a + b;
}

func main() {
	localVariable()
	globalVariable()
	localGloba()
	formalParameter()
}
