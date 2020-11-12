/*
   @Author: Jason Chen
   @Date: 2020/11/10 10:32
   @Code: UTF-8
*/
package sliceDemo

import (
	"fmt"
)
/*
1. go语言的切片
2. 切片是可动态变换的序列，是对数组的引用， 引用类型，遵循引用传递机制
3. slice类型写作[]T, T是slice元素类型， var s1 []int, s1就是切片变量
*/

func SliceBase()  {
	//创建一个数组
	var arr1 [5]int = [...]int{11,22,33,44,55}
	/*
		创建切片，通过对数组的索引切片
		s1 是切片名
		arr1[1:3]代表slice引用数组区间，索引1到索引3的值，注意取头不取尾
	*/
	s1 := arr1[1:4]
	fmt.Println(arr1)
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
}

func SliceTheory()  {
	//创建数组， Months月份，一月到12月份
	months := [...]string{
		"","January","February","March","April","May","June","July","August","September","October","November","December",
	}
	//创建切片，对数组的引用
	s1 := months[4:7]
	s2 := months[6:9]
	fmt.Println(s1)
	fmt.Println(s2)

	//指针： 指针指向slice第一个元素对应的数组地址
	fmt.Printf("slice first address %p\n", &s1[0])
	fmt.Printf("对应数组元素的地址%p\n", &months[4])

}

var slice0 []int = make([]int, 10)
var slice1 = make([]int, 10)
var slice2 = make([]int, 10, 10)

func SliceWays()  {
	//内置的make函数，参数（类型，len.cap）, 注意cap大于len，容量可以省略，默认等于长度，切片有默认值
	
	fmt.Printf("make全局slice0 ：%v\n", slice0)
	fmt.Printf("make全局slice1 ：%v\n", slice1)
	fmt.Printf("make全局slice2 ：%v\n", slice2)

	fmt.Println("------------------------------------------------------------------------")


	slice3 := make([]int, 10)
	slice4 := make([]int, 10)
	slice5 := make([]int, 10, 10)
	slice5[0] = 11
	slice5[1] = 22

	fmt.Printf("make局部slice3 : %v\n", slice3)
	fmt.Printf("make局部slice4 : %v\n", slice4)
	fmt.Printf("make局部slice5 : %v\n", slice5)
}

func SliceAppend()  {
	/*
	append原理就是对底层数组扩容，go会创建新的数组，将原本元素拷贝到新的数组中
	slice重新引用新的数组
	这个数组不可见
	*/

	//创建切片
	var slice1 []int = []int{100, 200, 300}
	fmt.Printf("slice1容量=%v 长度=%v\n", cap(slice1), len(slice1))
	//给切片追加新元素
	//容量扩容机制是2倍扩容
	slice1 = append(slice1, 400)
	fmt.Printf("slice1扩容后容量=%v 长度=%v\n", cap(slice1), len(slice1))
	fmt.Println(slice1)
	//切片扩容切片,slice1... 语法糖代表展开切片元素
	slice1=append(slice1,slice1...)
	fmt.Println(slice1)
}

func SliceFor()  {
	var arr [5]int = [...]int{11, 22, 33, 44, 55}
	s1 := arr[1:4]
	//for循环遍历
	for i := 0; i < len(s1); i++ {
		fmt.Printf("s1[%v]=%v\n", i, s1[i])
	}
	fmt.Println()
	//for range方式遍历切片
	for i, v := range s1 {
		fmt.Printf("索引i=%v 值v=%v\n", i, v)
	}
}

func SliceCopy()  {
	var slice1 []int = []int{11,22,33,44}
	//make创建切片，长度是10
	var slice2 = make([]int, 10)
	copy(slice2, slice1) //把slice1 的值拷贝给slice2
	fmt.Println(slice1)//[11 22 33 44]
	fmt.Println(slice2) //[11 22 33 44 0 0 0 0 0 0]

	//slice1和slice2，两者的数据石象湖独立的，互不影响
	slice1[0] = 1233
	fmt.Println(slice1)
	fmt.Println(slice2)

}

func SliceAll()  {
	//全切片
	//官网资料
	// https://golang.google.cn/ref/spec#Slice_expressions

	data := [...]int{0,1,2,3,4,5,6: 0}
	s := data[1:2:3]
	fmt.Printf("s的容量是：%v\n", cap(s))
	s = append(s, 100, 200)         // 一次 append 两个值，超出 s.cap 限制。
	fmt.Println(s, data)            // 重新分配底层数组，与原数组无关。
	fmt.Printf("s的容量=%v\n", cap(s)) //二倍扩容
	fmt.Println(&s[0], &data[0])    // 比对底层数组起始指针。

}

func SliceAndString()  {
	str1 := "yugo niubi"
	//string是不可变的，也无法通过切片修改值
	//str1[0] = 's'  编译器失败
	//修改string的方法，需要string转化为[]byte，修改后转为string
	arr1 := []byte(str1) //类型强转
	arr1[0] = 'g'
	str1 = string(arr1)
	fmt.Printf("str1=%v\n", str1)
	//[]byte只能处理英文和数字，不能处理汉字，汉字3个字节，会出现乱码
	//将string转为[]rune，按字符处理，兼容汉字
	arr2 := []rune(str1)
	arr2[0] = '于'
	str1 = string(arr2)
	fmt.Printf("str1=%v\n", str1)
}

func SliceDemo()  {
	fmt.Println("slice ")
	//基础
	SliceBase()
	//切片原理
	SliceTheory()
	//切片方法
	SliceWays()
	//添加
	SliceAppend()
	//遍历
	SliceFor()
	//拷贝
	SliceCopy()
	//全切片
	SliceAll()
	//字符串切片的两种方式
	SliceAndString()
}

