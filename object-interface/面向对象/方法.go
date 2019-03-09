/**
方法
如果要访问结构体成员，需要使用点号 . 操作符，格式为：
结构体.成员名

    方法的定义方法实际上也是函数，只是在声明时，在关键字 func 和方法名之间增加了一个参数
        普通的函数定义 func 方法名(入参) 返回值
        自定义类型的方法定义 func (接收者) 方法名(入参) 返回值
	方法的值传递和指针传递
 */

package main

import "fmt"

type person struct {
	name string
	age  int
}

// 普通函数
// 无返回值
func simple(p person) {
	fmt.Println(p)
}
// 有返回值
func simpleR(p person) person {
	return p;
}

// 传递指针(即地址)
func printPerson( p *person ) {
	fmt.Println(p.name)
	fmt.Println(p.age)
}


func main() {
	// 实例化类
	xh := person{"小红",13}
	simple(xh) // {小红 13}

	xl := simpleR(xh)
	fmt.Println(xl) // {小红 13}

	// 打印xh信息
	printPerson(&xh) // 小红 13

}
