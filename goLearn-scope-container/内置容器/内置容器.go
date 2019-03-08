/**
    内建容器
		1. 数组
        2. Slice 切片
        3. Map 映射
 */
package main

import "fmt"

/**
1. 数组
声明数组 形式: var variable_name [SIZE] variable_type
数组赋值
初始化数组 var variable_name = [SIZE]variable_type{}
访问数组元素
 */

func array() {
	// 声明数组
	var arr [2] int
	// 数组赋值
	arr[0] = 1
	arr[1] = 2

	// 初始化数组元素
	var brr = [2]int{1, 2}

	fmt.Println("arr数组数据: ", arr)
	fmt.Println("brr数组数据: ", brr)

	// 访问数组元素
	for i := 0; i < len(arr); i++ {
		fmt.Println("arr数组元素: ", arr[i])
	}

}

/**
2. 切片
Go 语言切片是对数组的抽象。
Go 数组的长度不可改变，在特定场景中这样的集合就不太适用
Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

定义切片
	形式:
	var identifier []type
	var slice1 []type = make([]type, len)
	也可以简写为
	slice1 := make([]type, len)
初始化切片
len() 和 cap() 函数
	切片是可索引的，并且可以由 len() 方法获取长度
	切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少
空(nil)切片
切片截取
append() 和 copy() 函数
	如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来
 */

func slice() {
	// 1.定义切片
	var a []int
	// 使用make函数创建一个int切片，长度和容量都是10
	b := make([]int, 10)

	// 2.初始化切片
	arr := [5]int{1, 2, 3, 4, 5}

	// 3.len() 和 cap() 函数
	// a切片是空的 为0
	fmt.Println(len(a),len(b),len(arr),cap(arr))
	// 空(nil)切片
	if (a == nil) {
		fmt.Println("切片是空的")
	}

	// 4.切片截取
	/* 创建切片 */
	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	numbers1 := make([]int,0,5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)


	// 4.append() 和 copy() 函数
	var brr []int
	printSlice(brr)

	/* 允许追加空切片 */
	brr = append(brr, 0)
	printSlice(brr)

	/* 向切片添加一个元素 */
	brr = append(brr, 1)
	printSlice(brr)

	/* 同时添加多个元素 */
	brr = append(brr, 2,3,4)
	printSlice(brr)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	brr1 := make([]int, len(brr), (cap(brr))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(brr1,brr)
	printSlice(brr1)
}


func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

/**
3. map集合
Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，
这是因为 Map 是使用 hash 表来实现的。

定义map
	声明变量，默认 map 是 nil
	var map_variable map[key_data_type]value_data_type
	使用 make 函数
	map_variable := make(map[key_data_type]value_data_type)
map插入数据
delete() 函数 删除元素
 */
func mapCon()  {

	// 定义map
	var mapA map[int]string
	// 初始化
	// 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
	mapA = make(map[int]string)

	// 插入数据
	mapA [1] = "我是key1对应的值"
	mapA [2] = "我是key2对应的值"
	mapA [3] = "我是key3对应的值"

	fmt.Println(mapA)

	// 遍历map
	for key := range mapA {
		fmt.Println("key值",key, "value值", mapA [key])
	}

	// 还可以这样遍历
	for key,value := range mapA{
		fmt.Println(key,value)
	}

	// 查看元素在集合中是否存在
	// captial 若存在返回key对应的value 不存在没有值 一般都是通过ok进行判断
	// ok 不存在 false/ 存在 true
	captial, ok := mapA [ 4 ]
	if (ok) {
		fmt.Println("存在的", captial)
	} else {
		fmt.Println("不存在")
	}

	// delete() 函数 删除元素
	delete(mapA, 1)

	// 删除之后的遍历map 发现key1已经没了
	for key := range mapA {
		fmt.Println("key值",key, "value值", mapA [key])
	}


}


// 依次调用
func main() {
	array()
	slice()
	mapCon()
}
