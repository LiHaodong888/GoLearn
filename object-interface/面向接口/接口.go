/**
Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口
 */
package main

import "fmt"

// 定义接口 type interface_name interface {}
type Phone interface {
	// 接口方法
	call()
}

//  定义结构体NokiaPhone
type NokiaPhone struct {
}

// 使用值接收者实现接口 call()
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

// 定义结构体IPhone
type IPhone struct {
}
// 使用指针接收者实现接口
func (iPhone *IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}


func main() {
	// 初始化接口
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}
