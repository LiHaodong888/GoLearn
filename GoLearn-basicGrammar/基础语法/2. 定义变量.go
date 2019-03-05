/**
Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。
声明变量的一般形式是使用 var 关键字
var identifier type 变量名在前 类型在后 这个有所不同

定义变量分为四种
1. 指定变量类型
2. 根据值自行判定变量类型
3. 最简体
4. 聚合声明
 */
package main

import "fmt"

// 定义具体变量 使用默认初值
func variableDefaultValue() {
	var a int    // 0
	var b string // ""
	var c bool   // false
	fmt.Println(a, b, c)
}

// 定义变量，赋初值
func variableInitialValue() {
	var a, b int = 1, 2
	var s string = "www.lhdyx.cn"
	fmt.Println(a, b, s)
}

// 类型推断
func variableTypeDeduction() {
	var a, b, c, d = 1, 2, true, "www.lhdyx.cn"
	fmt.Println(a, b, c, d)
}

// 最简定义变量方式
func variableShorter() {
	a, b, c, d := 3, 2, true, "hi"
	fmt.Println(a, b, c, d)
}

// 聚合调用
func variablePolymerization() {
	var a, b, c string
	a, b, c = "1","2","3"
	fmt.Println(a, b, c)
}
// 依次调用
func main() {
	variableDefaultValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	variablePolymerization()
}
