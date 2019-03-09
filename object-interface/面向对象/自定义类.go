/**
Go语言的结构体相当于java的对象
Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
结构体是由一系列具有相同类型或不同类型的数据构成的数据集合

1. 定义类(结构体)
2. 实例化类
3. 访问类的属性
 */

package main

import "fmt"

/**
	1. 定义类
	结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体有中有一个或多个成员。
	type 语句设定了结构体的名称。结构体的格式如下
	type 类名 struct {
   member definition;
   member definition;
   ...
   member definition;
	}
	 */
// 定义User类
type user struct {
	name string
	age  int
}

// 自定义 admin 类
type admin struct {
	// 自定义类
	person user
	// 内置类型
	level string
}

func main() {

	/**
	2. 实例化类
	一旦定义了结构体类型，它就能用于变量的声明，语法格式如下
	variable_name := structure_variable_type {value1, value2...valuen}
		或
	variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
	 */
	// 创建 user 变量，所有属性初始化为其零值
	var bill user
	fmt.Println(bill) // {  "" 0}
	// 创建 user 变量，并初始化属性值
	xm := user{
		name: "小明",
		age:  12}
	fmt.Println(xm) // {小明 12}

	// 直接使用属性值，属性值的顺序要与 struct 中定义的一致
	xm2 := user{"小明", 12}
	fmt.Println(xm2) // {小明 12}

	// 含有自定义类型的 struct 进行初始化
	fred := admin{xm,"admin"}

	fmt.Println(fred) // {{小明 12} admin}

	name := xm.name
	fmt.Println("小明的姓名: ", name) // 小明的姓名:  小明

}
