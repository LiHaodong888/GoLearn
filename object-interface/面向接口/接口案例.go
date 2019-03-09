/**
假设某公司有两个员工，一个普通员工和一个高级员工， 但是基本薪资是相同的，高级员工多拿奖金。计算公司为员工的总开支
 */
package main

import (
	"fmt"
)

// 薪资计算器接口
type SalaryCalculator interface {
	CalculateSalary() int
}
// 普通挖掘机员工
type Contract struct {
	empId  int
	basicpay int
}
// 有蓝翔技校证的员工
type Permanent struct {
	empId  int
	// 底薪
	basicpay int
	// 奖金
	jj int
}
// 普通挖掘机员工实现接口
func (c Contract) CalculateSalary() int {
	// 只有底薪
	return c.basicpay
}

// 蓝翔技校证的员工实现接口
func (p Permanent) CalculateSalary() int {
	// 底薪+奖金
	return p.basicpay + p.jj
}

// 总开支
func totalExpense(s []SalaryCalculator) {
	// 初始化总支出为0
	expense := 0
	// 进行遍历加值
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("总开支 $%d", expense)
}

func main() {
	pemp1 := Permanent{1,3000,10000}
	pemp2 := Permanent{2, 3000, 20000}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)
}