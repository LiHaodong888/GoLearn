/**
Go 语言内置的运算符有：
    算术运算符
    关系运算符
    逻辑运算符
    位运算符
    赋值运算符
    其他运算符
 */
package main

import "fmt"

// 算术运算符
func arithmeticOperator() {
	var a int = 20;
	var b int = 10;
	var c int

	// 相加
	c = a + b
	fmt.Println(c)
	// 相减
	c = a - b
	fmt.Println(c)
	// 相乘
	c = a * b
	fmt.Println(c)
	// 相除
	c = a / b
	fmt.Println(c)
	// 求余
	c = a % b
	fmt.Println(c)
	// 自增
	a++
	fmt.Println(a)
	// 自减
	b--
	fmt.Println(b)
}

// 关系运算符
func relationalOperator() {
	var a int = 20;
	var b int = 10;
	// 检查两个值是否相等，如果相等返回 True 否则返回 False
	fmt.Println(a == b)
	// 检查两个值是否不相等，如果不相等返回 True 否则返回 False
	fmt.Println(a != b)
	// 检查左边值是否大于右边值，如果是返回 True 否则返回 False
	fmt.Println(a > b)
	// 检查左边值是否小于右边值，如果是返回 True 否则返回 False
	fmt.Println(a < b)
	// 检查左边值是否大于等于右边值，如果是返回 True 否则返回 False
	fmt.Println(a >= b)
	// 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False
	fmt.Println(a <= b)
}

// 逻辑运算符
func logicalOperators() {
	var a bool = true
	var b bool = false
	// 逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False
	fmt.Println(a && b)
	// 逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False
	fmt.Println(a || b)
	// 逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True
	fmt.Println(!(a && b))

}

// 位运算符
// 位运算符对整数在内存中的二进制位进行操作
func bitOperator() {
	// 假定 A = 60; B = 13; 其二进制数转换为
	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0
	// 按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与
	c = a & b /* 12 = 0000 1100 */
	fmt.Println(c)

	// 按位或运算符"|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或
	c = a | b /* 61 = 0011 1101 */
	fmt.Println(c)

	// 按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。
	c = a ^ b /* 49 = 0011 0001 */
	fmt.Println(c)

	// 左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。
	c = a << 2 /* 240 = 1111 0000 */
	fmt.Println(c)

	// 右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。
	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Println(c)

}

// 赋值运算符
func assignmentOperator() {
	var a int = 21
	var c int

	// 简单的赋值运算符，将一个表达式的值赋给一个左值
	c = a
	fmt.Println(c)

	// 相加后再赋值 c += a 等于 c = c + a
	c += a
	fmt.Println(c)

	// 相减后再赋值 c -= a 等于 c = c - a
	c -= a
	fmt.Println(c)

	// 相乘后再赋值 c *= a 等于 c = c * a
	c *= a
	fmt.Println(c)

	// 相除后再赋值 c /= a 等于 c = c / a
	c /= a
	fmt.Println(c)

	c = 200;

	// 左移后赋值 c <<= a 等于 c = c << a
	c <<= 2
	fmt.Println(c)

	// 右移后赋值 c >>= a 等于 c = c >> a
	c >>= 2
	fmt.Println(c)

	// 按位与后赋值 c &= a 等于 c = c & a
	c &= 2
	fmt.Println(c)

	// 按位异或后赋值 c ^= a 等于 c = c ^ a
	c ^= 2
	fmt.Println(c)

	// 按位或后赋值 c |= a 等于 c = c | a
	c |= 2
	fmt.Println(c)
}

// 
func otherOperators() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	/*  & 和 * 运算符实例 */
	// &返回变量存储地址
	ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Println(a)
	// *是一个指针变量
	fmt.Println(*ptr)
}

// 依次调用
func main() {
	arithmeticOperator()
	relationalOperator()
	logicalOperators()
	bitOperator()
	assignmentOperator()
	otherOperators()
}
