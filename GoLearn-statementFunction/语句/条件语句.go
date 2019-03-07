/**
1. if语句 if 语句 由一个布尔表达式后紧跟一个或多个语句组成。
2. if ... else 语句 if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行
3. if嵌套 你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句
4. switch语句 switch 语句用于基于不同条件执行不同动作
5. select语句 select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行
 */
package main

import "fmt"

// 1. if语句
/**
if 布尔表达式 {
   在布尔表达式为 true 时执行
}
 */
func ifStatement() {
	/* 定义局部变量 */
	var a int = 10

	if (a > 10) {
		fmt.Println("a大于10")
	}
	fmt.Println("a小于10")
}

// 2. if ... else
/**
if 布尔表达式 {
   在布尔表达式为 true 时执行
} else {
在布尔表达式为 false 时执行
}
 */
func ifElse() {
	/* 定义局部变量 */
	var a int = 10

	if (a > 10) {
		fmt.Println("a大于10")
	} else {
		fmt.Println("a小于10")
	}

}

// 3. if嵌套
/**
if 布尔表达式 1 {
   在布尔表达式 1 为 true 时执行
if 布尔表达式 2 {
在布尔表达式 2 为 true 时执行
}
}
 */
func ifNesting() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	/* 判断条件 */
	if a == 100 {
		/* if 条件语句为 true 执行 */
		if b == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Println("a 的值为 100 ， b 的值为 200");
		}
	}
	fmt.Println("a 值为 :", a);
	fmt.Println("b 值为 :", b);
}

// 4. switch语句
/**
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
 */
func switchStatement(a int) {

	switch a {
	case 5:
		fmt.Println(a)
	case 6:
		fmt.Println(a)
	}
}

// 5. select语句 这个是与java不同的语句
/**
select是Go中的一个控制结构，类似于用于通信的switch语句。每个case必须是一个通信操作，要么是发送要么是接收。
select随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。一个默认的子句应该总是可运行的。
 */
// 信道只是 被声明，但没有被初始化，值为nil，因此，下面的select语句的case子句会被阻塞，
// 因为有default子句，因此，输出no communication
func selectStatement()  {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

}

func main() {
	ifStatement()
	ifElse()
	ifNesting()
	switchStatement(6)
	selectStatement()

}
