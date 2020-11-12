package arrayDemo

import (
	"fmt"

)

func ArrayBase()  {
	fmt.Println("array base....")
	/*
		定义数组：
			var 数组名 [数组大小]数据类型
			var a1 [5]int

			定义后五个元素都有默认值0

		数组赋值方式：
			a[0] = 1
			a[1] = 2

		数组的第一个元素的地址就是数组的首地址
		数组各个元素地址间隔，根据数组的数据类型决定， int64 8个字节 int32 4个字节

		数组使用步骤：

			声明数组
			给数组元素赋值
			使用数组
			数组索引从0开始，且不得越界否则panic
			Go数组是值类型，变量传递默认是值传递，因此会进行值拷贝
			修改原本的数组，可以使用引用传递（指针）
	*/
	var intArr [5]int
	fmt.Println("intArr默认值是：",intArr)

	intArr[0] = 1
	intArr[1] = 2
	intArr[2] = 3

	fmt.Println("intArr赋值后的值是：", intArr)
	fmt.Printf("intArr数组地址是=%p\n", &intArr)
	fmt.Printf("intArr数组第一个元素地址是=%p\n", &intArr[0])
	fmt.Printf("intArr数组第二个元素地址是=%p\n", &intArr[1])
	fmt.Printf("intArr数组第三个元素地址是=%p\n", &intArr[2])

	//全局声明
	//声明赋值方式
	var a1 [5]string = [5]string{"cat", "dog"}
	//自动类型推导， 未赋值的有默认值0
	var a2 =[5]int{1,2,3}
	//自动判断数组长度
	var a3 = [...]int{1,2,3,4,5}
	//指定索引赋值元素
	var a4 = [...]string{3:"dogaa", 6: "cataaa"}
	//结构体类型数组
	var a5 = [...]struct {
		name string
		age int
	}{
		{"jason", 10},
		{"sam", 20},
	}

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)


}

func ArrayFor()  {
	var a1 = [...]int{1,2,3,4,5,6}

	//通过索引取值
	for i := 0;i < len(a1); i++{
		fmt.Println(a1[i])
	}

	//for循环遍历数组，索引和值， index可以省略用占位符_
	for index, value := range a1 {
		fmt.Println(index, value)
	}
}

//函数接收值类型，默认有值拷贝
func test(arr [3]int) {
	arr[0] = 66
}

//函数修改原本数组，需要使用引用传递
func test2(arr *[3]int) {
	(*arr)[0] = 66 //可以缩写arr[0]=66 编译器自动识别,arr是指针类型
}

func ArrayDetail()  {
	//数组是多个相同类型数据的组合，且长度固定， 无法扩容
	var a1 [3]int
	a1[0] = 1
	//a2[1] = 11
	//必须赋值int类型数据，否则报错
	//a1[2] = 111.1
	//不得超出索引
	//a1[3]=111
	fmt.Println(a1)//有默认值[1 11 0]

	//声明arr数组，需要考虑传递函数参数时，数组的长度一致性
	arr := [3]int{11, 22, 33}
	//test函数不会修改数组
	test(arr)
	fmt.Println(arr)
	//test2修改了数组
	test2(&arr)
	fmt.Println(arr)
}

func ArrayDemo()  {
	//array
	ArrayBase()
	// 遍历数组
	ArrayFor()
	//细节
	ArrayDetail()
}

